package main

import (
	"fmt"
	"sort"
)

func main() {

	points := [][]int{
		// []int{-2147483648, 2147483647},
		// []int{10, 16},
		// []int{2, 8},
		// []int{1, 6},
		// []int{7, 12},
	}

	fmt.Printf("No of arrows : %d \n", findMinArrowShots(points))
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func findMinArrowShots(points [][]int) int {
	arrows := 0

	// re arrange the points from left to right
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	start := MinInt
	end := MinInt
	for _, v := range points {
		if v[0] > end {
			// this baloon is out of the current range
			// we now need a new arrow
			arrows++
			start = v[0]
			end = v[1]
		} else {
			// adjust range as needed

			// if point start is more than range start
			// move the range start
			if v[0] > start {
				start = v[0]
			}

			if v[1] < end {
				end = v[1]
			}
		}
	}

	return arrows
}
