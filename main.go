package main

import "fmt"

func foo(sum int, i int, digits []int, text string) (string, bool) {
	if i >= len(digits) {
		if sum == 0 {
			return text, true
		}
		return "", false
	}
	if text, ok := foo(sum-digits[i], i+1, digits, text+"-"); ok {
		return text, true
	}
	return foo(sum+digits[i], i+1, digits, text+"+")
}

func plus_minus(n int) string {
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	for i := 0; i < len(digits)/2; i++ {
		digits[i], digits[len(digits)-i-1] = digits[len(digits)-i-1], digits[i]
	}
	if text, ok := foo(digits[0], 1, digits, ""); ok {
		return text
	}
	return "not possible"
}

func main() {
	fmt.Println(plus_minus(123))
}
