package main

import (
	"fmt"
	"strings"
)

// Produit vendu dans la boutique
type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

// Catalogue avec la liste des produits
type Catalogue struct {
	Produits []Produit
}

// Ajoute un produit si l'ID n'existe pas déjà
func (c *Catalogue) AjouterProduit(p Produit) error {
	for _, produit := range c.Produits {
		if produit.ID == p.ID {
			return fmt.Errorf("un produit avec l'ID %d existe déjà", p.ID)
		}
	}

	c.Produits = append(c.Produits, p)
	return nil
}

// Cherche un produit avec son ID
func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for _, produit := range c.Produits {
		if produit.ID == id {
			return produit, nil
		}
	}

	return Produit{}, fmt.Errorf("aucun produit trouvé avec l'ID %d", id)
}

// Cherche les produits d'une catégorie
func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	resultats := []Produit{}

	for _, produit := range c.Produits {
		if strings.EqualFold(produit.Categorie, cat) {
			resultats = append(resultats, produit)
		}
	}

	return resultats
}

// Applique une réduction sur une catégorie
func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	nbModifies := 0

	for i := range c.Produits {
		if strings.EqualFold(c.Produits[i].Categorie, categorie) {
			c.Produits[i].Prix -= c.Produits[i].Prix * pct / 100
			nbModifies++
		}
	}

	return nbModifies
}

// Vend une quantité d'un produit
func (c *Catalogue) Vendre(id int, qte int) error {
	if qte <= 0 {
		return fmt.Errorf("la quantité doit être positive")
	}

	for i := range c.Produits {
		if c.Produits[i].ID == id {
			if c.Produits[i].Stock < qte {
				return fmt.Errorf("stock insuffisant pour %s", c.Produits[i].Nom)
			}

			c.Produits[i].Stock -= qte

			if c.Produits[i].Stock == 0 {
				c.Produits[i].Actif = false
			}

			return nil
		}
	}

	return fmt.Errorf("aucun produit trouvé avec l'ID %d", id)
}

// Génère un résumé du catalogue
func (c Catalogue) Rapport() string {
	var valeurTotale float64

	for _, produit := range c.Produits {
		valeurTotale += produit.Prix * float64(produit.Stock)
	}

	return fmt.Sprintf(
		"Nombre de produits : %d\nValeur totale du stock : %.2f €",
		len(c.Produits),
		valeurTotale,
	)
}

// Affiche un produit
func afficherProduit(p Produit) {
	fmt.Printf(
		"ID: %d | %s %s | %.2f € | Stock: %d | Catégorie: %s | Actif: %t\n",
		p.ID,
		p.Marque,
		p.Nom,
		p.Prix,
		p.Stock,
		p.Categorie,
		p.Actif,
	)
}

// Affiche une liste de produits
func afficherProduits(produits []Produit) {
	if len(produits) == 0 {
		fmt.Println("Aucun produit trouvé.")
		return
	}

	for _, produit := range produits {
		afficherProduit(produit)
	}
}

// Affiche les IDs disponibles
func afficherIDsProduits(c Catalogue) {
	fmt.Println("Produits disponibles :")

	for _, produit := range c.Produits {
		fmt.Printf("- ID %d : %s %s\n", produit.ID, produit.Marque, produit.Nom)
	}
}

// Affiche les catégories existantes
func afficherCategories(c Catalogue) {
	categories := []string{}

	for _, produit := range c.Produits {
		dejaExiste := false

		for _, categorie := range categories {
			if strings.EqualFold(categorie, produit.Categorie) {
				dejaExiste = true
			}
		}

		if !dejaExiste {
			categories = append(categories, produit.Categorie)
		}
	}

	fmt.Println("Catégories disponibles :")

	for _, categorie := range categories {
		fmt.Println("-", categorie)
	}
}

// Saisie d'un nouveau produit
func saisirProduit() Produit {
	var p Produit

	fmt.Print("ID : ")
	fmt.Scan(&p.ID)

	fmt.Print("Nom : ")
	fmt.Scan(&p.Nom)

	fmt.Print("Marque : ")
	fmt.Scan(&p.Marque)

	fmt.Print("Prix : ")
	fmt.Scan(&p.Prix)

	fmt.Print("Stock : ")
	fmt.Scan(&p.Stock)

	fmt.Print("Catégorie : ")
	fmt.Scan(&p.Categorie)

	p.Actif = true

	return p
}

