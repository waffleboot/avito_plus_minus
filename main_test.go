package main

import "testing"

var x string

func BenchmarkPlusMinus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = plus_minus(9223372036854775807)
	}
}

func BenchmarkPlusMinus2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = plus_minus2(9223372036854775807)
	}
}
