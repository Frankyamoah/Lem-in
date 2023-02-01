package lemin

type Room struct {
	Name     string
	Column   int
	Row      int
	NextRoom []*Room
	Start    bool
	End      bool
	Visited  int
	NumAnts  int
}

type Ant struct {
	Name     string
	Room     *Room
	PrevRoom *Room
	Path     string
}
