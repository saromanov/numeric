package stats

import "math"

// EuclideanDistance returns distance between two vectors
// https://en.wikipedia.org/wiki/Euclidean_distance
func EuclideanDistance(x, y []float64) (float64, error) {
	var distance float64
	for i := 0; i < len(x); i++ {
		distance = distance + ((x[i] - y[i]) * (x[i] - y[i]))
	}
	return math.Sqrt(distance), nil
}
