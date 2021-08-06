package main

import "testing"

var x string

var t []int = []int{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 7, 1, 2, 6, 1, 8, 4, 2, 1, 1, 2, 1, 7, 2, 9, 0, 1}

func BenchmarkPlusMinus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = plus_minus(t)
	}
}

func BenchmarkPlusMinus2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = plus_minus2(t)
	}
}
