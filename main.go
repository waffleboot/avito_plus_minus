package main

import (
	"fmt"
	"strings"
)

func split(n int) []int {
	digits := make([]int, 0, 20)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	for i := 0; i < len(digits)/2; i++ {
		digits[i], digits[len(digits)-i-1] = digits[len(digits)-i-1], digits[i]
	}
	return digits
}

func calculate(nums []int, i int, sum int, text string) (string, bool) {
	if i == len(nums) {
		if sum == 0 {
			return text, true
		}
		return "", false
	}
	if text, ok := calculate(nums, i+1, sum-nums[i], text+"-"); ok {
		return text, true
	}
	return calculate(nums, i+1, sum+nums[i], text+"+")
}

func plus_minus(n int) string {
	nums := split(n)
	if text, ok := calculate(nums, nums[0], 1, ""); ok {
		return text
	}
	return "not possible"
}

func plus_minus_2(n int) string {
	digits := split(n)
	var buf strings.Builder
	sum := digits[0]
	// fmt.Println(sum)
	for i := 1; i < len(digits); i++ {
		if sum >= 0 {
			sum -= digits[i]
			buf.WriteRune('-')
		} else {
			sum += digits[i]
			buf.WriteRune('+')
		}
		// fmt.Println(sum)
	}
	if sum == 0 {
		return buf.String()
	}
	return "not possible"
}

func main() {
	// n := 11
	// n := 123123123123
	n := 9223372036854775807
	fmt.Println(plus_minus(n))
	// fmt.Println(plus_minus_2(n))
}
