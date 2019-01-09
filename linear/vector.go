package linear

// Vector provides definition of the vector
type Vector []float64

// Len returns size of vector
func (v Vector) Len() int {
	return len(v)
}

// Dot provides implementation of dot product
func (v Vector) Dot(p Vector) Vector {
	return v
}
