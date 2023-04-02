package dev_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"testing"

	"github.com/lainio/err2/assert"
	"github.com/lainio/err2/try"
	"github.com/shynome/go-fsnet"
	"github.com/shynome/go-fsnet/dev"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/experimental/gojs"
)

var l net.Listener
var word = "hello world"

func TestMain(m *testing.M) {
	l = try.To1(net.Listen("tcp", "127.0.0.1:0"))
	defer l.Close()
	go http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, word)
	}))

	m.Run()
}

func TestFsNet(t *testing.T) {
	client := &http.Client{Transport: dev.Transport}
	addr := l.Addr().String()
	addr = fmt.Sprintf("http://%s", addr)
	req := try.To1(http.NewRequest(http.MethodGet, addr, nil))
	resp := try.To1(client.Do(req))
	b := try.To1(io.ReadAll(resp.Body))
	assert.Equal(string(b), word)
}

func TestWasiFsNet(t *testing.T) {
	buildWasm()

	ctx := context.Background()
	rtc := wazero.NewRuntimeConfigInterpreter()
	rt := wazero.NewRuntimeWithConfig(ctx, rtc)
	gojs.MustInstantiate(ctx, rt)

	mb := try.To1(os.ReadFile("testdata/main.wasm"))
	m := try.To1(rt.CompileModule(ctx, mb))

	mc := wazero.NewModuleConfig()
	fsc := wazero.NewFSConfig()
	fsc = fsc.WithFSMount(fsnet.New("/dev/"), "/dev")
	mc = mc.WithFSConfig(fsc)
	var stdout bytes.Buffer
	mc = mc.
		WithArgs("wasi", l.Addr().String()).
		WithStderr(os.Stderr).
		WithStdout(&stdout)

	gojs.Run(ctx, rt, m, gojs.NewConfig(mc))

	assert.Equal(stdout.String(), word)
}

func buildWasm() {
	cmd := exec.Command("go", "build", "-o", "testdata/main.wasm", "./testdata/main.go")
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	try.To(cmd.Run())
}
