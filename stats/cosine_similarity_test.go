package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCosineSimilarity(t *testing.T) {
	x := []float64{1.1, 1.1}
	y := []float64{1.1, 1.1}
	result, err := CosineSimilarity(x, y)
	assert.NoError(t, err)
	assert.Equal(t, result, float64(1.0000000000000002), "not expected output")
}
