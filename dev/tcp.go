package dev

import (
	"bufio"
	"net/http"
	"strconv"

	devnet "github.com/shynome/go-fsnet/dev/net"
)

type DevTCP struct{}

var tcpdev = &DevTCP{}

var portsMap = map[string]string{
	"https": "443",
	"http":  "80",
}

var Transport http.RoundTripper = tcpdev

func (*DevTCP) RoundTrip(req *http.Request) (*http.Response, error) {
	var portStr = req.URL.Port()
	if portStr == "" {
		portStr = portsMap[req.URL.Scheme]
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}
	addr := &devnet.Addr{
		NetType:  "tcp",
		Hostname: req.URL.Hostname(),
		Port:     port,
	}
	if req.URL.Scheme == "https" {
		addr.TLS = true
	}

	f, err := Dial("fsnet", addr)
	if err != nil {
		return nil, err
	}

	req.Write(f)
	resp, err := http.ReadResponse(bufio.NewReader(f), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
