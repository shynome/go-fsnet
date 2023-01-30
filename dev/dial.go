package dev

import (
	"net"
	"os"

	devnet "github.com/shynome/go-fsnet/dev/net"
)

func Dial(network string, raddr *devnet.Addr) (net.Conn, error) {
	f, err := OpenFile(raddr.String(), os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	conn := &devnet.Conn{
		ReadWriteCloser: f,
		Remote:          raddr,
	}
	return conn, nil
}
