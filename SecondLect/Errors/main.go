package main

import (
	"fmt"
)

type ErrNegativeNum float64

func (e ErrNegativeNum) Error() string{
	return fmt.Sprintf("Cannot Sqrt negative number %v", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeNum(x)
	}
	z := 1.0
	z -= (z*z - x) / (2*z)
	for z - (z - ((z*z - x) / (2*z))) > 1e-15 {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
