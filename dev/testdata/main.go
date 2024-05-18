package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shynome/err0/try"
	"github.com/shynome/go-fsnet/dev"
)

func init() {
	http.DefaultClient = &http.Client{Transport: dev.Transport}
}

func main() {
	go func() {
		for {
			time.Sleep(time.Second)
			log.Println("not hangup")
		}
	}()
	time.Sleep(5 * time.Second)
	addr := fmt.Sprintf("http://%s", os.Args[1])
	resp := try.To1(http.Get(addr))
	io.Copy(os.Stdout, resp.Body)
}
