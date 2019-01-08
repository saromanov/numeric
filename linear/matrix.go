package linear

import "errors"

var errNegativeNumber = errors.New("number should be positive")

// Matrix provides definition of the matrix
type Matrix []Vector

// Zeros defines zero matrix
func Zeros(n, m int) (Matrix, error) {
	if n < 0 || m < 0 {
		return nil, errNegativeNumber
	}

	res := make(Matrix, n)
	for i := 0; i < n; i++ {
		res[i] = make(Vector, m)
		for j := 0; j < m; j++ {
			res[i][j] = 0
		}
	}
	return res, nil
}
