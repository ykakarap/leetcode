package main

import (
	"fmt"
	"math"
)

func main() {
	n := 55
	fmt.Printf("The minimum no.of squares needed for '%v' is : '%v' \n", n, numSquares(n))
}

var cacheMap = map[int]int{}
var sqrtCache = map[int]int{}

func numSquares(n int) int {

	if r, ok := cacheMap[n]; ok {
		return r
	}

	// edge case
	if n <= 3 {
		return n
	}

	// worst case: all 1s
	result := n

	start := getClosestRoot(n)

	for i := start; i >= 1; i-- {
		tmp := i * i
		if tmp > n {
			break
		} else {
			result = min(result, 1+numSquares(n-tmp))
		}
	}
	cacheMap[n] = result
	return result
}

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func getClosestRoot(n int) int {
	if v, ok := sqrtCache[n]; ok {
		return v
	}

	v := int(math.Sqrt(float64(n)))
	sqrtCache[n] = v
	return v
}
