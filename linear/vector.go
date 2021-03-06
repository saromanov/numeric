package linear

import (
	"errors"
	"math"
)

var errVectorMismatch = errors.New("vector mismatch")

// Vector provides definition of the vector
type Vector []float64

// Len returns size of vector
func (v Vector) Len() int {
	return len(v)
}

// Dot provides implementation of dot product
func (v Vector) Dot(p Vector) (float64, error) {
	if v.Len() != p.Len() {
		return 0, errVectorMismatch
	}

	var result float64
	for i := 0; i < v.Len(); i++ {
		result += v[i] * p[i]
	}
	return result, nil
}

// Norm calculates of the vector norm
func (v Vector) Norm() float64 {
	var norm float64
	for _, x := range v {
		norm += x * x
	}
	return math.Sqrt(norm)
}

// Max returns maximum value of the vector
func (v Vector) Max() (float64, error) {
	if len(v) == 0 {
		return 0, errors.New("vector is empty")
	}
	var result float64
	result = v[0]
	for i := 1; i < len(v); i++ {
		if v[i] > result {
			result = v[i]
		}
	}
	return result, nil
}
