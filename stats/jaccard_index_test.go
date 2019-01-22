package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJaccardIndex(t *testing.T) {
	x := []float64{1.1, 2.2}
	y := []float64{5.5, 0.2}
	result := JaccardIndex(x, y)
	assert.Equal(t, result, float64(0), "not expected output")

	x = []float64{}
	y = []float64{}
	result = JaccardIndex(x, y)
	assert.Equal(t, result, float64(1), "not expected output")

	x = []float64{5, 2, 3, 4}
	y = []float64{7, 0, 2, 1}
	result = JaccardIndex(x, y)
	assert.Equal(t, result, float64(0.14285714285714285), "not expected output")

}
