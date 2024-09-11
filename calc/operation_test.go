package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add: attendu 5, obtenu %v", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	if result != 2 {
		t.Errorf("Subtract: attendu 2, obtenu %v", result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 4)
	if result != 12 {
		t.Errorf("Multiply: attendu 12, obtenu %v", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil || result != 5 {
		t.Errorf("Divide: attendu 5, obtenu %v avec erreur %v", result, err)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("Divide: division par zéro n'a pas généré d'erreur")
	}
}

func TestPower(t *testing.T) {
	result := Power(2, 3)
	if result != 8 {
		t.Errorf("Power: attendu 8, obtenu %v", result)
	}
}
