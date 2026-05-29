package main

import (
	"encoding/json"
	"fmt"
)

// Structure Personne avec des tags JSON
type Personne struct {
	Prenom string `json:"prenom,omitempty"`
	Nom    string `json:"nom,omitempty"`
	Age    int    `json:"age,omitempty"`
	Email  string `json:"email,omitempty"`
}

// Méthode qui retourne le nom complet de la personne
func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}

// Méthode qui retourne une présentation de la personne
func (p Personne) Presentation() string {
	return fmt.Sprintf(
		"Nom complet : %s\nAge : %d ans\nEmail : %s",
		p.NomComplet(),
		p.Age,
		p.Email,
	)
}

// Structure Adresse avec des tags JSON
type Adresse struct {
	Rue        string `json:"rue,omitempty"`
	Ville      string `json:"ville,omitempty"`
	CodePostal string `json:"code_postal,omitempty"`
}

// Méthode qui retourne l'adresse complète
func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

// En Go, il n'y a pas d'héritage.
// On utilise la composition par embedding : Employe contient Personne et Adresse.
type Employe struct {
	Personne
	Adresse
	Poste   string  `json:"poste,omitempty"`
	Salaire float64 `json:"-"` // Le salaire ne sera pas sérialisé en JSON
}

// Méthode qui affiche toutes les informations de l'employé
func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"--- Fiche Employé ---\n%s\nAdresse : %s\nPoste : %s\nSalaire : %.2f €",
		e.Presentation(),
		e.Adresse.Format(),
		e.Poste,
		e.Salaire,
	)
}

// Méthode avec pointeur pour modifier directement le salaire
func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire = e.Salaire + (e.Salaire * pct / 100)
}

// Structure Etudiant qui utilise aussi l'embedding avec Personne
type Etudiant struct {
	Personne
	Promo   string  `json:"promo,omitempty"`
	Moyenne float64 `json:"moyenne,omitempty"`
}

// Méthode qui retourne la mention selon la moyenne
func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "TB"
	case e.Moyenne >= 14:
		return "B"
	case e.Moyenne >= 12:
		return "AB"
	case e.Moyenne >= 10:
		return "P"
	default:
		return "Non admis"
	}
}

// Méthode qui affiche toutes les informations de l'étudiant
func (e Etudiant) FicheEtudiant() string {
	return fmt.Sprintf(
		"--- Fiche Étudiant ---\n%s\nPromo : %s\nMoyenne : %.2f\nMention : %s",
		e.Presentation(),
		e.Promo,
		e.Moyenne,
		e.MentionObtenue(),
	)
}

func main() {
	employe1 := Employe{
		Personne: Personne{
			Prenom: "Youssef",
			Nom:    "Myje",
			Age:    24,
			Email:  "youssef@gmail.com",
		},
		Adresse: Adresse{
			Rue:        "10 rue de Paris",
			Ville:      "Lyon",
			CodePostal: "69000",
		},
		Poste:   "Développeur",
		Salaire: 2800,
	}

	employe2 := Employe{
		Personne: Personne{
			Prenom: "Amine",
			Nom:    "Benali",
			Age:    30,
			Email:  "amine@gmail.com",
		},
		Adresse: Adresse{
			Rue:        "5 avenue Victor Hugo",
			Ville:      "Paris",
			CodePostal: "75000",
		},
		Poste:   "Chef de projet",
		Salaire: 3500,
	}

	etudiant1 := Etudiant{
		Personne: Personne{
			Prenom: "Sara",
			Nom:    "El Amrani",
			Age:    21,
			Email:  "sara@gmail.com",
		},
		Promo:   "Master 1",
		Moyenne: 15.5,
	}

	etudiant2 := Etudiant{
		Personne: Personne{
			Prenom: "Nabil",
			Nom:    "Ait",
			Age:    22,
			Email:  "nabil@gmail.com",
		},
		Promo:   "Bachelor 3",
		Moyenne: 17.2,
	}

	// Go gère automatiquement & et * pour les méthodes.
	// Même si AugmenterSalaire utilise *Employe, on peut écrire directement :
	employe1.AugmenterSalaire(10)

	// C'est équivalent à :
	// (&employe1).AugmenterSalaire(10)

	fmt.Println(employe1.FicheEmploye())
	fmt.Println()

	fmt.Println(employe2.FicheEmploye())
	fmt.Println()

	fmt.Println(etudiant1.FicheEtudiant())
	fmt.Println()

	fmt.Println(etudiant2.FicheEtudiant())
	fmt.Println()

	// Bonus : map pour créer un annuaire des contacts
	contacts := make(map[string]Personne)

	contacts[employe1.Email] = employe1.Personne
	contacts[employe2.Email] = employe2.Personne
	contacts[etudiant1.Email] = etudiant1.Personne
	contacts[etudiant2.Email] = etudiant2.Personne

	fmt.Println("--- Annuaire des contacts ---")

	for email, personne := range contacts {
		fmt.Println("Email :", email)
		fmt.Println("Nom complet :", personne.NomComplet())
		fmt.Println()
	}

	// Exemple de sérialisation JSON avec les struct tags
	jsonEmploye, err := json.MarshalIndent(employe1, "", "  ")

	if err != nil {
		fmt.Println("Erreur JSON :", err)
	} else {
		fmt.Println("--- Employé en JSON ---")
		fmt.Println(string(jsonEmploye))
	}
}