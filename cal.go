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

// Mod returns the remainder of division of two integers, returning an error if denominator is zero.
func Mod(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("error: denominator cannot be zero")
	}
	return a % b, nil
}

// Power returns base raised to the exponent power.
func Power(base, exp int) (int, error) {
	if exp < 0 {
		return 0, errors.New("error: negative exponent not supported")
	}
	res := 1
	for i := 0; i < exp; i++ {
		res *= base
	}
	return res, nil
}

// Absolute returns the absolute value of an integer.
func Absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// IsEven returns true if n is even, false otherwise.
func IsEven(n int) bool {
	return n%2 == 0
}

// Max returns the maximum of two integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of two integers.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
