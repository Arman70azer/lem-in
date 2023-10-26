package ia

import (
	func_cool "Lem-IN/func_utiles"
	"fmt"
	"os"
	"strings"
)

var ChangeLimite bool
var EndIsNothing bool

// Trouve tout les chemins possibles
func LesCheminsTrier(AllRoom func_cool.Rooms) func_cool.Rooms {
	var startRooms []string
	var endroom []string
	startRooms, endroom = func_cool.Trouversallefinetdebut(AllRoom)
	EndIsNothing = true

	for a := 0; a < len(endroom); a++ {
		split := strings.Split(endroom[a], "-")
		if isEndRoom(AllRoom, split[0]) || isEndRoom(AllRoom, split[1]) {
			EndIsNothing = false
		}
	}

	if EndIsNothing == true {
		fmt.Println()
		fmt.Println("C'est la fin... car il n'y a pas de fin mdr : Aucun chemins n'as de salle de type <fin>")
		os.Exit(12)
	}

	//Créer des chemins en fct du nombre de salles de départs et ainsi créer des chemins pour toutes les salles de départs
	for _, startRoom := range startRooms {
		usedRooms := make(map[string]bool)
		pathstart := []string{startRoom}
		AllRoom = backtrack(pathstart, endroom, AllRoom.Chemins, usedRooms, startRooms)
	}
	CheminFinEtStart(&AllRoom)

	OptimisationsDesChemins(&AllRoom)

	SuppressionsDesCheminsCroises1(&AllRoom)

	fmt.Println()
	fmt.Println("Chemins possibles sans croissements :", AllRoom.CheminsOptimaux)
	SimplifyPaths(&AllRoom)

	return AllRoom
}

// Va chercher tout les chemins possibles c'est la func la plus importante !!!
func backtrack(pathstart []string, endRooms []string, rooms []string, used map[string]bool, startRooms []string) func_cool.Rooms {
	EndIsNothing = false
	if len(pathstart) >= 1 {
		for _, endRoom := range endRooms {
			if pathstart[len(pathstart)-1] == endRoom {
				// Si le chemin a atteint "end", ajoutez-le à la liste des chemins
				func_cool.AllRoom.CheminsOptimaux = append(func_cool.AllRoom.CheminsOptimaux, append([]string(nil), pathstart...))
				return func_cool.AllRoom
			} else if pathstart[0] == endRoom {

				func_cool.AllRoom.CheminsOptimaux = append(func_cool.AllRoom.CheminsOptimaux, append([]string(nil), pathstart...))
				return func_cool.AllRoom

			}
		}
	}

	lastRoom := pathstart[len(pathstart)-1]

	for _, room := range rooms {
		// Vérifiez si la salle n'a pas déjà été utilisée et qu'elle n'est pas déjà dans le chemin
		if !used[room] && !contains(startRooms, room) && room != lastRoom {

			if LogicForTravel(func_cool.AllRoom, room, pathstart) {
				// Ajoutez la salle au chemin
				if len(pathstart) < 7 && ChangeLimite == false {
					pathstart = append(pathstart, room)
					used[room] = true

					// Récursivement, explorez le chemin suivant
					backtrack(pathstart, endRooms, rooms, used, startRooms)

					// Retirez la salle du chemin pour explorer d'autres possibilités
					pathstart = pathstart[:len(pathstart)-1]
					used[room] = false
				} else if ChangeLimite == true {
					pathstart = append(pathstart, room)
					used[room] = true

					// Récursivement, explorez le chemin suivant
					backtrack(pathstart, endRooms, rooms, used, startRooms)

					// Retirez la salle du chemin pour explorer d'autres possibilités
					pathstart = pathstart[:len(pathstart)-1]
					used[room] = false
				}

			}
		}
	}

	return func_cool.AllRoom
}

// Voit si il contient une salle spécifique (ici les salles de départ)
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func CheminFinEtStart(AllRoom *func_cool.Rooms) {
	for i := 0; i < len(AllRoom.Chemins); i++ {

		split := strings.Split(AllRoom.Chemins[i], "-")

		if (isStartRoom(func_cool.AllRoom, split[0]) || isStartRoom(func_cool.AllRoom, split[1])) && (isEndRoom(func_cool.AllRoom, split[0]) || isEndRoom(func_cool.AllRoom, split[1])) {

			spliteur := strings.Split(AllRoom.Chemins[i], " ")
			AllRoom.CheminsOptimaux = append(AllRoom.CheminsOptimaux, spliteur)
		}

	}
}
