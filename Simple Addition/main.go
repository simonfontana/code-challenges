package main

import (
  "math/big"
  "fmt"
)

func main() {
	var term1, term2 big.Int
	fmt.Scanf("%s\n%s", &term1, &term2)
	fmt.Print(term1.Add(&term1, &term2))
}
