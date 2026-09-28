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
	// Tests valid division.
	// Note: We deliberately omit division by zero test here so you can observe
	// uncovered lines in Codecov! You can add it later to achieve 100% coverage.
	result, err := Div(10, 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}
