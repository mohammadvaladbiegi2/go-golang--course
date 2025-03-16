package main

import "fmt"

func sum(numbers ...int) int {
	var r int
	for _, v := range numbers {
		r += v
	}
	return r
}
func main() {

	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}
