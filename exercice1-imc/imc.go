package main

import "fmt"

func main() {
	// Déclaration des variables pour stocker le poids et la taille
	var poids float64
	var taille float64

	// Déclaration des constantes pour les limites des catégories IMC
	const IMCMaigreur = 18.5
	const IMCNormal = 25.0
	const IMCSurpoids = 30.0

	// Constante contenant le prénom à afficher
	const Nom = "Youssef"

	// Message d'accueil
	fmt.Println("Bonjour", Nom)

	// Demande du poids à l'utilisateur
	fmt.Print("Entrez votre poids en kg : ")
	fmt.Scan(&poids)

	// Demande de la taille à l'utilisateur
	fmt.Print("Entrez votre taille en mètres : ")
	fmt.Scan(&taille)

	// Calcul de l'IMC avec la formule : poids / taille²
	imc := poids / (taille * taille)

	// Affichage de l'IMC avec deux chiffres après la virgule
	fmt.Printf("Votre IMC est : %.2f\n", imc)

	// Vérification de la catégorie selon la valeur de l'IMC
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