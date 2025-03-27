package main

import "testing"

func BenchmarkPowerMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PowerMath(5, 7)
	}
}

func BenchmarkPowerSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PowerSimple(5, 7)
	}
}

// how run
// cd to file then
// go test -bench=. -benchmem
