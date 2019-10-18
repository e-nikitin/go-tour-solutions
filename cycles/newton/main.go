package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	delta := 1.0
	for math.Abs(delta) > 0.0000000001 {
		delta = (z*z - x) / (2 * z)
		fmt.Println(delta)
		z -= delta
	}
	return z
}

func main() {
	fmt.Println(sqrt(4))
}
