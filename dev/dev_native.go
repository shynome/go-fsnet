//go:build !wasm
// +build !wasm

package dev

import (
	"io"
	"io/fs"

	"github.com/shynome/go-fsnet"
)

func init() {
	var fsnet = fsnet.New("")
	var OpenFile FileOpener = func(name string, flag int, perm fs.FileMode) (io.ReadWriteCloser, error) {
		f, err := fsnet.Open(name)
		if err != nil {
			return nil, err
		}
		return f.(io.ReadWriteCloser), nil
	}
	SetFileOpener(OpenFile)
}
