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

func nextFile(current string, path string) string {
	sortfiles := sortDir(path)
	nextFile := -1
	for k, v := range sortfiles {
		if v == current {
			nextFile = k + 1
			break
		}
	}
	if nextFile > 0 {
		return sortfiles[nextFile]
	}
	return ""
}
