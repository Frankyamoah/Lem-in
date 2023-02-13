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

func newGraph() *Graph {
	return &Graph{
		rooms: make(map[int]*Room),
	}
}

func (g *Graph) addRoom(id int, name string) *Room {
	r := &Room{
		id:   id,
		name: name,
	}
	g.rooms[id] = r
	return r
}

func (g *Graph) addLink(id1, id2 int) {
	r1 := g.rooms[id1]
	r2 := g.rooms[id2]
	r1.neighbours = append(r1.neighbours, r2)
	r2.neighbours = append(r2.neighbours, r1)
}
