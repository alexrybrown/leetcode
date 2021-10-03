package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	result, original := 0, x
	for x != 0 {
		result = result*10 + (x % 10)
		x /= 10
	}

	return result == original
}

func main() {
	fmt.Println(isPalindrome(121))
}
