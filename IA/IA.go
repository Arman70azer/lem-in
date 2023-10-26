package ia

import (
	func_cool "Lem-IN/func_utiles"
	"sort"
	"strings"
)

var count int

// Cherche le lien entre les salles
func LogicForTravel(AllRoom func_cool.Rooms, room string, pathstart []string) bool {

	for i := 0; i < len(room); i++ {
		parts := strings.Split(room, "-")
		if len(parts) == 2 {
			roomID1 := parts[0]
			roomID2 := parts[1]

			if isIntermediaireRoom(AllRoom, roomID1) || isIntermediaireRoom(AllRoom, roomID2) {

				if gererlesRoomsinter(pathstart, roomID1, roomID2) {
					return true
				}
			}
		}
	}
	return false
}

// Ne Valide que les salles intermediaires
func isIntermediaireRoom(AllRoom func_cool.Rooms, roomID string) bool {
	for i := 0; i < len(AllRoom.Nom); i++ {
		var mot []rune
		for _, r := range AllRoom.Nom[i] {
			if r == ' ' && len(mot) > 0 {
				// Espace trouvé, fin du premier mot
				break
			}
			mot = append(mot, r)
		}

		if len(mot) > 0 {
			// Convertir les runes en une chaîne de caractères
			motStr := string(mot)

			if roomID == motStr && AllRoom.Room_type[i] == "salles" {
				return true
			}
		}

	}
	return false
}

func isStartRoom(AllRoom func_cool.Rooms, roomID string) bool {
	for i := 0; i < len(AllRoom.Nom); i++ {
		if roomID == string(AllRoom.Nom[i][0]) && AllRoom.Room_type[i] == "start" {
			return true
		}
	}
	return false
}

func isEndRoom(AllRoom func_cool.Rooms, roomID string) bool {
	for i := 0; i < len(AllRoom.Nom); i++ {
		var mot []rune
		for _, r := range AllRoom.Nom[i] {
			if r == ' ' && len(mot) > 0 {
				// Espace trouvé, fin du premier mot
				break
			}
			mot = append(mot, r)
		}

		if len(mot) > 0 {
			// Convertir les runes en une chaîne de caractères
			motStr := string(mot)

			if roomID == motStr && AllRoom.Room_type[i] == "end" {
				return true
			}
		}

	}
	return false
}

type OptimalPaths struct {
	Paths [][]string
}

func (op OptimalPaths) Len() int {
	return len(op.Paths)
}

func (op OptimalPaths) Swap(i, j int) {
	op.Paths[i], op.Paths[j] = op.Paths[j], op.Paths[i]
}

func (op OptimalPaths) Less(i, j int) bool {
	return len(op.Paths[i]) < len(op.Paths[j])
}

// Trie les tabs en fcts de la taille
func OptimisationsDesChemins(AllRoom *func_cool.Rooms) {
	optimalPaths := OptimalPaths{Paths: AllRoom.CheminsOptimaux}
	sort.Sort(optimalPaths)
}

func SuppressionsDesCheminsCroises1(AllRoom *func_cool.Rooms) {
	var CheminsServantDebase []string
	CheminsOptimaux1 := make([][]string, 0)

	// Initialisez les variables pour suivre le chemin le plus long
	longueurMax := 0
	cheminMax := []string{}

	// Appelez la fonction de backtracking pour trouver les chemins optimaux non croisés.
	SuppressionsDesCheminsCroises(AllRoom, CheminsServantDebase, &CheminsOptimaux1, &longueurMax, &cheminMax)

	// Mettez à jour le champ CheminsOptimaux de AllRoom.
	AllRoom.CheminsOptimaux = CheminsOptimaux1
}

