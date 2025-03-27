package main

import "testing"

func TestSquere(t *testing.T) {
	for i := 1; i < 10000; i++ {
		result := squere(i)

		if result/i != i {
			t.Errorf("exsept : %d,result : %d", i*i, result)
		}
	}
}
