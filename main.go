package main

import (
	"fmt"
	"strings"
)

type collector struct {
	digits []int
	buf    []rune
}

func NewCollector(n int) *collector {
	c := &collector{
		digits: split(n),
	}
	c.buf = make([]rune, len(c.digits)-1)
	return c
}

func (c *collector) collect(sum int, i int) bool {
	if i+1 >= len(c.digits) {
		if sum == 0 {
			return true
		}
		return false
	}
	c.buf[i] = '-'
	if ok := c.collect(sum-c.digits[i+1], i+1); ok {
		return true
	}
	c.buf[i] = '+'
	return c.collect(sum+c.digits[i+1], i+1)
}

func split(n int) []int {
	digits := make([]int, 20)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	for i := 0; i < len(digits)/2; i++ {
		digits[i], digits[len(digits)-i-1] = digits[len(digits)-i-1], digits[i]
	}
	return digits
}

func plus_minus(n int) string {
	c := NewCollector(n)
	if ok := c.collect(c.digits[0], 0); ok {
		return string(c.buf)
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
	// n := 123
	// n := 123123123123
	n := 9223372036854775807
	fmt.Println(plus_minus(n))
	fmt.Println(plus_minus_2(n))
}
