package cal

import (
	"errors"
)

// Add returns the sum of two integers.
func Add(a, b int) (int, error) {
	return a + b, nil
}

// Sub returns the difference of two integers.
func Sub(a, b int) (int, error) {
	return a - b, nil
}

// Mul returns the product of two integers.
func Mul(a, b int) (int, error) {
	return a * b, nil
}

// Div returns the division result of two integers, returning an error if denominator is zero.
func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("error: denominator cannot be zero")
	}
	return a / b, nil
}
