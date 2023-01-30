# Description

wasi network over file system, like bash `/dev/tcp` or `/dev/udp`

# Usage

```go
// see dev/tcp_test.go
func TestFsNet(t *testing.T) {
	client := &http.Client{Transport: dev.Transport}
	addr := l.Addr().String()
	addr = fmt.Sprintf("http://%s", addr)
	req := try.To1(http.NewRequest(http.MethodGet, addr, nil))
	resp := try.To1(client.Do(req))
	b := try.To1(io.ReadAll(resp.Body))
	assert.Equal(string(b), word)
}
```
