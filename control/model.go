package control

import (
	"fmt"
	"strings"
)

type Antfarm struct {
	Start     string
	End       string
	Roomnames string
	Xcoords   int
	Ycoords   int
	Numants   int
}

func GetLinks(s string) string {

	var x []string

	for _, v := range s {
		if v == '-' {
			x = strings.Split(s, "-")

		}

	}
	new := strings.Join(x, ",")
	return new
}

// Graph represets an adjacency list graph
type Graph struct {
	// refers to the rooms
	vertices [] *Vertex 
}


// Vertex represents a graph vertex
type Vertex struct {
	// Names of the rooms
	key string
	// the connection to the two corresponding rooms
	adjacent []*Vertex
}

// AddVertex adds a Vertex to the graph

func (g *Graph) AddVertex(k string){
	if Duplicate(g.vertices, k){
		err := fmt.Errorf("Vertex %v not added as it's a dupicate key",k)
		fmt.Println(err.Error())
	} else {
	g.vertices = append(g.vertices, &Vertex{key:k})
	}
}

// Add edge (link) to the graph
func(g *Graph) AddEdge(from,to string){
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check for errors
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf(("Invalid edge (%v-->%v"), from, to)
		fmt.Println(err.Error())
	} else if Duplicate(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v-->%v", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}	
}

//getVertex returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex (k string) *Vertex {
	var finalVertex *Vertex
	for i, Value := range g.vertices {
		if Value.key == k {
			finalVertex = g.vertices[i]
		}
	}
	return finalVertex
}
// checks if there is duplicate room names in addVertex Function
func Duplicate(list []*Vertex, k string) bool {
	for _, value := range list {
		if k == value.key {
			return true
		}
	}
	return false
}
// Print will print the adjacent list for each vertex at the graph

func (g *Graph) Print(){
	for _, value := range g.vertices {
		fmt.Printf("\nVertex %v : ", value.key)
		for _, value := range value.adjacent{
			fmt.Printf("%v ", value.key)
		}
	}
	fmt.Println()
}