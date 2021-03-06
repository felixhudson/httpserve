package main

import (
	"fmt"
	"path/filepath"
	"strconv"
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

func Test_count_numbers(t *testing.T) {
	text := "a"
	expect := 0
	result := countNumberChars(text)

	if result != expect {
		t.Fatal(text + " returned wrong count of " + strconv.Itoa(result) + "::")
	}
	text = "a1"
	expect = 1
	result = countNumberChars(text)
	if result != expect {
		t.Fatal(text + " returned wrong count of " + strconv.Itoa(result) + "::")
	}
	text = "a10"
	expect = 2
	result = countNumberChars(text)
	if result != expect {
		t.Fatal(text + " returned wrong count of " + strconv.Itoa(result) + "::")
	}
}
func Test_pad(t *testing.T) {
	text := "a1"
	expect := "a01"
	result := padZeros(text, 2)
	if result != expect {
		t.Fatal(text + " returned wrong text of " + result + "::")
	}
	text = "a1a"
	expect = "a01a"
	result = padZeros(text, 2)
	if result != expect {
		t.Fatal(text + " returned wrong text of " + result + "::")
	}
	text = "a10a"
	expect = "a10a"
	result = padZeros(text, 2)
	if result != expect {
		t.Fatal(text + " returned wrong text of " + result + "::")
	}
	text = "a1a1"
	expect = "a01a01"
	result = padZeros(text, 2)
	if result != expect {
		t.Fatal(text + " returned wrong text of " + result + "::")
	}
}

func Test_hash_sort(t *testing.T) {
	// give a map of [string]string
	data := make(map[string]string)
	data["01.txt"] = "1.txt"
	data["11.txt"] = "11.txt"
	data["02.txt"] = "2.txt"
	result := sortmap(data)
	expect := []string{"1.txt", "2.txt", "11.txt"}
	for k, v := range result {

		if v != expect[k] {
			t.Fatal("Expected:" + expect[k] + " Got " + v + "::")
		}
	}

}

func Test_natural_sort(t *testing.T) {
	files := []string{"1.txt", "11.txt", "2.txt"}
	expect := []string{"1.txt", "2.txt", "11.txt"}
	result := NaturalSort(files)
	if len(result) != len(expect) {
		t.Fatal("Natural sort lengths dont match " + strconv.Itoa(len(result)) + ":" + strconv.Itoa(len(expect)))
	}
	for k, v := range result {

		if v != expect[k] {
			t.Fatal("Expected:" + expect[k] + " Got " + v + "::")
		}
	}

}
