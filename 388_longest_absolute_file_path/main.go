package main

import (
	"fmt"
	"strings"
)

func main() {
	// input := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	input2 := "dir\n        file.txt"
	fmt.Printf("Size of longest file path is : %v \n", lengthLongestPath(input2))
	// for _, v := range input2 {
	// 	fmt.Printf("got : %v - %v \n", v, string(v))
	// }
}

func lengthLongestPath(input string) int {
	start := 0
	end := 0
	parts := []string{}
	best := 0
	steps := 0
	spaces := 0
	for k, v := range input {
		if v == '\n' {
			end = k
			// fmt.Printf("Start is %v , end is %v \n", start, end)
			part := input[start:end]
			// part = strings.TrimSpace(part)
			// fmt.Printf("Part is %v \n", part)
			parts = append(parts, part)
			// fmt.Println(parts)
			if isfile(part) {
				fullName := strings.Join(parts, "/")
				l := len(fullName)
				if l > best {
					best = l
				}
			}
		} else {
			if v == 32 {
				spaces++
				continue
			}
			if v == '\t' {
				steps++
				start = k + 1
			} else {
				fmt.Printf("Spaces is %v\n", spaces)
				if spaces%4 == 0 {
					// steps++
					start = k
				}
				spaces = 0
				if steps != 0 {
					parts = parts[:steps]
				}
				steps = 0
			}
		}
	}

	// fmt.Printf("Start at end is : %v \n", start)
	part := input[start:]
	// part = strings.TrimSpace(part)
	parts = append(parts, part)
	fmt.Println(parts)
	if isfile(part) {
		fullName := strings.Join(parts, "/")
		l := len(fullName)
		if l > best {
			best = l
		}
	}

	return best
}

func isfile(name string) bool {
	ok := false

	for _, v := range name {
		if v == '.' {
			ok = true
			break
		}
	}

	return ok
}
