package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1}
	k := 1
	fmt.Printf("No.of subarrays is : %v \n", numSubarrayProductLessThanK(nums, k))

}

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	prod := 1
	ans, left := 0, 0
	for right, v := range nums {
		fmt.Printf("Right : %v \n", right)
		prod *= v
		for prod >= k {
			fmt.Printf("Left : %v \n", left)
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}
