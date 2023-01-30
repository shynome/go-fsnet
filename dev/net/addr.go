package devnet

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strconv"
)

type Addr struct {
	NetType  string // tcp or udp
	Hostname string
	Port     int

	TLS bool
}

var _ net.Addr = (*Addr)(nil)

func (Addr) Network() string { return "fsnet" }

func (a *Addr) String() string {
	s := fmt.Sprintf("/dev/%s/%s/%d", a.NetType, a.Hostname, a.Port)
	if a.TLS {
		s += "/tls"
	}
	return s
}

func (a *Addr) Address() string {
	return fmt.Sprintf("%s:%d", a.Hostname, a.Port)
}

var addrSpitter = regexp.MustCompile(`^\/dev\/(tcp|udp)\/(.+)\/(\d+)(/tls|)$`)

var ErrAddrMatchFailed = errors.New("addr match failed")

func ParseAddr(name string) (*Addr, error) {
	n := addrSpitter.FindStringSubmatch(name)
	if len(n) != 5 {
		return nil, ErrAddrMatchFailed
	}

	port, err := strconv.Atoi(n[3])
	if err != nil {
		return nil, err
	}

	addr := &Addr{
		NetType:  n[1],
		Hostname: n[2],
		Port:     port, // n[3]
		TLS:      n[4] != "",
	}

	return addr, nil
}
