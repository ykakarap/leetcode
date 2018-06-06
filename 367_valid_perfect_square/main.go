package main

import (
	"fmt"
)

func main() {

	num := 2147483647
	fmt.Printf("Number '%v'is perfect square : %v \n", num, isPerfectSquare2(num))

}

// checking by actually squaring all numbers less than n/2
func isPerfectSquare(num int) bool {
	if num == 1 || num == 0 {
		return true
	}
	for i := 1; i <= num/2; i++ {
		if i*i > num {
			return false
		}
		if i*i == num {
			return true
		}
	}

	return false
}

// sum of consequitive odd numbers is euqal to a perfect square
// using the above logic to check if we are getting 0 by subracting consequitive odd numbers
func isPerfectSquare2(num int) bool {
	for i := 1; ; i += 2 {
		num -= i
		if num == 0 {
			return true
		}
		if num < 0 {
			return false
		}
	}
}
