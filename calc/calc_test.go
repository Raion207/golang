package calc

import (
	"testing"
)

func TestConvertInputToValue(t *testing.T) {
	value, err := ConvertInputToValue("5")
	if err != nil || value != 5 {
		t.Errorf("ConvertInputToValue: attendu 5, obtenu %v avec erreur %v", value, err)
	}
}

func TestCalculate(t *testing.T) {
	result, err := Calculate("+", 2, 2)
	if err != nil || result != 4 {
		t.Errorf("Calculate: attendu 4, obtenu %v avec erreur %v", result, err)
	}
}

func TestRun(t *testing.T) {
	result, err := Run("3 + 3")
	if err != nil || result != 6 {
		t.Errorf("Run: attendu 6, obtenu %v avec erreur %v", result, err)
	}
}
