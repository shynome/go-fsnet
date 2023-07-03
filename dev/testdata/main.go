package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lainio/err2/try"
	"github.com/shynome/go-fsnet/dev"
)

func init() {
	http.DefaultClient = &http.Client{Transport: dev.Transport}
}

func main() {
	addr := fmt.Sprintf("http://%s", os.Args[1])
	resp := try.To1(http.Get(addr))
	io.Copy(os.Stdout, resp.Body)
}
