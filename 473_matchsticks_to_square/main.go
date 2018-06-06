package main

import (
	"fmt"
	"sort"
)

func main() {
	sticks := []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}
	fmt.Printf("Is it possible to form a square with the given sticks ? -> %v \n", makesquare(sticks))
}

func makesquare(nums []int) bool {

	if len(nums) == 0 {
		return false
	}

	perimeter := 0
	for _, v := range nums {
		perimeter += v
	}

	// perimeter should be a multiple for 4 to form a square
	if perimeter%4 != 0 {
		return false
	}

	side := perimeter / 4

	// sort the nums in decreasing order
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	if nums[0] > side {
		return false
	}

	return isOk(nums, 0, []int{0, 0, 0, 0}, side)
}

func isOk(nums []int, index int, sides []int, edge int) bool {

	if index == len(nums) {
		if sides[0] != edge || sides[1] != edge || sides[2] != edge || sides[3] != edge {
			return false
		} else {
			return true
		}
	}

	for i := 0; i < 4; i++ {
		if sides[i]+nums[index] > edge {
			continue
		} else {
			sides[i] += nums[index]
			if isOk(nums, index+1, sides, edge) {
				return true
			} else {
				sides[i] -= nums[index]
			}
		}
	}

	return false
}
