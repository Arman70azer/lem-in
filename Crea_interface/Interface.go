package Interface

import (
	func_utiles "Lem-IN/func_utiles"
	"fmt"
	"strconv"
	"strings"
)

// Fonction pour trouver les valeurs maximales de X et Y dans les données des salles
func findMaxXY(salles []string) (maxX, maxY int) {

	var decoupe [][]string
	var temp []string
	for _, salle := range salles {
		temp = append(temp, salle)
		if len(temp) == 1 {
			decoupe = append(decoupe, temp)
			temp = nil // Réinitialise temp pour le prochain groupe de trois valeurs
		}
	}

	for _, salle := range decoupe {
		var x int
		var y int
		Xconvertit := salle[0]
		parts := strings.Split(Xconvertit, " ")
		if len(parts) > 1 {
			Xconvertit1 := parts[1]
			x, _ = strconv.Atoi(Xconvertit1)
		}
		Yconvertit := salle[0]
		partsY := strings.Split(Yconvertit, " ")
		if len(parts) > 1 {
			Yconvertit1 := partsY[2]
			y, _ = strconv.Atoi(Yconvertit1)
		}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
	return maxX, maxY
}

// Créer le visuel de la fourmilière
func CreaSallesInterface(AllRooms func_utiles.Rooms) {

	maxX, maxY := findMaxXY(AllRooms.Nom)

	// Créer une matrice pour représenter le plan
	plan := make([][]string, maxY+1)
	for i := range plan {
		plan[i] = make([]string, maxX+1)
		for j := range plan[i] {
			plan[i][j] = "  " // Trois espaces entre les chiffres
		}
	}

	// Remplir la matrice avec les salles
	for i := 0; i < len(AllRooms.Nom); i++ {
		char := strings.Split(AllRooms.Nom[i], " ")
		x, _ := strconv.Atoi(string(char[1]))
		y, _ := strconv.Atoi(string(char[2]))

		if AllRooms.Room_type[i] == "start" {
			Salle := (char[0])
			plan[y][x] += fmt.Sprint("Start:", Salle) // Concaténer le numéro de la salle
		} else if AllRooms.Room_type[i] == "end" {
			Salle := (char[0])
			plan[y][x] += fmt.Sprint("End:", Salle)
		} else {
			Salle := (char[0])
			plan[y][x] += fmt.Sprint(Salle)
		}
	}

	// Afficher le plan
	for i := maxY; i >= 0; i-- {
		for j := 0; j <= maxX; j++ {
			fmt.Print(plan[i][j])
		}
		fmt.Println()
	}
}
