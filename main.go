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

func margin(dp [][]int) (int, int) {
	l, r := len(dp[0]), 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if dp[i][j] > 0 {
				if j < l {
					l = j
				}
				if r < j {
					r = j
				}
			}
		}
	}
	return l, r
}

func backtrack(nums []int, mid int, dp [][]int) {
	l, r := margin(dp)
	dp[len(nums)-1][mid+0] = 2
	for i := len(nums) - 2; i >= 0; i-- {
		for j := l; j <= r; j++ {
			if dp[i][j] == 1 {
				if dp[i+1][j+nums[i+1]] == 2 {
					dp[i][j] = 2
				} else if dp[i+1][j-nums[i+1]] == 2 {
					dp[i][j] = 2
				}
			}
		}
	}
}

func path(nums []int, mid int, dp [][]int) string {
	sum := nums[0]
	ans := make([]rune, len(nums)-1)
	for i := 1; i < len(nums); i++ {
		if dp[i][mid+sum-nums[i]] == 2 {
			ans[i-1] = '-'
			sum -= nums[i]
		} else {
			ans[i-1] = '+'
			sum += nums[i]
		}

	}
	return string(ans)
}

func dump(dp [][]int, all bool) {
	l, r := margin(dp)
	for i := 0; i < len(dp); i++ {
		var buf strings.Builder
		for j := l; j <= r; j++ {
			if dp[i][j] == 2 {
				buf.WriteRune('2')
			} else if dp[i][j] == 1 && all {
				buf.WriteRune('1')
			} else {
				buf.WriteRune(' ')
			}
		}
		fmt.Println(buf.String())
	}
	fmt.Println()
}

func plus_minus2(n int) string {

	const mid = 100

	nums := split(n)

	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2*mid)
	}

	dp[0][mid+nums[0]] = 1

	for i := 1; i < len(nums); i++ {
		for sum := -mid; sum < mid; sum++ {
			if dp[i-1][mid+sum] > 0 {
				dp[i][mid+sum+nums[i]] = 1
				dp[i][mid+sum-nums[i]] = 1
			}
		}
	}

	if dp[len(nums)-1][mid+0] == 0 {
		return "not possible"
	}

	backtrack(nums, mid, dp)

	return path(nums, mid, dp)

}

func main() {
	// n := 2234
	// n := 11211
	// n := 123123123123
	n := 9223372036854775807
	// n := 9223372031
	// n := 9223372031
	// fmt.Println(n)
	fmt.Println(plus_minus(n))
	fmt.Println(plus_minus2(n))
}
