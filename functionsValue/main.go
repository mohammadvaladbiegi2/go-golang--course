package main

import (
	"fmt"
	"time"
)

func main() {
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(5, 10))

	//  anonymouse func
	func(a, b int) int {
		return a * b
	}(5, 5)

	// func paramert

	a := func(a int, t func()) int {
		t()
		return a + 2
	}

	a(5, func() {
		fmt.Println(time.Now().Hour())
	})
}
