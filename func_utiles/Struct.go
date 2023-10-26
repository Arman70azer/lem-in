package Coord_Nom

type Rooms struct {
	Nom             []string
	Room_type       []string
	Chemins         []string
	CheminsOptimaux [][]string
	Nombres_fourmis int
	LF
}

type LF struct {
	CheminsPossible map[int][]string
	EtapesFourmis   map[int]int
	IndexEtapes     map[int]int
	LenOriginel     int
}
