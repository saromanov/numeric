package stats

import (
	"testing"
)

func TestEuclideanDistance(t *testing.T) {
	x := []float64{1.1,2.2}
	y := []float64{5.5,0.2}
	_, err := EuclideanDistance(x, y)
	if err != nil {
		t.Fatal(err)
	}
}
