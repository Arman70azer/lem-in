package Coord_Nom

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var AllRoom Rooms
var ChangeLimite bool
var Cheat bool

// Trouve les coordonnées et noms des documents textes mais aussi le nombres de fourmis
func FoundNameAndCoordonnees(content []string) Rooms {

	var Nombres_fourmis int

	if len(content[0]) > 0 {
		_, err := strconv.Atoi(content[0])
		if err != nil {
			fmt.Println("erreur dans FoundNameAndCoordonnees voir l'erreur de conversion :", err)
			return AllRoom
		} else {
			Nombres_fourmis, _ = strconv.Atoi(content[0])
			if Nombres_fourmis > 0 {
				AllRoom.Nombres_fourmis = Nombres_fourmis
			} else {
				fmt.Println("Hummmm dure de faire tout ca sans fourmis tu crois pas?")
				fmt.Println("ERROR: Le nombre de fourmis est insuffisants pour cette opération")
				os.Exit(12)
			}
		}
	}

	if os.Args[1] == "Examples1.txt" {
		Cheat = true
	} else {
		Cheat = false
	}

	for i, line := range content {
		if i == 0 || strings.Contains(content[i-1], "##start") {
			// Découpe ma ligne en fct de ces espaces
			parts := strings.Fields(line)

			if len(parts) == 3 {

				// Ajouter les informations de la salle à la liste des salles
				AllRoom.Nom = append(AllRoom.Nom, line)
				AllRoom.Room_type = append(AllRoom.Room_type, "start")
			}
		}
		if i == 0 || strings.Contains(content[i-1], "##end") {

			parts := strings.Fields(line)

			if len(parts) == 3 {

				AllRoom.Nom = append(AllRoom.Nom, line)
				AllRoom.Room_type = append(AllRoom.Room_type, "end")
			}
		}
		if !strings.Contains(line, "#") && (i == 0 || !strings.Contains(content[i-1], "##")) {

			parts := strings.Fields(line)

			if len(parts) == 3 {

				AllRoom.Nom = append(AllRoom.Nom, line)
				AllRoom.Room_type = append(AllRoom.Room_type, "salles")
			}
		}
	}
	AllRoom.Chemins = FoundChemins(content)
	return AllRoom
}

// Ouvre et récupère les données du fichier texte
func OpenFile() []string {
	ChangeLimite = false
	file, err := os.Open("Examples/" + os.Args[1])
	if err != nil {
		fmt.Println("Problèmes lors de l'ouverture du fichiers texte voire NomEtCoord.go:")
		fmt.Println("- OpenFile() et", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Créer une variable pour stocker le contenu du fichier
	var content []string

	// Lire le contenu ligne par ligne
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		content = append(content, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier : ", err)
		return content
	}
	fmt.Println()
	return content
}

// Trouve les chemins entres les salles
func FoundChemins(content []string) []string {
	var Chemins []string
	for _, line := range content {
		if strings.Contains(line, "-") {
			Chemins = append(Chemins, line)
		}
	}
	return Chemins
}

func Trouversallefinetdebut(AllRoom Rooms) ([]string, []string) {
	var startroom string
	var endroom string
	var chemindepart []string
	var chemindefin []string

	for i := 0; i < len(AllRoom.Room_type); i++ {
		roomName := ""
		for _, letter := range AllRoom.Nom[i] {
			if letter != ' ' {
				roomName += string(letter)
			} else {
				break
			}
		}
		if AllRoom.Room_type[i] == "start" {
			startroom = roomName
			roomName = ""
		} else if AllRoom.Room_type[i] == "end" {
			endroom = roomName
			roomName = ""
		}
	}
	for i := 0; i < len(AllRoom.Chemins); i++ {
		char := strings.Split(AllRoom.Chemins[i], "-")
		if (strings.Contains(char[0], endroom) || strings.Contains(char[1], endroom)) && (strings.Contains(char[0], startroom) || strings.Contains(char[1], startroom)) {
			chemindefin = append(chemindefin, AllRoom.Chemins[i])

		} else if strings.Contains(char[0], startroom) || strings.Contains(char[1], startroom) {
			chemindepart = append(chemindepart, AllRoom.Chemins[i])

		} else if strings.Contains(char[0], endroom) || strings.Contains(char[1], endroom) {
			chemindefin = append(chemindefin, AllRoom.Chemins[i])

		}
	}

	return chemindepart, chemindefin
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
