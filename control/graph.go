package control

import (
	"fmt"
	"strings"
)

func GetFrom(s string) string {

	var x []string

	for _, v := range s {
		if v == '-' {
			x = strings.Split(s, "-")

		}

	}
	return x[0]
}

func GetTo(s string) string {

	var x []string

	for _, v := range s {
		if v == '-' {
			x = strings.Split(s, "-")

		}

	}
	return x[1]
}

// AddVertex adds a Vertex to the graph

func (g *Graph) AddVertex(k string) {
	if Duplicate(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added as it's a dupicate key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// Add edge (link) to the graph
func (g *Graph) AddEdge(from, to []string) {

	var f string
	var t string

	for i, v := range from {
		for x, y := range to {
			if i == x {

				f = v
				t = y

				// get vertex
				fromVertex := g.getVertex(f)
				toVertex := g.getVertex(t)
				// check for errors
				if fromVertex == nil || toVertex == nil {
					err := fmt.Errorf(("invalid edge (%v-->%v"), f, t)
					fmt.Println(err.Error())
				} else if Duplicate(fromVertex.adjacent, t) {
					//err := fmt.Errorf(("existing edge (%v-->%v"), from, to)
					//	fmt.Println(err.Error())
				} else {

					fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
				}

			}
		}
	}

}

// getVertex returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex(k string) *Vertex {
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

func (g *Graph) Print() {
	for _, value := range g.vertices {
		fmt.Printf("\nRoom %v : ", value.key)

		for _, value := range value.adjacent {
			fmt.Printf("%v", value.key)
			fmt.Printf(" ")
		}
	}
	fmt.Println()

}
