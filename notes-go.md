# Notes de cours — Langage Go

## Les entiers

En Go, les entiers permettent de stocker des nombres sans virgule.
Exemples : `int`, `int8`, `int16`, `int32`, `int64`, `uint`.

`int` est le type le plus utilisé pour les nombres entiers classiques.

---

## Variables et constantes

Une variable peut changer de valeur pendant l’exécution du programme.

```go
var age int = 20
age = 21
```

Une constante ne peut pas être modifiée après sa déclaration.

```go
const PI = 3.14
```

---

## Différence entre `:=` et `=`

`:=` sert à déclarer une variable et lui donner une valeur directement.

```go
nom := "Youssef"
```

`=` sert à modifier une variable déjà existante.

```go
nom = "Ali"
```

---

## Les slices

Un slice est une liste dynamique.
Contrairement à un tableau, sa taille peut changer.

```go
notes := []int{12, 15, 18}
notes = append(notes, 20)
```

---

## `len` et `cap`

`len` donne le nombre d’éléments présents dans un slice.

```go
len(notes)
```

`cap` donne la capacité du slice, c’est-à-dire l’espace prévu en mémoire.

```go
cap(notes)
```

---

## `copy` des slices

`copy` permet de copier les éléments d’un slice dans un autre slice.

```go
source := []int{1, 2, 3}
destination := make([]int, len(source))

copy(destination, source)
```

---

## `iota`

`iota` permet de créer des constantes avec des valeurs automatiques.

```go
const (
	Faible = iota
	Moyen
	Eleve
)
```

Ici, `Faible = 0`, `Moyen = 1`, `Eleve = 2`.

---

## Maps

Une map permet de stocker des données sous forme clé / valeur.

```go
ages := map[string]int{
	"Youssef": 24,
	"Ali": 22,
}
```

On accède à une valeur grâce à sa clé.

```go
fmt.Println(ages["Youssef"])
```

---

## Boucle `for`

En Go, il n’y a qu’une seule boucle : `for`.

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

Elle peut aussi remplacer une boucle `while`.

```go
for condition {
	// code
}
```

---

## Fonctions variadiques

Une fonction variadique peut recevoir plusieurs valeurs du même type.

```go
func afficherNoms(noms ...string) {
	for _, nom := range noms {
		fmt.Println(nom)
	}
}
```

---

## Closure

Une closure est une fonction qui retourne une autre fonction ou qui garde une variable en mémoire.

```go
func compteur() func() int {
	n := 0

	return func() int {
		n++
		return n
	}
}
```

---

## `fallthrough`

`fallthrough` permet de continuer dans le cas suivant d’un `switch`.

```go
switch niveau {
case 1:
	fmt.Println("Niveau 1")
	fallthrough
case 2:
	fmt.Println("Niveau 2")
}
```

Il faut l’utiliser avec attention, car il force le passage au cas suivant.

---

## Structures

Une structure permet de regrouper plusieurs informations dans un seul type.

```go
type Personne struct {
	Nom string
	Age int
}
```

On peut ensuite créer une personne.

```go
p := Personne{Nom: "Youssef", Age: 24}
```

---

## Notion de classe en Go

Go n’a pas de classes comme Java.
À la place, Go utilise des `structs` avec des méthodes.

Cela permet de représenter des objets sans utiliser le système classique des classes.

---

## Méthodes

Une méthode permet de donner un comportement à une struct.

```go
func (p Personne) SePresenter() {
	fmt.Println("Bonjour, je suis", p.Nom)
}
```

Ici, `SePresenter()` est une méthode liée à `Personne`.

---

## Pointeurs

Un pointeur permet de travailler sur l’adresse mémoire d’une variable.

```go
func modifierAge(p *Personne) {
	p.Age = 25
}
```

L’intérêt est de modifier directement la valeur originale, sans faire de copie.

---

## Objets en Go

En Go, on peut créer des objets avec des structs et des méthodes.

```go
type Produit struct {
	Nom  string
	Prix float64
}
```

Même s’il n’y a pas de classes, on peut organiser le code de manière orientée objet.

---

## Embedding

Go n’utilise pas l’héritage classique.
À la place, il utilise la composition avec l’embedding.

```go
type Personne struct {
	Nom string
}

type Employe struct {
	Personne
	Poste string
}
```

Ici, `Employe` contient `Personne`.

---

## Struct tags

Les struct tags sont des métadonnées ajoutées aux champs d’une struct.
Ils sont souvent utilisés pour le JSON.

```go
type Personne struct {
	Nom string `json:"nom,omitempty"`
}
```

`omitempty` permet de ne pas afficher le champ s’il est vide.

```go
MotDePasse string `json:"-"`
```

`json:"-"` permet de ne pas sérialiser le champ.

---

## Visibilité en Go

En Go, la visibilité dépend de la casse.

Un nom qui commence par une majuscule est public/exporté.

```go
type Personne struct {}
```

Un nom qui commence par une minuscule est privé au package.

```go
type personne struct {}
```

---

## Camel Case et Pascal Case

Camel Case commence par une minuscule.
En Go, cela rend l’élément privé au package.

```go
nomProduit
calculerPrix
```

Pascal Case commence par une majuscule.
En Go, cela rend l’élément visible depuis un autre package.

```go
NomProduit
CalculerPrix
```

---

## `private`, `public`, `protected`

Go n’utilise pas les mots-clés `private`, `public` ou `protected`.

La visibilité se fait uniquement avec la casse :

* majuscule : public/exporté ;
* minuscule : privé au package ;
* pas de vrai `protected` comme en Java.

---

## OOP en Go

Go permet de faire de la programmation orientée objet, mais différemment de Java.

Il n’y a pas de classes ni d’héritage classique.
On utilise :

* des structs ;
* des méthodes ;
* des interfaces ;
* l’embedding.

---

## Surcharge de fonction

Go ne permet pas la surcharge de fonctions.

En Java, on peut avoir plusieurs fonctions avec le même nom mais des paramètres différents.

En Go, chaque fonction doit avoir un nom unique dans le même package.

Avantage : le code est plus simple et plus lisible.
Inconvénient : il faut parfois créer plusieurs noms de fonctions.

---

## Gestion des erreurs : Go vs Java

En Java, on utilise souvent `try / catch` pour gérer les exceptions.

```java
try {
	// code
} catch (Exception e) {
	// erreur
}
```

En Go, on retourne une erreur avec `error`.

```go
resultat, err := calculer()

if err != nil {
	fmt.Println("Erreur :", err)
}
```

Go préfère une gestion explicite des erreurs.

---

## Mot-clé `defer`

`defer` permet d’exécuter une instruction à la fin d’une fonction.

```go
defer fmt.Println("Fin du programme")
```

Il est souvent utilisé pour fermer un fichier ou une connexion.

```go
defer fichier.Close()
```
