package main

import (
	"fmt"
)

func main() {

	n := 1690
	fmt.Printf("%vth ugly number is : %v \n", n, nthUglyNumber(n))
}

func nextUglyNumber(uglies []int, pos2, pos3, pos5 int) (int, int, int, int) {
	n2 := uglies[pos2] * 2
	n3 := uglies[pos3] * 3
	n5 := uglies[pos5] * 5

	n := min(n2, n3, n5)

	if n == n2 {
		pos2++
	}

	if n == n3 {
		pos3++
	}

	if n == n5 {
		pos5++
	}

	return n, pos2, pos3, pos5

}

func min(a, b, c int) int {
	m := a
	if b < a {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}

func nthUglyNumber(n int) int {
	var next int
	pos2, pos3, pos5 := 0, 0, 0
	uglies := make([]int, 0, 1690)
	uglies = append(uglies, 1)
	for k := 0; k < n; k++ {
		next, pos2, pos3, pos5 = nextUglyNumber(uglies, pos2, pos3, pos5)
		uglies = append(uglies, next)
	}
	fmt.Println(len(uglies))
	return uglies[n-1]
}

func isUgly(num int) bool {
	n := num
	for n%2 == 0 {
		n = n / 2
	}

	for n%3 == 0 {
		n = n / 3
	}

	for n%5 == 0 {
		n = n / 5
	}

	if n == 1 {
		return true
	}

	return false
}
