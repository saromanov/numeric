package linear

import "errors"

var errNegativeNumber = errors.New("number should be positive")

// Matrix provides definition of the matrix
type Matrix []Vector

// Zeros defines zero matrix
func Zeros(r, c int) (Matrix, error) {
	if r < 0 || c < 0 {
		return nil, errNegativeNumber
	}

	res := make(Matrix, r)
	for i := 0; i < r; i++ {
		res[i] = make(Vector, c)
		for j := 0; j < c; j++ {
			res[i][j] = 0
		}
	}
	return res, nil
}
