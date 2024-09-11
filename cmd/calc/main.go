package main

import (
	"bufio"
	"fmt"
	"os"
	"rattrapage-golang/calc"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bienvenue dans la calculatrice. Entrez une opération (ex: 2 + 2) ou tapez 'quit' pour quitter.")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1] // enlever le caractère de nouvelle ligne

		if input == "quit" {
			break
		}

		result, err := calc.Run(input)
		if err != nil {
			fmt.Println("Erreur:", err)
		} else {
			fmt.Println("Résultat:", result)
			calc.StoreInMemory(result)
		}
	}
}
