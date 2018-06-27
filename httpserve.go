package main

// Inspired by  https://github.com/inconshreveable/srvdir

import (
	"net/http"
)

func main() {
	p := http.Dir("/")
	fs := http.FileServer(p)
	http.ListenAndServe("0.0.0.0:8888", fs)
}
