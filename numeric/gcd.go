package numeric

// GCD provides calculation for greastest common divisor
// of two numbers
func GCD(x, y int64) int64 {
	var c int64
	for y > 0 {
		c = x % y
		x = y
		y = c
	}
	return x
}
