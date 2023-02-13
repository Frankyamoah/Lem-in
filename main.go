package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	data, err := readData()
	if err != nil {
		fmt.Println(err)
		return
	}
	printOutput(data)
}

func readData() ([]string, error) {
	if len(os.Args) != 2 {
		return nil, fmt.Errorf("ERROR: Incorrect number of arguments.\ninput format: go run . example00.txt")
	}

	path := os.Args[1]
	file, err := os.Open("examples/" + path)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Failed to open the file: %v", err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ERROR: Failed to read the file: %v", err)
	}

	return data, scanner.Err()
}

func printOutput(data []string) {
	antNbr, allRooms, allLinks := filterData(data)
	if antNbr <= 0 {
		fmt.Println("ERROR: Invalid number of ants. Must be > 0")
		return
	}

	graph := newGraph()

	//add rooms to the graph
	roomIDs := make(map[string]int)
	for id, room := range allRooms {
		roomIDs[room] = id
		graph.addRoom(id, room)
	}

	//add links to the graph
	for _, link := range allLinks {
		parts := strings.Split(link, "-")
		id1 := roomIDs[parts[0]]
		id2 := roomIDs[parts[1]]
		graph.addLink(id1, id2)
	}

	//assign start and end points
	startRoom := graph.rooms[0]
	endRoom := graph.rooms[len(graph.rooms)-1]

	paths := graph.findPaths(startRoom, endRoom)
	if len(paths) == 0 {
		fmt.Println("ERROR: no path found or text file is formatted incorrectly")
		return
	}

	validPaths := findCompatiblePaths(paths)
	bestPath := pathAssign(paths, validPaths, antNbr)

	path := os.Args[1]
	bytes, err := ioutil.ReadFile("examples/" + path)

	if err != nil {
		fmt.Println(err)
		return
	}

	content := string(bytes)
	fmt.Println(content)
	fmt.Println()
	printAntSteps(paths, bestPath)
}
