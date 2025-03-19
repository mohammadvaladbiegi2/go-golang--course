package main

import "fmt"

type Parametr struct {
	num1 int
	num2 int
}

func main() {
	firstnum := Parametr{
		num1: 6,
		num2: 7,
	}

	fmt.Println(firstnum.add())
	fmt.Println(firstnum.minus())
	fmt.Println(firstnum.power())
	fmt.Println(firstnum.devide())

}

func (num Parametr) add() int {
	return num.num1 + num.num2
}
func (num Parametr) minus() int {
	return num.num1 - num.num2
}
func (num Parametr) power() int {
	return num.num1 * num.num2
}
func (num Parametr) devide() int {
	return num.num1 / num.num2
}
