package stats

import (
	"errors"
	"math"
)

var errLengthMismatch = errors.New("size of slices is not equal")

// EuclideanDistance returns distance between two vectors
// https://en.wikipedia.org/wiki/Euclidean_distance
func EuclideanDistance(x, y []float64) (float64, error) {
	if err := validateDistanceInput(x, y); err != nil {
		return 0, err
	}
	var distance float64
	for i := 0; i < len(x); i++ {
		distance = distance + ((x[i] - y[i]) * (x[i] - y[i]))
	}
	return math.Sqrt(distance), nil
}

// ManhattanDistance return distance between two vectors
// https://en.wikipedia.org/wiki/Taxicab_geometry
func ManhattanDistance(x, y []float64) (float64, error) {
	if err := validateDistanceInput(x, y); err != nil {
		return 0, err
	}
	var d float64
	for i := 0; i < len(x); i++ {
		d += math.Abs(x[i] - y[i])
	}
	return d, nil
}

// validateDistanceInput returns nil if input for get distance
// between two slices is ok
func validateDistanceInput(x, y []float64) error {
	if len(x) != len(y) {
		return errLengthMismatch
	}

	return nil
}
