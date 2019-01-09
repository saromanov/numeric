package linear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	v := Vector([]float64{1, 2, 5})
	assert.Equal(t, v.Len(), 3, "length of vector is not equal")

	v = Vector([]float64{1})
	assert.Equal(t, v.Len(), 1, "length of vector is not equal")

	v = Vector([]float64{})
	assert.Equal(t, v.Len(), 0, "length of vector is not equal")

	v = Vector([]float64{5, 4, 7, 8, 5, 2, 3})
	assert.Equal(t, v.Len(), 7, "length of vector is not equal")
}
