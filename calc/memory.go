package calc

// Variable pour stocker le résultat en mémoire
var memory float64

// Fonction pour stocker une valeur en mémoire
func StoreInMemory(value float64) {
	memory = value
}

// Fonction pour rappeler la valeur stockée en mémoire
func RecallMemory() float64 {
	return memory
}
