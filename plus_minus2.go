package main

import (
	"fmt"
	"strings"
	"sync"
)

const mid = 70

func inside(k int) bool {
	return 0 <= k && k < 2*mid
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
