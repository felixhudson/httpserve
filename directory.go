package main

import (
	"path/filepath"
)

func isGallery(path string) bool {
	return true
}

// need to write a func that returns numerical sorted files in a directory
func sortDir(path string) []string {
	files, _ := filepath.Glob(path + "/*")
	return files
	//return []string{"one.jpg", "two.jpg"}
}

func findNextFile(current string, path string) string {
	sortfiles := sortDir(path)
	return nextFile(current, sortfiles)
}

func nextFile(current string, files []string) string {
	nextFile := -1
	for k, v := range files {
		if v == current {
			nextFile = k + 1
			break
		}
	}
	if nextFile > 0 {
		return files[nextFile]
	}
	return ""
}

func padZeros(input string, length int) string {
	// find first number
	start := 0
	for k, v := range input {
		if v == '0' {
			start = k
		}
	}

	return input[start:3]
}

func NaturalSort(files []string) []string {
	//find the longest number string
	//longest := 3

	// append 0's to the length of the longest number
	// sort by first key
	// return the last
	return []string{"foo"}
}

func NextTwoFiles(path string) string {
	files := sortDir(path)
	return files[0]
}
