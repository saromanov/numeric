package numeric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	r := GCD(100, 24)
	assert.Equal(t, r, int64(4), "numbers is not equal")
	r = GCD(24, 100)
	assert.Equal(t, r, int64(4), "numbers is not equal")
	r = GCD(5, 24)
	assert.Equal(t, r, int64(1), "numbers is not equal")
	r = GCD(5, 10)
	assert.Equal(t, r, int64(5), "numbers is not equal")
	r = GCD(5, -10)
	assert.Equal(t, r, int64(5), "numbers is not equal")
	r = GCD(-5, 10)
	assert.Equal(t, r, int64(5), "numbers is not equal")
}
