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
	// Scans text file then applies conditions to differentiate file to read coordinates 
	getCoOrds := bufio.NewScanner(data)
	for getCoOrds.Scan() {
		line++
		if line > 1 {
			//Ignores ##Start and ##End 
			if strings.Contains(getCoOrds.Text(), "#") {
				emptyString = ""
				// Ignores links between rooms	
			} else if strings.Contains(getCoOrds.Text(), "-") {
				emptyString = ""
			} else {
				// adds rest of scanned file to empty string and seperates by \n
				emptyString = getCoOrds.Text() + "\n"
				// adds empty string to getcoordinate variable
				getCoOrd += emptyString
				// restarts empty string
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
	// reads file
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
			// searches file looking for linked rooms containing -
			if strings.Contains(linksInfo.Text(), "-") {
				// places links in empty string then places in links variable
				emptyString = linksInfo.Text()
				links = append(links, emptyString)
			} else {
				// restarts empty string
				emptyString = ""
			}
		}
	}
// loops through links
	for i := range links {
		// Second Loop through each link to get to and from room
		for j := range links[i] {
			if links[i][j] == '-' {
				// splits to and from room by - 
				linkString := strings.Split(links[i], "-")
				// searches list of rooms
				for k := range roomList {
					//Second loop through list
					for o := range roomList {
						// checks if to and from room are equal to names of rooms
						if linkString[0] == roomList[k].Name && roomList[o].Name == linkString[1] {
							//creates link to next room based on the example.txt
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
// Function to determine the startroom/
func AssignStart() {
	data, err1 := os.Open(os.Args[1])
	if err1 != nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("File Error")
		log.Fatal()
	}
	var getStart []string
	var startLine string
	// gathers information for new 
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
