package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return	fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}


func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z:= float64(x)/2
	iterations := 0
        guess := 0.0
	for math.Abs(z-guess) >0.000000000000001 && iterations <10{
	        guess = z
		z -= (z*z - x) / (2*z)
		iterations ++
	}
	return z, nil
}
func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
