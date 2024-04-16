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

	"github.com/shynome/err0/try"
	"github.com/shynome/go-fsnet"
	"github.com/shynome/go-fsnet/dev"
	"github.com/stretchr/testify/assert"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
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
	assert.Equal(t, string(b), word)
}

func TestWasiFsNet(t *testing.T) {
	buildWasm()

	ctx := context.Background()
	rtc := wazero.NewRuntimeConfigInterpreter()
	rt := wazero.NewRuntimeWithConfig(ctx, rtc)
	wasi_snapshot_preview1.MustInstantiate(ctx, rt)

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

	rt.InstantiateModule(ctx, m, mc)

	assert.Equal(t, stdout.String(), word)
}

func buildWasm() {
	cmd := exec.Command("go", "build", "-o", "testdata/main.wasm", "./testdata/main.go")
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "GOOS=wasip1", "GOARCH=wasm")
	try.To(cmd.Run())
}
