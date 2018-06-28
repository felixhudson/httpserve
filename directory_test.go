package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_directory(t *testing.T) {
	if !isGallery("/") {
		t.Fatal()
	}
}

func Test_sort(t *testing.T) {
	foo := sortDir("/Users")
	if len(foo) <= 0 {
		t.Fatal()
	}
}

func Test_next(t *testing.T) {
	files := []string{"first.txt", "seccond.txt"}
	foo := nextFile("fake", files)
	if foo != "" {
		fmt.Println("failed test")
		t.Fatal()
	}

	foo = nextFile(files[0], files)
	if foo != "seccond.txt" {
		fmt.Printf("foo: %v\n", foo)
		fmt.Println("failed test")
		t.Fatal()
	}
}

func Test_glob(t *testing.T) {
	files, _ := filepath.Glob("/*")
	if len(files) <= 0 {
		t.Fatal("got zero length files list")
	}
}
