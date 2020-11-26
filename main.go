package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	s := float64(0)
	for i:=0; i < 10; i ++ {
		z -= (z*z - x) / (2*z)
		//1e-10は1×10の-10乗
		if math.Abs(z - s) < 1e-10 {
			break;
		}
		s = z
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}