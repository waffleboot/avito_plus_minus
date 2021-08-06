package main

import (
	"fmt"
)

func main() {
	// n := 2234
	// n := 11211
	// n := 123123123123
	// n := split(9223372036854775807)
	// n := 9223372031
	// n := 9223372031
	// fmt.Println(n)
	n := []int{9, 2, 2, 3, 1, 7, 2, 0, 3, 6, 8, 5, 0, 1, 2, 1, 4, 4, 7, 7, 5, 8, 0, 7, 1, 2, 6, 1, 8, 4, 2, 1, 1, 2, 1, 7, 2, 9, 0, 1, 1, 6, 7, 1, 5, 2, 4}
	fmt.Println(plus_minus2(n, true))
	fmt.Println(plus_minus(n))
}
