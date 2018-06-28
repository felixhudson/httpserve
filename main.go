package main

// Inspired by  https://github.com/inconshreveable/srvdir

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func myHandler(w http.ResponseWriter, req *http.Request) {
	log(w, req)
	cwd, _ := os.Getwd()
	fmt.Printf("%v", req.URL)
	if strings.Contains(req.URL.Path, "jpg") {
		relativePath := strings.TrimPrefix(req.URL.Path, "/test")
		staticPath := "/static" + relativePath
		dir, file := filepath.Split(relativePath)
		cwdpath := filepath.SplitList(cwd)[0]

		nextPath := findNextFile(file, filepath.Join(cwdpath, dir))
		fmt.Printf("nextpath: %s", nextPath)
		// nextPath := findNextFile(splits[0:len(splits)-2], splits[len(splits)-1])
		page := pagestart + showimage(staticPath, nextPath) + pageend
		fmt.Fprint(w, page)
	} else {
		f, _ := os.Open(cwd)
		myDirList(w, req, f)
		// fmt.Printf(NextTwoFiles("/"))
		// fmt.Printf("%v", w)
		// fmt.Printf("%v", req)
	}
}

type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}

var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	// "&#34;" is shorter than "&quot;".
	`"`, "&#34;",
	// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	"'", "&#39;",
)

type byName []os.FileInfo

func (s byName) Len() int           { return len(s) }
func (s byName) Less(i, j int) bool { return s[i].Name() < s[j].Name() }
func (s byName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func myDirList(w http.ResponseWriter, r *http.Request, f File) {
	dirs, err := f.Readdir(-1)
	if err != nil {
		//logf(r, "http: error reading directory: %v", err)
		//Error(w, "Error reading directory", StatusInternalServerError)
		return
	}
	//sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })
	sort.Sort(byName(dirs))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<pre>\n")
	for _, d := range dirs {
		name := d.Name()
		if d.IsDir() {
			name += "/"
		}
		// name may contain '?' or '#', which must be escaped to remain
		// part of the URL path, and not indicate the start of a query
		// string or fragment.
		url := url.URL{Path: name}
		fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", url.String(), htmlReplacer.Replace(name))
	}
	fmt.Fprintf(w, "</pre>\n")
}

func log(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(w, "Got request %s %s", req.URL.Path, req.URL.Host)
	t := string(time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))
	fmt.Printf("Got request %s %s %s\n", t, req.URL.Path, req.URL.Host)
}

func main() {
	stuff()
	// todo Change to the current directory, or one given on command line
	cwd, _ := os.Getwd()
	p := http.Dir(cwd)
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

func stuff() {
	fmt.Println("Stuff!")
	fmt.Println(os.Args)
	cwd, _ := os.Getwd()
	fmt.Println(cwd)
	fmt.Println("end of stuff!")
}
