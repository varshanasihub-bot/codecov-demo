package cal

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result, err := Add(2, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func TestSub(t *testing.T) {
	result, err := Sub(5, 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 3 {
		t.Errorf("expected 3, got %d", result)
	}
}

func TestMul(t *testing.T) {
	result, err := Mul(4, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 12 {
		t.Errorf("expected 12, got %d", result)
	}
}

func TestDiv(t *testing.T) {
	result, err := Div(10, 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func TestDivByZero(t *testing.T) {
	result, err := Div(10, 0)
	if err == nil {
		t.Fatal("expected an error for division by zero, got nil")
	}
	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
	expectedMsg := "error: denominator cannot be zero"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message %q, got %q", expectedMsg, err.Error())
	}
}

func TestNegativeNumbers(t *testing.T) {
	// Test Add with negative numbers
	resAdd, err := Add(-5, -3)
	if err != nil || resAdd != -8 {
		t.Errorf("expected -8, got %d (err: %v)", resAdd, err)
	}

	// Test Sub with negative numbers
	resSub, err := Sub(-5, -3)
	if err != nil || resSub != -2 {
		t.Errorf("expected -2, got %d (err: %v)", resSub, err)
	}

	// Test Mul with negative numbers
	resMul, err := Mul(-5, 3)
	if err != nil || resMul != -15 {
		t.Errorf("expected -15, got %d (err: %v)", resMul, err)
	}

	// Test Div with negative numbers
	resDiv, err := Div(-10, 2)
	if err != nil || resDiv != -5 {
		t.Errorf("expected -5, got %d (err: %v)", resDiv, err)
	}
}

// Failing test cases to demonstrate test failures and failure rate in Codecov:

func TestAddFail(t *testing.T) {
	result, err := Add(2, 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// Intentionally expecting 5 instead of 4 to cause a test failure
	if result != 5 {
		t.Errorf("FAIL INTENDED: expected 5, got %d", result)
	}
}

func TestMulFail(t *testing.T) {
	result, err := Mul(3, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// Intentionally expecting 10 instead of 9 to cause a test failure
	if result != 10 {
		t.Errorf("FAIL INTENDED: expected 10, got %d", result)
	}
}
