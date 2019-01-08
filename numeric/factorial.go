package numeric

import (
	"errors"
)

var errNegativeNumber = errors.New("number most be positive")

// Factorial provides calculation of factorial
func Factorial(x float64) (float64, error) {
	if x < 0 {
		return 0, errNegativeNumber
	}
	var (
		res float64 = 1
		i   float64 = 2
	)

	for i <= x {
		res *= i
		i++
	}
	return res, nil
}
