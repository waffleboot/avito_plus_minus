package main

import (
	"fmt"
	"strings"
	"sync"
)

const mid = 70

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

func plus_minus(nums []int) string {
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

func inside(k int) bool {
	return 0 <= k && k < 2*mid
}

func backtrack(nums []int, dp [][]int) {
	dp[len(dp)-1][mid+0] = 2
	for i := len(nums) - 2; i >= 0; i-- {
		d := nums[i+1]
		for sum := 0; sum < len(dp[0]); sum++ {
			if dp[i][sum] == 1 {
				if inside(sum+d) && dp[i+1][sum+d] == 2 || inside(sum-d) && dp[i+1][sum-d] == 2 {
					dp[i][sum] = 2
				}
			}
		}
	}
}

var pathPool sync.Pool = sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}

func path(nums []int, mid int, dp [][]int) string {
	sum := nums[0]
	ans := pathPool.Get().(*strings.Builder)
	ans.Grow(len(nums) - 1)
	for i := 1; i < len(dp); i++ {
		d := nums[i]
		if inside(mid+sum-d) && dp[i][mid+sum-d] == 2 {
			ans.WriteRune('-')
			sum -= d
		} else {
			ans.WriteRune('+')
			sum += d
		}
	}
	pathPool.Put(ans)
	return ans.String()
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

func plus_minus2(nums []int) string {

	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2*mid)
	}

	dp[0][mid+nums[0]] = 1

	for i := 1; i < len(nums); i++ {
		d := nums[i]
		for sum := -mid; sum < mid; sum++ {
			if inside(mid+sum) && dp[i-1][mid+sum] > 0 {
				if inside(mid + sum + d) {
					dp[i][mid+sum+d] = 1
				}
				if inside(mid + sum - d) {
					dp[i][mid+sum-d] = 1
				}
			}
		}
	}

	if dp[len(dp)-1][mid+0] == 0 {
		return "not possible"
	}

	backtrack(nums, dp)

	return path(nums, mid, dp)

}

func main() {
	// n := 2234
	// n := 11211
	// n := 123123123123
	// n := split(9223372036854775807)
	// n := 9223372031
	// n := 9223372031
	// fmt.Println(n)
	n := []int{9, 2, 2, 3, 1, 7, 2, 0, 3, 6, 8, 5, 0, 1, 2, 1, 4, 4, 7, 7, 5, 8, 0, 7, 1, 2, 6, 1, 8, 4, 2, 1, 1, 2, 1, 7, 2, 9, 0, 1, 1, 6, 7, 1, 5, 2, 4}
	fmt.Println(plus_minus(n))
	fmt.Println(plus_minus2(n))
}
