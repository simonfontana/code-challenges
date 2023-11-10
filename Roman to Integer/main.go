package main

import "fmt"

func main() {
	// test
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if translate(s[i]) < translate(s[i+1]) {
				sum -= translate(s[i])
				continue
			}
		}
		sum += translate(s[i])
	}
	return sum
}

func translate(s byte) int {
	switch s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
