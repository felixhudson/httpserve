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

	foo := nextFile("/Users", "doesntexist")
	if foo != "" {
		fmt.Println("failed test")
		t.Fatal()
	}

	foo = nextFile("/Users", "Felix")
	if foo != "Public" {
		fmt.Printf("foo: %v\n", foo)
		fmt.Println("failed test")
		t.Fatal()
	}
}

func Test_glob(t *testing.T) {
	files, _ := filepath.Glob("/*")
	t.Log(files)
}
