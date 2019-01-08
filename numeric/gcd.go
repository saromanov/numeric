package numeric

// GCD provides calculation for greastest common divisor
// of two numbers
func GCD(x, y int64) int64 {
	if x < 0 {
		x = 0 - x
	}
	if y < 0 {
		y = 0 - y
	}
	var c int64
	for y > 0 {
		c = x % y
		x = y
		y = c
	}
	return x
}
