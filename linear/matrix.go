package linear

import "errors"

var errNegativeNumber = errors.New("number should be positive")

// Matrix provides definition of the matrix
type Matrix []Vector

// Zeros defines zero matrix
func Zeros(c, r int) (Matrix, error) {
	if r < 0 || c < 0 {
		return nil, errNegativeNumber
	}

	res := make(Matrix, c)
	for i := 0; i < c; i++ {
		res[i] = make(Vector, r)
		for j := 0; j < r; j++ {
			res[i][j] = 0
		}
	}
	return res, nil
}

// Columns returns number of columns on matrix
func (m Matrix) Columns() int {
	return len(m)
}

// Rows return number of rows on matrix
func (m Matrix) Rows() int {
	return len(m[0])
}

// Multiply provides multiply of two matrix
// A x B = (m x n) x (n x m), where n, m are integers who define
// * the dimensions of matrices A, B.
func (m Matrix) Multiply(n Matrix) (Matrix, error) {
	return m, nil
}

// Max return maximum value along the axis
func (m Matrix) Max(axis int) (float64, error) {
	if axis == 1 {
		var result float64
		for _, rows := range m {
			r, err := rows.Max()
			if err != nil {
				return 0, err
			}
			if r > result {
				result = r
			}
		}
		return result, nil
	}
	return 0, nil
}
