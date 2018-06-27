package main

// Inspired by  https://github.com/inconshreveable/srvdir

import (
	"net/http"
)

func main() {
	p := http.Dir("/")
	fs := http.FileServer(p)
	http.Handle("/static", fs)
}
