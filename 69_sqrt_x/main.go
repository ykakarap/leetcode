package main

import (
	"fmt"
)

func main() {
	x := 12
	fmt.Printf("Sqrt of %v is : %v \n", x, mySqrt(x))
	// e := 0.000001
	// fx := float64(x)
	// fmt.Printf("Sqrt of %v accuarte upto 5 decimals : %v \n", fx, babylonian(fx, e))
}

func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}

	return sqrt(x, 1, x)
}

func sqrt(x, start, end int) int {
	var m, m2, ans int
	for start <= end {
		m = (start + end) / 2
		m2 = m * m

		if m2 == x {
			ans = m
			break
		}

		if m2 > x {
			end = m - 1
		} else {
			ans = m
			start = m + 1
		}
	}

	return ans
}

func babylonian(n float64, e float64) float64 {
	var x, y float64
	x = n / 2
	y = 1
	for x-y > e {
		x = (x + y) / 2
		y = n / x
	}
	return x
}
