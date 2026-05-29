package main

import "fmt"

func main() {
	// Groupe de constantes avec iota pour représenter les niveaux
	const (
		NiveauFaible = iota
		NiveauMoyen
		NiveauBien
		NiveauExcellent
	)

	var prenom string
	var nombreNotes int
	var somme float64

	// Demande du prénom de l'étudiant
	fmt.Print("Entrez le prénom de l'étudiant : ")
	fmt.Scan(&prenom)

	// Demande du nombre de notes
	fmt.Print("Combien de notes voulez-vous entrer ? ")
	fmt.Scan(&nombreNotes)

	// Création d'un slice vide pour stocker les notes
	notes := []float64{}

	// Boucle for unique pour saisir les notes et calculer la somme
	for i := 0; i < nombreNotes; i++ {
		var note float64

		fmt.Printf("Entrez la note %d : ", i+1)
		fmt.Scan(&note)

		notes = append(notes, note)
		somme += note
	}

	// Calcul de la moyenne
	moyenne := somme / float64(len(notes))

	fmt.Println()
	fmt.Println("Analyse des notes de", prenom)
	fmt.Println("Notes :", notes)
	fmt.Printf("Moyenne : %.2f\n", moyenne)

	var niveau int

	// Détermination du niveau selon la moyenne
	if moyenne < 10 {
		niveau = NiveauFaible
	} else if moyenne < 14 {
		niveau = NiveauMoyen
	} else if moyenne < 17 {
		niveau = NiveauBien
	} else {
		niveau = NiveauExcellent
	}

	// Switch avec fallthrough pour afficher plusieurs messages selon le niveau
	switch niveau {
	case NiveauExcellent:
		fmt.Println("Niveau : Excellent")
		fallthrough
	case NiveauBien:
		fmt.Println("Bon travail, les résultats sont solides.")
		fallthrough
	case NiveauMoyen:
		fmt.Println("Il faut continuer les efforts.")
	case NiveauFaible:
		fmt.Println("Il faut revoir les bases.")
	default:
		fmt.Println("Niveau inconnu.")
	}
}