package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{5, 4, 3, 2, 1}
	fmt.Printf("Ranks are : %v \n", findRelativeRanks(nums))
}

type Item struct {
	Score int
	Index int
}

func findRelativeRanks(nums []int) []string {
	original := []Item{}
	l := len(nums)
	for k, v := range nums {
		original = append(original, Item{
			Score: v,
			Index: k,
		})
	}

	sort.Slice(original, func(i, j int) bool {
		return original[i].Score > original[j].Score
	})

	result := make([]string, l, l)
	for k, v := range original {
		if k == 0 {
			result[v.Index] = "Gold Medal"
			continue
		}
		if k == 1 {
			result[v.Index] = "Silver Medal"
			continue
		}
		if k == 2 {
			result[v.Index] = "Bronze Medal"
			continue
		}
		result[v.Index] = strconv.Itoa(k + 1)
	}

	return result

}
