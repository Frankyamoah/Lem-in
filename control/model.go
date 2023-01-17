package control

// stores the information from input file into sections
type Antfarm struct {
	Numants   int
	Start     string
	End       string
	Roomnames []string
	Xcoords   []int
	Ycoords   []int
	From      []string
	To        []string
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
