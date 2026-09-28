package cal

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	time.Sleep(15 * time.Millisecond)
	result, err := Add(2, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func TestSub(t *testing.T) {
	time.Sleep(20 * time.Millisecond)
	result, err := Sub(5, 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 3 {
		t.Errorf("expected 3, got %d", result)
	}
}

func TestMul(t *testing.T) {
	time.Sleep(25 * time.Millisecond)
	result, err := Mul(4, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 12 {
		t.Errorf("expected 12, got %d", result)
	}
}

func TestDiv(t *testing.T) {
	time.Sleep(10 * time.Millisecond)
	result, err := Div(10, 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func TestDivByZero(t *testing.T) {
	time.Sleep(12 * time.Millisecond)
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
	time.Sleep(30 * time.Millisecond)
	resAdd, err := Add(-5, -3)
	if err != nil || resAdd != -8 {
		t.Errorf("expected -8, got %d (err: %v)", resAdd, err)
	}

	resSub, err := Sub(-5, -3)
	if err != nil || resSub != -2 {
		t.Errorf("expected -2, got %d (err: %v)", resSub, err)
	}

	resMul, err := Mul(-5, 3)
	if err != nil || resMul != -15 {
		t.Errorf("expected -15, got %d (err: %v)", resMul, err)
	}

	resDiv, err := Div(-10, 2)
	if err != nil || resDiv != -5 {
		t.Errorf("expected -5, got %d (err: %v)", resDiv, err)
	}
}

func TestMod(t *testing.T) {
	time.Sleep(14 * time.Millisecond)
	res, err := Mod(10, 3)
	if err != nil || res != 1 {
		t.Errorf("expected 1, got %d (err: %v)", res, err)
	}
}

func TestModByZero(t *testing.T) {
	time.Sleep(10 * time.Millisecond)
	_, err := Mod(10, 0)
	if err == nil {
		t.Fatal("expected error for modulo by zero")
	}
}

func TestPower(t *testing.T) {
	time.Sleep(22 * time.Millisecond)
	res, err := Power(2, 4)
	if err != nil || res != 16 {
		t.Errorf("expected 16, got %d (err: %v)", res, err)
	}
}

func TestPowerNegativeExp(t *testing.T) {
	time.Sleep(11 * time.Millisecond)
	_, err := Power(2, -3)
	if err == nil {
		t.Fatal("expected error for negative exponent")
	}
}

func TestAbsolute(t *testing.T) {
	time.Sleep(16 * time.Millisecond)
	if res := Absolute(-42); res != 42 {
		t.Errorf("expected 42, got %d", res)
	}
	if res := Absolute(42); res != 42 {
		t.Errorf("expected 42, got %d", res)
	}
}

func TestIsEven(t *testing.T) {
	time.Sleep(12 * time.Millisecond)
	if !IsEven(4) {
		t.Error("expected 4 to be even")
	}
	if IsEven(5) {
		t.Error("expected 5 to be odd")
	}
}

func TestMaxAndMin(t *testing.T) {
	time.Sleep(18 * time.Millisecond)
	if max := Max(10, 20); max != 20 {
		t.Errorf("expected Max(10, 20) = 20, got %d", max)
	}
	if max := Max(30, 5); max != 30 {
		t.Errorf("expected Max(30, 5) = 30, got %d", max)
	}
	if min := Min(10, 20); min != 10 {
		t.Errorf("expected Min(10, 20) = 10, got %d", min)
	}
	if min := Min(30, 5); min != 5 {
		t.Errorf("expected Min(30, 5) = 5, got %d", min)
	}
}

// Table-driven tests for large number coverage
func TestTableDriven(t *testing.T) {
	time.Sleep(25 * time.Millisecond)
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Zero addition", 0, 0, 0},
		{"Large addition", 100000, 200000, 300000},
		{"Zero multiplication", 500, 0, 0},
		{"Identity multiplication", 500, 1, 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Add(tt.a, tt.b)
			if err != nil {
				t.Errorf("%s: unexpected error: %v", tt.name, err)
			}
			if tt.a == 500 && tt.b == 0 {
				res, _ = Mul(tt.a, tt.b)
			}
			if tt.a == 500 && tt.b == 1 {
				res, _ = Mul(tt.a, tt.b)
			}
			if res != tt.expected {
				t.Errorf("%s: expected %d, got %d", tt.name, tt.expected, res)
			}
		})
	}
}

// Failing test cases to demonstrate failure reporting in Codecov Test Analytics:

func TestAddFail(t *testing.T) {
	time.Sleep(18 * time.Millisecond)
	result, err := Add(2, 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 5 {
		t.Errorf("FAIL INTENDED: expected 5, got %d", result)
	}
}

func TestMulFail(t *testing.T) {
	time.Sleep(22 * time.Millisecond)
	result, err := Mul(3, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 10 {
		t.Errorf("FAIL INTENDED: expected 10, got %d", result)
	}
}

func TestPowerFail(t *testing.T) {
	time.Sleep(15 * time.Millisecond)
	result, err := Power(2, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// Intentionally expecting 9 instead of 8 to cause a failure
	if result != 9 {
		t.Errorf("FAIL INTENDED: expected 9, got %d", result)
	}
}
