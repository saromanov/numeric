package stats

import (
	"errors"
	"math"
)

var errEmptySlice = errors.New("slice is empty")

// CosineSimilarity provides calculation
// of cosine similarity
// https://en.wikipedia.org/wiki/Cosine_similarity
func CosineSimilarity(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0, errLengthMismatch
	}

	if len(x) == 0 || len(y) == 0 {
		return 0, errEmptySlice
	}

	var product float64
	for i := 0; i < len(x); i++ {
		product += (x[i] * y[i])
	}

	return product / (norm(x) * norm(y)), nil

}

func norm(x []float64) float64 {
	var result float64
	for i := 0; i < len(x); i++ {
		result += (x[i] * x[i])
	}
	return math.Sqrt(result)
}
