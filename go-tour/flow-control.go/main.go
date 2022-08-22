package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	z := 1.0
	iter := 1
	var prev float64
	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(prev-z) < 0.000000001 {
			break
		}
		fmt.Println(z)
		prev = z
		iter = iter + 1
	}

	return z
}

func main() {
	fmt.Println(Sqrt(4))
}
