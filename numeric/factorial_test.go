package numeric

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	f, err := Factorial(8)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, f, float64(40320), "numbers is not equal")

	f, err = Factorial(5)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, f, float64(120), "numbers is not equal")

	f, err = Factorial(1)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, f, float64(1), "numbers is not equal")

	f, err = Factorial(0)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, f, float64(1), "numbers is not equal")

	_, err = Factorial(-1)
	if err == nil {
		log.Fatalf("factorial shouldn't calculate negative numbers")
	}
}
