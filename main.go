package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", "80", "port to check")
	ip := flag.String("ip", "127.0.0.1", "ip address to check")
	path := flag.String("path", "/health", "path to check")
	flag.Parse()

	resp, err := http.Get(fmt.Sprintf("http://%s:%s/%s", *ip, *port, *path))
	if err != nil || resp.StatusCode != 200 {
		os.Exit(1)
	}
	os.Exit(0)
}
