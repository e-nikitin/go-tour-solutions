package main

import (
	"fmt"
	"math"
)

const delta = 1e-10

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Sqrt: negative number %g", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	z := x
	
	for {
		n := z - (z*z-x)/(2*z)
		
		if math.Abs(n-z) < delta {
			break
		}
		z = n
	}
	
	return z, nil
	
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
