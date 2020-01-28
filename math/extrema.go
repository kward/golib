/*
Package math provides basic constants and mathematical functions that aren't
provided by default.
*/
package math

// Max returns the larger of x or y.
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

var Max = MaxInt

// Min returns the smaller of x or y.
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

var Min = MinInt
