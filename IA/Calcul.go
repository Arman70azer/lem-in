package ia

import (
	func_cool "Lem-IN/func_utiles"
	"fmt"
)

var countFinal int

func Printeur(AllRoom *func_cool.Rooms) {
	GenererDonneesLF(AllRoom)

	// Liste des chemins disponibles pour chaque étape
	chemins := AllRoom.LF.CheminsPossible

	// Nombre de fourmis à chaque étape
	nombreFourmis := AllRoom.LF.EtapesFourmis

	//index pour chaque etapes
	indexetapes := AllRoom.LF.IndexEtapes

	var lignes []string
	var count int
	var indiceChemin int
	a := AllRoom.LF.LenOriginel

	var StockeLeschemins []string

	// Boucle pour parcourir chaque étape
	for etape := 1; etape <= 100000; etape++ {
		ligne := ""

		// Parcours de chaque étape
		for etapeActuelle := 1; etapeActuelle <= 100000; etapeActuelle++ {

			if count == a {
				break
			}
			if nombreFourmis[etapeActuelle] > 0 {

				// Chercher un chemin disponible
				cheminsDisponibles := chemins[etapeActuelle]

				indiceChemin = (indexetapes[etapeActuelle] - 1) % len(cheminsDisponibles)

				cheminDisponible := cheminsDisponibles[indiceChemin]

				if !elementDansTableau(cheminDisponible, StockeLeschemins) || nombreFourmis[etapeActuelle] == 1 {

					StockeLeschemins = append(StockeLeschemins, cheminDisponible)
					ligne += fmt.Sprintf("L%d-%s ", etapeActuelle, cheminDisponible)
					count++

					nombreFourmis[etapeActuelle]--
					indexetapes[etapeActuelle]++
				}

			}
		}
		fmt.Println(ligne)
		countFinal++

		StockeLeschemins = nil

		if ligne == "" {
			break
		}

		// Affichage de la ligne correspondant à l'étape
		lignes = append(lignes, ligne)
		a = a + AllRoom.LenOriginel
		count = 0
	}

	fmt.Println("Nombres de lignes:", countFinal-1)

}

// Remplis et créer les tableaux avec les données de AllRoom
func GenererDonneesLF(AllRoom *func_cool.Rooms) {
	// Initialisation des cartes
	chemins := make(map[int][]string)
	etapesFourmis := make(map[int]int)
	indexetapes := make(map[int]int)
	longueur := AllRoom.Nombres_fourmis
	var cheminsEtape []string
	count := 0

	//Récupère la len(AllRoom.CheminsOptimaux) avant qu'elle soit perdu
	AllRoom.LF.LenOriginel = len(AllRoom.CheminsOptimaux)

	// Remplissage des cartes en fonction de la longueur spécifiée
	for fourmis := 1; fourmis <= longueur; fourmis++ {

		if len(AllRoom.CheminsOptimaux) >= 2 {
			// Obtenez la longueur du dernier et de l'avant-dernier tableaux
			lastLength := len(AllRoom.CheminsOptimaux[len(AllRoom.CheminsOptimaux)-1])
			prevlastLength := len(AllRoom.CheminsOptimaux[len(AllRoom.CheminsOptimaux)-1])

			if fourmis == longueur-(lastLength-prevlastLength) {
				RemoveLastSlice(AllRoom)
				if count-1 == len(AllRoom.CheminsOptimaux) {
					count--
				}
			}
		}

		if count == len(AllRoom.CheminsOptimaux) {
			count = 0
		}

		cheminsEtape = append(AllRoom.CheminsOptimaux[count])
		count++

		chemins[fourmis] = cheminsEtape
		cheminsEtape = nil
		// Génération du nombre de fourmis pour cette étape (à titre d'exemple, vous pouvez ajuster selon vos besoins)
		etapesFourmis[fourmis] = len(chemins[fourmis])

		// Initialisation de l'index pour chaque étape
		indexetapes[fourmis] = 1
	}
	AllRoom.LF.CheminsPossible = chemins
	AllRoom.LF.EtapesFourmis = etapesFourmis
	AllRoom.LF.IndexEtapes = indexetapes
}

func RemoveLastSlice(AllRoom *func_cool.Rooms) {
	if len(AllRoom.CheminsOptimaux) > 0 {

		AllRoom.CheminsOptimaux = AllRoom.CheminsOptimaux[:len(AllRoom.CheminsOptimaux)-1]

	}
}

func elementDansTableau(element string, tableau []string) bool {
	for _, valeur := range tableau {
		if valeur == element {
			return true
		}
	}
	return false
}
