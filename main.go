package main

// Inspired by  https://github.com/inconshreveable/srvdir

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func myHandler(w http.ResponseWriter, req *http.Request) {
	log(w, req)
	fmt.Printf("%v", req.URL)
	if strings.Contains(req.URL.Path, "jpg") {
		page := pagestart + showimage(req.URL.Path, "next.jpg") + pageend
		fmt.Fprint(w, page)
	} else {
		fmt.Fprintf(w, NextTwoFiles("/"))
		fmt.Printf("%v", w)
		fmt.Printf("%v", req)
	}
}

func log(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(w, "Got request %s %s", req.URL.Path, req.URL.Host)
	t := string(time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))
	fmt.Printf("Got request %s %s %s", t, req.URL.Path, req.URL.Host)
}

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
	mux.HandleFunc("/test/", myHandler)
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe("0.0.0.0:8888", mux)
}
