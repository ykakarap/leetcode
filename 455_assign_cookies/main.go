package main

import (
	"fmt"
	"sort"
)

func main() {

	g := []int{1, 2}
	s := []int{1, 2, 3}
	fmt.Printf("No of content children : %v \n", findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) int {
	// sort greed factors
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	// sort cookies
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	childrenCount := len(g)

	childIndex := 0
	content := 0

	for _, v := range s {
		if v >= g[childIndex] {
			childIndex++
			content++

			if childIndex >= childrenCount {
				break
			}
		}
	}

	return content
}
