package main

type pair struct {
	i   int
	sum int
}

func plus_minus(nums []int) string {
	memo := make(map[pair]struct{})
	if text, ok := calculate(nums, 1, nums[0], "", memo); ok {
		return text
	}
	return "not possible"
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
