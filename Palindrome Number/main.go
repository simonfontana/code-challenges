package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Test
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}
	return true
}
