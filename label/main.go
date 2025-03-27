package main

import "fmt"

func main() {
	{
		sum := 10

		if sum > 0 {

			sum++

			if sum > 15 {
				goto label

			}

			fmt.Println("sum", sum)
		}

	label:
		fmt.Println("exit conditians")
	}

	i := 0

loop:
	i++
	fmt.Println("i", i)
	if i < 10 {
		goto loop
	}

	fmt.Println("exit")
}
