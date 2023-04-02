package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lainio/err2/try"
	"github.com/shynome/go-fsnet/dev"
)

func main() {
	client := &http.Client{Transport: dev.Transport}
	addr := fmt.Sprintf("http://%s", os.Args[1])
	req := try.To1(http.NewRequest(http.MethodGet, addr, nil))
	resp := try.To1(client.Do(req))
	io.Copy(os.Stdout, resp.Body)
}
