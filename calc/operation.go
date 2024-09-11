package calc

import (
	"errors"
	"math"
)

// Fonction pour l'addition
func Add(a, b float64) float64 {
	return a + b
}

// Fonction pour la soustraction
func Subtract(a, b float64) float64 {
	return a - b
}

// Fonction pour la multiplication
func Multiply(a, b float64) float64 {
	return a * b
}

// Fonction pour la division avec gestion de la division par zéro
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division par zéro impossible")
	}
	return a / b, nil
}

// Fonction pour l'exponentiation
func Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}
