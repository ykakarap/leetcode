package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Printf("Next permutation is  : %v \n", nums)
}

func nextPermutation(nums []int) {

	// edge case
	// if the nums are is proper descending order then returning the number is propert ascending order is the answer
	if sort.SliceIsSorted(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	}) {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		return
	}

	l := len(nums)
	point := 0
	for k, v := range nums {
		if k+1 != l {
			if v < nums[k+1] {
				point = k
			}
		}
	}

	// sort after point
	partial := nums[point+1:]
	sort.Slice(partial, func(i, j int) bool {
		return partial[i] < partial[j]
	})

	// swapper point
	var swapperPoint int
	for k, v := range partial {
		if nums[point] < v {
			swapperPoint = k
			break
		}
	}

	tmp := partial[swapperPoint]
	partial[swapperPoint] = nums[point]
	nums[point] = tmp
}
