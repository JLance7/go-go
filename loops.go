package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for ; x < math.Pow(z, 2); {
		z += 0.1
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
