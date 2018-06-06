package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 3}
	fmt.Printf("Output is : %v \n", majorityElement(nums))

}

func majorityElement(nums []int) []int {
	presence := map[int]int{}
	results := []int{}
	nby3 := len(nums) / 3

	for _, v := range nums {
		if _, ok := presence[v]; ok {
			// update the presence count
			presence[v] = presence[v] + 1
		} else {
			presence[v] = 1

		}
		if presence[v] == nby3+1 {
			results = append(results, v)
		}
	}

	return results
}
