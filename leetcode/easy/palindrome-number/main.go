package main

import (
	"fmt"
)

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	div := 1

	for x >= 10*div {
		div *= 10
	}

	for x != 0 {
		right := x % 10
		left := x / div

		if left != right {
			return false
		}

		x = (x % div) / 10
		div /= 100
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(121))  // true
	fmt.Println(isPalindrome(-121)) // false
	fmt.Println(isPalindrome(10))   // false
}
