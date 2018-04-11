package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0
	for i := 1; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Printf("My Sqrt(%d) is %g\n", 2, Sqrt(2))
	fmt.Printf("math.Sqrt(%d) is %g\n", 2, math.Sqrt(2))
}
