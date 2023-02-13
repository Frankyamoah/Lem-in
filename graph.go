package main

//create a graph structure with nodes containing
//id, room, and neighbours

type Room struct {
	id         int
	name       string
	neighbours []*Room
}

type Graph struct {
	rooms map[int]*Room
}

// newGraph creates a new Graph object and initializes its rooms map
func newGraph() *Graph {
	return &Graph{
		rooms: make(map[int]*Room),
	}
}

// addRoom adds a new room to the graph with the given id and name
func (g *Graph) addRoom(id int, name string) *Room {
	// create a new Room object with the given id and name
	r := &Room{
		id:   id,
		name: name,
	}

	// add the room to the graph's rooms map
	g.rooms[id] = r

	// return the room
	return r
}

// addLink creates a two-way link between two rooms in the graph with the given ids
func (g *Graph) addLink(id1, id2 int) {
	// retrieve the rooms with the given ids from the graph's rooms map
	r1 := g.rooms[id1]
	r2 := g.rooms[id2]

	// add each room to the other's neighbours list to create a two-way link
	r1.neighbours = append(r1.neighbours, r2)
	r2.neighbours = append(r2.neighbours, r1)
}

