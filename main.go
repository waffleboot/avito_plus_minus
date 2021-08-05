package main

import (
	"fmt"
	"strings"
)

type pair struct {
	i   int
	sum int
}

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

func calculate(nums []int, i int, sum int, text string, memo map[pair]struct{}) (string, bool) {
	if _, ok := memo[pair{i: i, sum: sum}]; ok {
		return "", false
	}
	if i == len(nums) {
		if sum == 0 {
			return text, true
		}
		return "", false
	}
	if text, ok := calculate(nums, i+1, sum-nums[i], text+"-", memo); ok {
		return text, true
	}
	if text, ok := calculate(nums, i+1, sum+nums[i], text+"+", memo); ok {
		return text, true
	}
	memo[pair{i: i, sum: sum}] = struct{}{}
	return "", false
}

func plus_minus(n int) string {
	nums := split(n)
	memo := make(map[pair]struct{})
	if text, ok := calculate(nums, 1, nums[0], "", memo); ok {
		return text
	}
	return "not possible"
}

func plus_minus2(n int) string {

	const mid = 100

	nums := split(n)

	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, 2*mid)
	}

	dp[0][mid+nums[0]] = true

	for i := 1; i < len(nums); i++ {
		for sum := -mid; sum < mid; sum++ {
			if dp[i-1][mid+sum] {
				dp[i][mid+sum+nums[i]] = true
				dp[i][mid+sum-nums[i]] = true
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		var buf strings.Builder
		for j := 0; j < len(dp[i]); j++ {
			if dp[i][j] {
				buf.WriteRune('+')
			} else {
				buf.WriteRune(' ')
			}
		}
		fmt.Println(buf.String())
	}

	if !dp[len(nums)-1][mid+0] {
		return "not possible"
	}

	var sum int
	ans := make([]rune, len(nums)-1)
	for i := len(nums) - 1; i > 0; i-- {
		if dp[i-1][mid+sum-nums[i]] {
			ans[i-1] = '+'
			sum -= nums[i]
		} else {
			ans[i-1] = '-'
			sum += nums[i]
		}
	}

	return string(ans)

}

func main() {
	// n := 2234
	// n := 11211
	// n := 123123123123
	// n := 9223372036854775807
	n := 9223372031
	fmt.Println(plus_minus(n))
	fmt.Println(plus_minus2(n))
}
