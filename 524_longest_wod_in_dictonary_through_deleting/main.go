package main

import (
	"fmt"
)

func main() {
	s := "abpcplea"
	d := []string{"ale", "apple", "monkey", "plea"}
	fmt.Printf("Longest word found is : %s \n", findLongestWord(s, d))

}

func findLongestWord(s string, d []string) string {
	result := ""

	for _, v := range d {
		if present(v, s) {
			if len(v) > len(result) {
				result = v
			} else if len(v) == len(result) {
				if v < result {
					result = v
				}
			}
		}
	}

	return result
}

func present(word, available string) bool {
	present := false
	pos := 0
	l := len(word)
	require := word[pos]

	for k := range available {
		// check if requirement is met
		if available[k] == require {
			// requirement met

			// check if we achieved the word
			if pos == l-1 {
				// the entire word found
				present = true
				break
			}

			// word not finished, move forward
			pos++
			require = word[pos]
		}
	}

	return present
}
