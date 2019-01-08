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
