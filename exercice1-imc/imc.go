package main

import "fmt"

func main() {
	var poids float64
	var taille float64

	const IMCMaigreur = 18.5
	const IMCNormal = 25.0
	const IMCSurpoids = 30.0
	const Nom = "Youssef"

	fmt.Println("Bonjour", Nom)
	fmt.Print("Entrez votre poids en kg : ")
	fmt.Scan(&poids)

	fmt.Print("Entrez votre taille en mètres : ")
	fmt.Scan(&taille)

	imc := poids / (taille * taille)

	fmt.Printf("Votre IMC est : %.2f\n", imc)

	if imc < IMCMaigreur {
		fmt.Println("Catégorie : Maigreur")
	} else if imc < IMCNormal {
		fmt.Println("Catégorie : Normal")
	} else if imc < IMCSurpoids {
		fmt.Println("Catégorie : Surpoids")
	} else {
		fmt.Println("Catégorie : Obésité")
	}
}