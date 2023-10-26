package main

import (
	Interface "Lem-IN/Crea_interface"
	ia "Lem-IN/IA"
	func_cool "Lem-IN/func_utiles"
	"time"

	"fmt"
)

func main() {

	début := time.Now()

	//Début du setup du code
	content := func_cool.OpenFile()

	AllRoom := func_cool.FoundNameAndCoordonnees(content)

	Interface.CreaSallesInterface(AllRoom)
	//Fin de la configuration

	AllRoom = ia.LesCheminsTrier(AllRoom)

	if ia.EndIsNothing == true {
		return
	}

	fmt.Println()
	fmt.Println("Nombres de fourmis:", AllRoom.Nombres_fourmis)
	fmt.Println("Etapes:", AllRoom.CheminsOptimaux)
	fmt.Println()
	ia.Printeur(&AllRoom)
	fin := time.Now()
	tempsDExecution := fin.Sub(début)

	fmt.Printf("LEM-IN a mis %v pour s'exécuter.\n", tempsDExecution)
	fmt.Println()
}
