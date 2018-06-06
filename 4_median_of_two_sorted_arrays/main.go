package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 6}
	nums2 := []int{3, 4}

	fmt.Printf("Median is : %v \n", findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i := 0
	j := 0

	l1 := len(nums1)
	l2 := len(nums2)

	merged := make([]int, 0, l1+l2)

	for i < l1 || j < l2 {
		if i >= l1 && j < l2 {
			merged = append(merged, nums2[j:]...)
			break
		}
		if j >= l2 && i < l1 {
			merged = append(merged, nums1[i:]...)
			break
		}

		if nums1[i] <= nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	total := l1 + l2

	if total%2 == 0 {
		mid := total / 2
		return (float64(merged[mid-1]) + float64(merged[mid])) / 2
	}

	return float64(merged[total/2])
}
