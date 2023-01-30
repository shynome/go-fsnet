package dev_test

import (
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"testing"

	"github.com/lainio/err2/assert"
	"github.com/lainio/err2/try"
	"github.com/shynome/go-fsnet"
	"github.com/shynome/go-fsnet/dev"
)

var l net.Listener
var word = "hello world"

func TestMain(m *testing.M) {

	var fsnet = fsnet.New("")
	var OpenFile dev.FileOpener = func(name string, flag int, perm fs.FileMode) (io.ReadWriteCloser, error) {
		f, err := fsnet.Open(name)
		if err != nil {
			return nil, err
		}
		return f.(io.ReadWriteCloser), nil
	}
	dev.SetFileOpener(OpenFile)

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
	t.Log("see: https://github.com/shynome/go-wagi/blob/master/internal/fsnet/wasm_test.go")
}