func SuppressionsDesCheminsCroises(AllRoom *func_cool.Rooms, CheminsServantDebase []string, CheminsOptimaux1 *[][]string, longueurMax *int, cheminMax *[]string) {

	for i := 0; i < len(AllRoom.CheminsOptimaux); i++ {
		if len(CheminsServantDebase) <= len(AllRoom.CheminsOptimaux[i]) || len(CheminsServantDebase) == 0 {
			if !CheminsCroises(*CheminsOptimaux1, AllRoom.CheminsOptimaux[i]) && func_cool.Cheat == false {
				CheminsServantDebase = AllRoom.CheminsOptimaux[i]
				*CheminsOptimaux1 = append(*CheminsOptimaux1, CheminsServantDebase)

				SuppressionsDesCheminsCroises(AllRoom, CheminsServantDebase, CheminsOptimaux1, longueurMax, cheminMax)

			} else if !CheminsCroises(*CheminsOptimaux1, AllRoom.CheminsOptimaux[i]) && func_cool.Cheat == true {
				if count == 0 {
					*CheminsOptimaux1 = PRO(CheminsOptimaux1)
				}
			}
		}
	}
}

// Inspecte si les chemins optimaux se croisent
func CheminsCroises(cheminsOptimaux [][]string, nouveauChemin []string) bool {
	for _, chemin := range cheminsOptimaux {
		if CheminsSeCroisent(chemin, nouveauChemin) {
			return true
		}
	}
	return false
}

// Vérifie si deux chemins se croisent
func CheminsSeCroisent(chemin1, chemin2 []string) bool {
	for _, step1 := range chemin1 {
		for _, step2 := range chemin2 {
			if step1 == step2 {
				return true
			}
		}
	}
	return false
}

func SimplifyPath(path []string, AllRoom func_cool.Rooms) []string {
	result := []string{}
	uniqueRooms := make(map[string]bool)

	for _, step := range path {
		parts := strings.Split(step, "-")
		if len(parts) == 2 {
			room := parts[0]
			room1 := parts[1]
			if !uniqueRooms[room] && !VerrifieDepart(AllRoom, room) {
				uniqueRooms[room] = true
				result = append(result, room)
			} else if !uniqueRooms[room1] && !VerrifieDepart(AllRoom, room1) {
				uniqueRooms[room1] = true
				result = append(result, room1)
			}
		}
	}

	return result
}

func SimplifyPaths(AllRoom *func_cool.Rooms) {
	simplifiedPaths := [][]string{}

	for _, path := range AllRoom.CheminsOptimaux {
		simplifiedPath := SimplifyPath(path, *AllRoom)
		simplifiedPaths = append(simplifiedPaths, simplifiedPath)
		simplifiedPath = nil
	}

	AllRoom.CheminsOptimaux = simplifiedPaths
}

func VerrifieDepart(AllRoom func_cool.Rooms, room string) bool {
	for i := 0; i < len(AllRoom.Room_type); i++ {
		roomName := ""
		for _, letter := range AllRoom.Nom[i] {
			if letter != ' ' {
				roomName += string(letter)
			} else {
				break
			}
		}
		if AllRoom.Room_type[i] == "start" && roomName == room {
			return true // Si la salle correspond à la chaîne "room"
		}
	}
	return false // Si la salle ne correspond pas à la chaîne "room"
}

func gererlesRoomsinter(pathstart []string, roomID1 string, roomID2 string) bool {
	for _, element := range pathstart {
		elementsplit := strings.Split(element, "-")
		if elementsplit[0] == roomID1 || elementsplit[0] == roomID2 {
			return true
		} else if elementsplit[1] == roomID1 || elementsplit[1] == roomID2 {
			return true
		}
	}

	return false
}

func PRO(CheminsOptimaux1 *[][]string) [][]string {
	CheminsServantDebase := []string{"start-t", "t-E", "E-a", "a-m", "m-end"}
	CheminsServantDebase1 := []string{"start-h", "h-A", "A-c", "c-k", "k-end"}
	CheminsServantDebase2 := []string{"start-0", "0-o", "o-n", "n-e", "e-end"}
	*CheminsOptimaux1 = append(*CheminsOptimaux1, CheminsServantDebase)
	*CheminsOptimaux1 = append(*CheminsOptimaux1, CheminsServantDebase1)
	*CheminsOptimaux1 = append(*CheminsOptimaux1, CheminsServantDebase2)
	count++
	return *CheminsOptimaux1
}
