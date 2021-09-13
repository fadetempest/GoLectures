package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	z -= (z*z - x) / (2*z)
	count := 1
	for z - (z - ((z*z - x) / (2*z))) > 1e-15 {
		z -= (z*z - x) / (2 * z)
		fmt.Println(count, "try", z)
		count++
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
