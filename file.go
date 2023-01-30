package fsnet

import (
	"io"
	"io/fs"
	"net"
	"time"
)

type File struct {
	conn net.Conn
	path string
}

var _ fs.File = (*File)(nil)

func NewFile(name string, conn net.Conn) *File {
	return &File{
		conn: conn,
		path: name,
	}
}

func (fc *File) Stat() (fs.FileInfo, error) { return fc, nil }
func (fc *File) Read(p []byte) (n int, err error) {
	n, err = fc.conn.Read(p)
	return
}
func (fc *File) Close() (err error) {
	err = fc.conn.Close()
	return
}

var _ io.Writer = (*File)(nil)

func (fc *File) Write(p []byte) (n int, err error) {
	n, err = fc.conn.Write(p)
	return
}

var _ fs.FileInfo = (*File)(nil)

func (info *File) Name() string       { return info.path }
func (info *File) Size() int64        { return 4096 }
func (info *File) Mode() fs.FileMode  { return fs.ModeType }
func (info *File) ModTime() time.Time { return time.Now() }
func (info *File) IsDir() bool        { return false }
func (info *File) Sys() any           { return nil }
