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

func TestDot(t *testing.T) {
	v := Vector([]float64{1, 2, 5})
	p := Vector([]float64{2, 2, 2})
	r, err := v.Dot(p)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, r, float64(16), "result should be equal")

	v = Vector([]float64{10, -5})
	p = Vector([]float64{20, 1})
	r, err = v.Dot(p)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, r, float64(195), "result should be equal")

	v = Vector([]float64{1, 2, 5, 8, 5, 10})
	p = Vector([]float64{2, 2, 2, 3, 2, 5})
	r, err = v.Dot(p)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, r, float64(100), "result should be equal")

	v = Vector([]float64{1, 2, 5, 8, 5, 10})
	p = Vector([]float64{2})
	_, err = v.Dot(p)
	if err == nil {
		t.Fatal("vector mismatch")
	}
}

func TestNorm(t *testing.T) {
	v := Vector([]float64{1, 2, 5})
	assert.Equal(t, v.Norm(), 5.477225575051661, "norm mismatch")
}
