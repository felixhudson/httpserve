package main

// Inspired by  https://github.com/inconshreveable/srvdir

import (
	"fmt"
	"net/http"
)

func main() {
	// todo Change to the current directory, or one given on command line
	p := http.Dir("/")
	fs := http.FileServer(p)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "home page")
	})
	mux.HandleFunc("/s2/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "home page 2")
	})
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe("0.0.0.0:8888", mux)
}
