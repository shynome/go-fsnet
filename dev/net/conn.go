package devnet

import (
	"io"
	"net"
	"time"
)

type Conn struct {
	io.ReadWriteCloser
	Network string
	Remote  *Addr
}

var _ net.Conn = (*Conn)(nil)

func (c *Conn) RemoteAddr() net.Addr { return c.Remote }
func (c *Conn) LocalAddr() net.Addr {
	return &Addr{
		NetType:  c.Network,
		Hostname: "wasi",
		Port:     0,
	}
}

func (c *Conn) SetDeadline(t time.Time) error      { return nil }
func (c *Conn) SetReadDeadline(t time.Time) error  { return nil }
func (c *Conn) SetWriteDeadline(t time.Time) error { return nil }
