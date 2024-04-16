package fsnet

import (
	"crypto/tls"
	"io/fs"
	"net"
	"path/filepath"
	"time"

	devnet "github.com/shynome/go-fsnet/dev/net"
)

type Net struct {
	basedir string
}

func New(dir string) fs.FS {
	return &Net{basedir: dir}
}

var _ fs.FS = (*Net)(nil)

func (n *Net) Open(name string) (f fs.File, err error) {
	name = filepath.Join(n.basedir, name)
	if name+"/" == n.basedir {
		return &dir{name}, nil
	}
	addr, err := devnet.ParseAddr(name)
	if err != nil {
		return nil, fs.ErrNotExist
	}
	var conn net.Conn
	if addr.TLS {
		conn, err = tls.Dial(addr.NetType, addr.Address(), nil)
	} else {
		conn, err = net.Dial(addr.NetType, addr.Address())
	}
	if err != nil {
		return
	}
	return NewFile(name, conn), nil
}

type dir struct{ path string }

var _ fs.File = (*dir)(nil)
var _ fs.FileInfo = (*dir)(nil)

func (d *dir) Stat() (fs.FileInfo, error) { return d, nil }
func (dir) Read([]byte) (int, error)      { return 0, nil }
func (dir) Close() error                  { return nil }

func (d dir) Name() string     { return d.path }
func (dir) Size() int64        { return 0 }
func (dir) Mode() fs.FileMode  { return fs.ModeType }
func (dir) ModTime() time.Time { return time.Now() }
func (dir) IsDir() bool        { return true }
func (dir) Sys() any           { return nil }
