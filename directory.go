package main

import (
	"path/filepath"
	"strings"
	"unicode"
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
	// lets cheat by always having a non-number at the end of the string!
	input = input + string("_")

	// find first number
	start := -1
	result := make([]string, 0)
	// strlength := len(input)
	for k, v := range input {
		if unicode.IsDigit(v) {
			if start == -1 {
				start = k
			}
		} else {
			if start != -1 {
				if k-start > length {
					// insert the padding zeros
					result = append(result, string('0'))
				}
				// append the rest of the numeral
				result = append(result, input[start:k])
				// reset the counts
				start = -1
				// continue
			}
			result = append(result, string(v))
		}
	}

	return strings.Join(result, "")
}

// NaturalSort will sort how a human wants to
func NaturalSort(files []string) []string {
	//find the longest number string
	//longest := 3

	// append 0's to the length of the longest number
	// sort by first key
	// return the last
	return []string{"foo"}
}
func countNumberChars(filename string) int {
	count, longest := 0, 0
	for _, v := range filename {
		if !unicode.IsDigit(v) {
			count = 0
		} else {
			count++
		}

		if count > longest {
			longest = count
		}
	}
	return longest
}

// NextTwoFiles will return the next two files in a sorted directory
func NextTwoFiles(path string) string {
	files := sortDir(path)
	return files[0]
}
