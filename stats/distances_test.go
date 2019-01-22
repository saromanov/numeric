package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEuclideanDistance(t *testing.T) {
	x := []float64{1.1, 2.2}
	y := []float64{5.5, 0.2}
	result, err := EuclideanDistance(x, y)
	assert.NoError(t, err)
	assert.Equal(t, result, float64(4.833218389437829), "not expected output")

}
