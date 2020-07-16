package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64{
	z:= float64(x)/2
	iterations := 0
        guess := 0.0
	for math.Abs(z-guess) >0.000000000000001 && iterations <10{
	        guess = z
		z -= (z*z - x) / (2*z) 
		iterations ++
	        fmt.Printf("Guess is %f and there has been %d iterations\n",z,iterations)
	}
	return z

}


func main(){
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(1))
}
