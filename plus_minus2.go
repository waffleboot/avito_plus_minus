package main

import (
	"fmt"
	"strings"
)

const mid = 70 // min 10

func inside(k int) bool {
	return 0 <= k && k < 2*mid
}

func plus_minus2(nums []int, verbose bool) string {

	dp := make([][]int, len(nums))
	dp[0] = make([]int, 2*mid)

	dp[0][mid+nums[0]] = 3

	for i, d := range nums[1:] {
		dp[i+1] = make([]int, 2*mid)
		for sum := -mid; sum < mid; sum++ {
			if inside(mid+sum) && dp[i][mid+sum] > 0 {
				if inside(mid + sum + d) {
					dp[i+1][mid+sum+d] = 1
				}
				if inside(mid + sum - d) {
					dp[i+1][mid+sum-d] = 1
				}
			}
		}
	}

	if dp[len(dp)-1][mid+0] == 0 {
		return "not possible"
	}

	backtrack(nums, dp)

	if verbose {
		defer func() {
			dump(dp, 2)
		}()
	}

	return path(nums, mid, dp)

}

func backtrack(nums []int, dp [][]int) {

	a := make([]int, 0, mid)
	b := make([]int, 0, mid)

	a = append(a, mid+0)
	dp[len(dp)-1][mid+0] = 2

	for i := len(nums) - 1; i > 0; i-- {
		d := nums[i]
		for _, sum := range a {
			if inside(sum-d) && dp[i-1][sum-d] == 1 {
				b = append(b, sum-d)
				dp[i-1][sum-d] = 2
			}
			if inside(sum+d) && dp[i-1][sum+d] == 1 {
				b = append(b, sum+d)
				dp[i-1][sum+d] = 2
			}
		}
		a, b = b, a[:0]
	}
}

func path(nums []int, mid int, dp [][]int) string {
	sum := nums[0]
	ans := make([]rune, len(nums)-1)
	for i := 1; i < len(dp); i++ {
		d := nums[i]
		if inside(mid+sum-d) && dp[i][mid+sum-d] == 2 {
			dp[i][mid+sum-d] = 3
			ans[i-1] = '-'
			sum -= d
		} else {
			dp[i][mid+sum+d] = 3
			ans[i-1] = '+'
			sum += d
		}
	}
	return string(ans)
}

func margin(dp [][]int, min int) (int, int) {
	l, r := len(dp[0]), 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if dp[i][j] >= min {
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

func dump(dp [][]int, min int) {
	l, r := margin(dp, min)
	if r < l {
		return
	}
	for i := 0; i < len(dp); i++ {
		var buf strings.Builder
		for j := l; j <= r; j++ {
			if dp[i][j] >= min {
				switch dp[i][j] {
				case 3:
					buf.WriteRune('+')
				case 2:
					buf.WriteRune('.')
				case 1:
					buf.WriteRune(' ')
				}
			} else {
				buf.WriteRune(' ')
			}
		}
		fmt.Println(buf.String())
	}
	fmt.Println()
}
