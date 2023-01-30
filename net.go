package fsnet

import (
	"crypto/tls"
	"io/fs"
	"net"
	"path/filepath"

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
