package linear

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeros(t *testing.T) {
	res, err := Zeros(2, 2)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, res, Matrix([]Vector{
		Vector{0, 0}, Vector{0, 0},
	}), "numbers is not equal")

	res, err = Zeros(3, 2)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, res, Matrix([]Vector{
		Vector{0, 0}, Vector{0, 0}, Vector{0, 0},
	}), "numbers is not equal")

	_, err = Zeros(-3, 2)
	if err == nil {
		log.Fatal("shouldn't allow negative numbers")
	}
	_, err = Zeros(3, -2)
	if err == nil {
		log.Fatal("shouldn't allow negative numbers")
	}

}

func TestRows(t *testing.T) {
	res, err := Zeros(2, 2)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, res.Rows(), 2, "numbers is not equal")

	res, err = Zeros(3, 2)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, res.Rows(), 2, "numbers is not equal")

	res, err = Zeros(2, 7)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, res.Rows(), 7, "numbers is not equal")
}

func TestColumns(t *testing.T) {
	res, err := Zeros(2, 2)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, res.Columns(), 2, "numbers is not equal")

	res, err = Zeros(3, 2)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, res.Columns(), 3, "numbers is not equal")

	res, err = Zeros(2, 7)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, res.Columns(), 2, "numbers is not equal")
}

func TestMatrixMax(t *testing.T) {
	m := Matrix([]Vector{
		Vector{4, 2},
		Vector{6, 7},
		Vector{3, 7},
		Vector{4, 7},
		Vector{9, 7},
	})
	result, err := m.Max(1)
	assert.NoError(t, err)
	assert.Equal(t, result, 9, "numbers is not equal")
}
