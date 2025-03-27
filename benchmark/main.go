package main

import (
	"math"
)

func PowerMath(x float64, y float64) float64 {
	return math.Pow(x, y)
}

func PowerSimple(x float64, y float64) float64 {
	res := float64(0)
	for i := float64(0); i < y; i++ {
		res *= i
	}

	return res
}

func main() {

}
