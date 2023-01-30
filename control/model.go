package control

// stores the information from input file into sections
type Antfarm struct {
	Roomnames []string   // name of room
	Adjacent  [][]string // list of all connecting
	Distance  int        // current position from start
	Previous  *Vertex    // previous room
	Numants   int
	Occupied  bool
	Antnames  []int    // which ant is in which room
	Roomtype  []string // type of room(start,end or middle)
}

// Graph represets an adjacency list graph
type Graph struct {
	// refers to the rooms
	vertices []*Vertex
}

// Vertex represents a graph vertex
type Vertex struct {
	// Names of the rooms
	key string
	// the connection to the two corresponding rooms
	adjacent []*Vertex
}
