package main

import (
	"fmt"
)

// Cette fonction reçoit deux nombres et un opérateur
// Elle retourne le résultat ou une erreur
func calculer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division par zéro impossible")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("opérateur inconnu")
	}
}

// Cette fonction crée une opération sous forme de closure
func creerOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a, b float64) float64 {
			return a + b
		}
	case "-":
		return func(a, b float64) float64 {
			return a - b
		}
	case "*":
		return func(a, b float64) float64 {
			return a * b
		}
	case "/":
		return func(a, b float64) float64 {
			return a / b
		}
	default:
		return nil
	}
}

func main() {
	var a, b float64
	var op string

	for {
		fmt.Print("Entrez le premier nombre ou 'quit' : ")
		fmt.Scan(&op)

		if op == "quit" {
			fmt.Println("Fin du programme")
			break
		}

		fmt.Sscan(op, &a)

		fmt.Print("Entrez le deuxième nombre : ")
		fmt.Scan(&b)

		fmt.Print("Entrez l'opérateur (+, -, *, /) : ")
		fmt.Scan(&op)

		resultat, err := calculer(a, b, op)

		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Println("Résultat :", resultat)
		}
	}
}