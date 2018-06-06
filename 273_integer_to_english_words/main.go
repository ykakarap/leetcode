package main

import (
	"fmt"
)

func main() {
	num := 212456000
	fmt.Printf("Number %v is : '%v' \n", num, numberToWords(num))
}

var (
	singles = map[int]string{
		0: "",
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
	}

	teens = map[int]string{
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
	}

	tens = map[int]string{
		2: "Twenty",
		3: "Thirty",
		4: "Forty",
		5: "Fifty",
		6: "Sixty",
		7: "Seventy",
		8: "Eighty",
		9: "Ninety",
	}
)

func numberToWords(num int) string {

	if num == 0 {
		return "Zero"
	}

	var result string

	// first 3 digits
	hundereds := convert3Digits(num % 1000)
	if hundereds != "" {
		result = hundereds
	}

	thousands := convert3Digits((num % 1000000) / 1000)
	if thousands != "" {
		if result != "" {
			result = " Thousand " + result
		} else {
			result = " Thousand"
		}
		result = thousands + result
	}

	millions := convert3Digits((num % 1000000000) / 1000000)
	if millions != "" {
		if result != "" {
			result = " Million " + result
		} else {
			result = " Million"
		}
		result = millions + result
	}

	billions := convert3Digits((num % 1000000000000) / 1000000000)
	if billions != "" {
		if result != "" {
			result = " Billion " + result
		} else {
			result = " Billion"
		}
		result = billions + result
	}

	return result
}

func convert3Digits(num int) string {

	if num == 0 {
		return ""
	}

	var lhs, rhs string

	// get middle number
	m := (num % 100) / 10
	if m == 1 {
		rhs = teens[num%100]
	} else if m == 0 {
		rhs = singles[num%10]
	} else {
		s := singles[num%10]
		if s != "" {
			rhs = tens[m] + " " + s
		} else {
			rhs = tens[m]
		}
	}

	// get the first number
	f := num / 100
	if f > 0 {
		lhs = singles[f]
	}

	if lhs != "" {
		if rhs != "" {
			return lhs + " Hundred " + rhs
		} else {
			return lhs + " Hundred"
		}
	} else {
		return rhs
	}
}
