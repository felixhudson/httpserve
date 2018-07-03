package main

import (
	"fmt"
	"path/filepath"
	"sort"
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
	// input = input + string("_")

	// find first number
	start := -1
	result := make([]string, 0)
	state := "other"
	strlength := len(input)
	for k, v := range input {
		if k+1 == strlength {
			if state == "other" {
				state = "otherend"
			} else {
				state = "numberend"
			}
		}

		switch state {
		case "number":
			if !unicode.IsDigit(v) {
				if k-start < length {
					result = append(result, string('0'))
				}
				result = append(result, input[start:k])
				result = append(result, string(v))
				// reset the counts
				state = "other"
			}
		case "other":
			if unicode.IsDigit(v) {
				state = "number"
				start = k
			} else {
				result = append(result, string(v))
			}
		case "otherend":
			if unicode.IsDigit(v) {
				// we have reached a single digit at the end of the input
				result = append(result, string('0'))

			}
			result = append(result, string(v))

		case "numberend":
			if unicode.IsDigit(v) {
				if k-start < length {
					result = append(result, string('0'))
				}
				result = append(result, string(input[start:k]))
			} else {
				if k-start < length {
					result = append(result, string('0'))
				}
				result = append(result, string(input[start:k]))
				result = append(result, string(v))
			}

		}
	}
	return strings.Join(result, "")
}

// NaturalSort will sort how a human wants to
func NaturalSort(files []string) []string {
	//find the longest number string
	longest := 0
	temp := 0
	for _, v := range files {
		temp = countNumberChars(v)
		if temp > longest {
			longest = temp
		}
	}

	// append 0's to the length of the longest number
	paddedfilenames := make([]string, 0)
	for _, v := range files {
		paddedfilenames = append(paddedfilenames, padZeros(v, longest))
	}
	sorted := sort.StringSlice(paddedfilenames)
	for _, v := range sorted {
		fmt.Printf("%v", v)
	}
	// sort by first key
	// return the last
	return sorted
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
