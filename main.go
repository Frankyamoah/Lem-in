package main

import (
	"fmt"
	lemin "lemin/control"
	"os"
)

func main() {
	lemin.GetAnts()
	lemin.GetRooms()
	lemin.LinkRooms()
	lemin.AssignStart()
	lemin.AssignEnd()
	lemin.AllPaths(lemin.Start)
	lemin.Final()
	lemin.Sort()
	lemin.AssignPaths()
	lemin.TraversePath(lemin.Start)
	file, _ := os.ReadFile(os.Args[1])
	fmt.Println(string(file) + "\n")
	fmt.Println(lemin.AntPath)
}