// Produits de départ
func preRemplirCatalogue(c *Catalogue) {
	produits := []Produit{
		{ID: 1, Nom: "iPhone15", Marque: "Apple", Prix: 969.99, Stock: 10, Categorie: "Smartphone", Actif: true},
		{ID: 2, Nom: "MacBookAir", Marque: "Apple", Prix: 1299.99, Stock: 5, Categorie: "Ordinateur", Actif: true},
		{ID: 3, Nom: "GalaxyS24", Marque: "Samsung", Prix: 899.99, Stock: 8, Categorie: "Smartphone", Actif: true},
		{ID: 4, Nom: "ThinkPad", Marque: "Lenovo", Prix: 1099.99, Stock: 6, Categorie: "Ordinateur", Actif: true},
		{ID: 5, Nom: "SourisMX", Marque: "Logitech", Prix: 79.99, Stock: 20, Categorie: "Accessoire", Actif: true},
	}

	for _, produit := range produits {
		err := c.AjouterProduit(produit)

		if err != nil {
			fmt.Println("Erreur :", err)
		}
	}
}

// Menu principal
func afficherMenu() {
	fmt.Println()
	fmt.Println("===== TechShop - Boutique CLI =====")
	fmt.Println("1. Ajouter un produit")
	fmt.Println("2. Chercher un produit")
	fmt.Println("3. Soldes par catégorie")
	fmt.Println("4. Vendre un produit")
	fmt.Println("5. Rapport")
	fmt.Println("0. Quitter")
	fmt.Print("Votre choix : ")
}

func main() {
	catalogue := Catalogue{
		Produits: []Produit{},
	}

	preRemplirCatalogue(&catalogue)

	var choix int

	for {
		afficherMenu()
		fmt.Scan(&choix)

		switch choix {
		case 1:
			fmt.Println()
			fmt.Println("--- Ajouter un produit ---")

			produit := saisirProduit()
			err := catalogue.AjouterProduit(produit)

			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Produit ajouté.")
			}

		case 2:
			var sousChoix int

			fmt.Println()
			fmt.Println("--- Chercher un produit ---")
			fmt.Println("1. Par ID")
			fmt.Println("2. Par catégorie")
			fmt.Print("Votre choix : ")
			fmt.Scan(&sousChoix)

			if sousChoix == 1 {
				var id int

				afficherIDsProduits(catalogue)

				fmt.Print("ID du produit : ")
				fmt.Scan(&id)

				produit, err := catalogue.TrouverParID(id)

				if err != nil {
					fmt.Println("Erreur :", err)
				} else {
					afficherProduit(produit)
				}

			} else if sousChoix == 2 {
				var categorie string

				afficherCategories(catalogue)

				fmt.Print("Catégorie : ")
				fmt.Scan(&categorie)

				resultats := catalogue.TrouverParCategorie(categorie)
				afficherProduits(resultats)

			} else {
				fmt.Println("Choix invalide.")
			}

		case 3:
			var categorie string
			var reduction float64

			fmt.Println()
			fmt.Println("--- Soldes par catégorie ---")

			afficherCategories(catalogue)

			fmt.Print("Catégorie : ")
			fmt.Scan(&categorie)

			fmt.Print("Réduction en % : ")
			fmt.Scan(&reduction)

			nb := catalogue.AppliquerReduction(categorie, reduction)

			if nb == 0 {
				fmt.Println("Aucun produit modifié.")
			} else {
				fmt.Println(nb, "produit(s) modifié(s).")
			}

		case 4:
			var id int
			var qte int

			fmt.Println()
			fmt.Println("--- Vendre un produit ---")

			afficherIDsProduits(catalogue)

			fmt.Print("ID du produit : ")
			fmt.Scan(&id)

			fmt.Print("Quantité : ")
			fmt.Scan(&qte)

			err := catalogue.Vendre(id, qte)

			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Vente effectuée.")
			}

		case 5:
			fmt.Println()
			fmt.Println("--- Rapport ---")
			fmt.Println(catalogue.Rapport())

		case 0:
			fmt.Println("Au revoir.")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}