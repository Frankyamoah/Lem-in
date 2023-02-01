package lemin

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// find path from start to end
var (
	roomPaths []string
	roomList  []*Room
)

// to initialise rooms with their own address
func GetRooms() {
	data, err1 := os.Open(os.Args[1])
	if err1 != nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("File Error")
		log.Fatal()
	}

	var emptyString string
	var getCoOrd string
	line := 0
	getCoOrds := bufio.NewScanner(data)
	for getCoOrds.Scan() {
		line++
		if line > 1 {
			if strings.Contains(getCoOrds.Text(), "#") {
				emptyString = ""
			} else if strings.Contains(getCoOrds.Text(), "-") {
				emptyString = ""
			} else {
				emptyString = getCoOrds.Text() + "\n"
				getCoOrd += emptyString
				emptyString = ""
			}
		}
	}
	var a []string
	var rooms *Room
	var rowInt int
	var columnInt int

	// this is to add co-ordinates to their respective room struct
	for i := 0; i < len(getCoOrd); i++ {
		if getCoOrd[i] != 10 {
			emptyString += string(getCoOrd[i])
		}
		if getCoOrd[i] == 10 {
			a = strings.Split(emptyString, " ")
			columnInt, _ = strconv.Atoi(a[1])
			rowInt, _ = strconv.Atoi(a[2])
			rooms = &Room{Name: a[0]}
			rooms.Column = columnInt
			rooms.Row = rowInt
			roomList = append(roomList, rooms)
			emptyString = ""
		}
	}
}

// this links the room to their respective next room(s)
func LinkRooms() {
	data, err1 := os.Open(os.Args[1])
	if err1 != nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("File Error")
		log.Fatal()
	}
	var emptyString string
	var links []string
	line := 0
	linksInfo := bufio.NewScanner(data)
	for linksInfo.Scan() {
		line++
		if line > 1 {
			if strings.Contains(linksInfo.Text(), "-") {
				emptyString = linksInfo.Text()
				links = append(links, emptyString)
			} else {
				emptyString = ""
			}
		}
	}

	for i := range links {
		for j := range links[i] {
			if links[i][j] == '-' {
				linkString := strings.Split(links[i], "-")
				for k := range roomList {
					for o := range roomList {
						if linkString[0] == roomList[k].Name && roomList[o].Name == linkString[1] {
							roomList[k].NextRoom = append(roomList[k].NextRoom, roomList[o])
						} else if linkString[1] == roomList[k].Name && roomList[o].Name == linkString[0] {
							roomList[k].NextRoom = append(roomList[k].NextRoom, roomList[o])
						}
					}
				}
			}
		}
	}
}

// function to assign the start room for ants.
var (
	Start    *Room
	lenStart int
)

func AssignStart() {
	data, err1 := os.Open(os.Args[1])
	if err1 != nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("File Error")
		log.Fatal()
	}
	var getStart []string
	var startLine string

	startInfo := bufio.NewScanner(data)
	for startInfo.Scan() {
		getStart = append(getStart, startInfo.Text())
	}

	for i := range getStart {
		if getStart[i] == "##start" {
			startLine = getStart[i+1]
		}
	}
	a := strings.Split(startLine, " ")
	for _, ele := range roomList {
		if ele.Name == a[0] {
			ele.Start = true
		}
	}
	for i := range roomList {
		if roomList[i].Start {
			Start = roomList[i]
		}
	}
	if Start == nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("No Start Room Found")
		log.Fatal()
	}
	lenStart = len(Start.NextRoom)
	if lenStart == 0 {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("No rooms connecting to Start")
		log.Fatal()
	}
	roomPaths = make([]string, 5)
}

// function to assign the end room for ants
var End *Room

func AssignEnd() {
	data, err1 := os.Open(os.Args[1])
	if err1 != nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("File Error")
		log.Fatal()
	}
	var getEnd []string
	var endLine string

	// this is to get coords by removing # and -
	endInfo := bufio.NewScanner(data)
	for endInfo.Scan() {
		getEnd = append(getEnd, endInfo.Text())
	}

	for i := range getEnd {
		if getEnd[i] == "##end" {
			endLine = getEnd[i+1]
		}
	}
	a := strings.Split(endLine, " ")
	for _, ele := range roomList {
		if ele.Name == a[0] {
			ele.End = true
			End = ele
		}
	}
	if End == nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("No End Room assigned")
		log.Fatal()
	}
	if End.NextRoom == nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("No rooms connecting to end")
		log.Fatal()
	}
	End.NumAnts = len(ants)
}
