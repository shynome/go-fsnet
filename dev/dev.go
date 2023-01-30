package dev

import (
	"io"
	"io/fs"
	"os"
)

type FileOpener func(name string, flag int, perm fs.FileMode) (io.ReadWriteCloser, error)

var OpenFile FileOpener = func(name string, flag int, perm fs.FileMode) (io.ReadWriteCloser, error) {
	return os.OpenFile(name, flag, perm)
}

func SetFileOpener(f FileOpener) {
	OpenFile = f
}
