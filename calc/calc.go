package calc

import (
	"fmt"
	"strconv"
	"strings"
)

// Fonction principale pour calculer une opération basée sur l'entrée utilisateur
func Calculate(operation string, a, b float64) (float64, error) {
	switch operation {
	case "+":
		return Add(a, b), nil
	case "-":
		return Subtract(a, b), nil
	case "*":
		return Multiply(a, b), nil
	case "/":
		return Divide(a, b)
	case "^":
		return Power(a, b), nil
	default:
		return 0, fmt.Errorf("opération inconnue : %s", operation)
	}
}

// Fonction pour convertir une entrée utilisateur en commande calculable
func ConvertInputToValue(input string) (float64, error) {
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("entrée invalide : %s", input)
	}
	return value, nil
}

// Fonction pour exécuter le calcul basé sur une chaîne d'entrée comme "2 + 3"
func Run(input string) (float64, error) {
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		return 0, fmt.Errorf("format d'entrée invalide")
	}

	a, err := ConvertInputToValue(parts[0])
	if err != nil {
		return 0, err
	}

	b, err := ConvertInputToValue(parts[2])
	if err != nil {
		return 0, err
	}

	return Calculate(parts[1], a, b)
}
