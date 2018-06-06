package main

import (
	"fmt"
)

func main() {

	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Printf("Original array is : %v \n", nums)
	c := removeElement(nums, val)
	fmt.Printf("Modified array is : %v \n", nums)
	fmt.Printf("Count is : %v \n", c)
}

func removeElement(nums []int, val int) int {
	l := len(nums)
	count := l
	i := 0
	travelled := 0
	for {
		if travelled == l {
			break
		}
		v := nums[i]
		if v == val {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, v)
			count--
		} else {
			i++
		}
		travelled++
	}

	return count
}
