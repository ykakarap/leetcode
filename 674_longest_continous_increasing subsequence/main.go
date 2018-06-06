package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 4, 2, 3, 4, 5}
	fmt.Printf("Answer  : %v \n", findLengthOfLCIS(nums))
}

func findLengthOfLCIS(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	current := 1
	best := 1
	for k, v := range nums {
		if k == 0 {
			continue
		}
		if v > nums[k-1] {
			current++
		} else {
			if current > best {
				best = current
			}
			current = 1
		}
	}

	if current > best {
		best = current
	}

	return best

}
