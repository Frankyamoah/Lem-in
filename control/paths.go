package lemin

import "strings"

// find path from start to end
var count int

func AllPaths(r *Room) {
	prevRoom := r
	nextRoom := r.NextRoom
	visitedCounter := 0
	lenCounter := 0
	if prevRoom.End {
		roomPaths[count] += prevRoom.Name
		roomPaths[count] = StartEnd(roomPaths[count])
		Start.Visited = 0
		End.Visited = 0
		count++
		AllPaths(Start)
	}
	for _, ele := range nextRoom {
		if ele.Visited == 1 {
			visitedCounter++
		}
	}

	if visitedCounter == len(nextRoom) {
		prevRoom.Visited = 1
		dEndNameSlice := strings.Split(roomPaths[count], ",")
		for _, room := range roomList {
			if len(dEndNameSlice) >= 2 {
				if dEndNameSlice[len(dEndNameSlice)-2] == room.Name {
					dEndNameSlice = Remove(dEndNameSlice, len(dEndNameSlice)-2)
					roomPaths[count] = strings.Join(dEndNameSlice, ",")
				}
			}
		}
	} else {
		for _, roomele := range nextRoom {
			lenCounter++
			if prevRoom == Start {
				for i, rooms := range nextRoom {
					if count < lenStart {
						rNamesSlice := strings.Split(roomPaths[count], ",")
						if !Contains(rNamesSlice, rooms.Name) && (rooms.Visited == 0) && !strings.HasPrefix(rooms.Name, "G") {
							roomPaths[count] += prevRoom.Name + ","
							prevRoom.Visited = 1
							AllPaths(nextRoom[i])
						}
					}
				}
			}
			if roomele.End {
				if count < lenStart {
					roomPaths[count] += prevRoom.Name + ","
					prevRoom.Visited = 1
					AllPaths(roomele)
				}
			} else if lenCounter == len(nextRoom) {
				for _, check := range nextRoom {
					for _, endNextRooms := range End.NextRoom {
						if check.Name == endNextRooms.Name && check.Visited == 0 {
							if count < lenStart {
								rNamesSlice := strings.Split(roomPaths[count], ",")
								if !Contains(rNamesSlice, check.Name) && !strings.HasPrefix(check.Name, "G") {
									if prevRoom != End {
										check.Visited = 1
										endNextRooms.Visited = 1
										prevRoom.Visited = 1
										roomPaths[count] += prevRoom.Name + ","
										roomPaths[count] += check.Name + ","
										AllPaths(End)
									}
								}
							}
						}
					}
				}
				for i, rooms := range nextRoom {
					if count < lenStart {
						rNamesSlice := strings.Split(roomPaths[count], ",")
						if !Contains(rNamesSlice, rooms.Name) && (rooms.Visited == 0) && !strings.HasPrefix(rooms.Name, "G") {
							if prevRoom != End {
								roomPaths[count] += prevRoom.Name + ","
								prevRoom.Visited = 1
								AllPaths(nextRoom[i])
							}
						}
					}
				}
			}
		}
	}
}

// this returns the index of the smallest path
var roomLength []int

func MinPath(roomLength []int) int {
	min := roomLength[0]
	index := 0
	for i, room := range roomLength {
		if room < min {
			min = room
			index = i
		}
	}
	return index
}

// this assigns the appropriate path to each ant with the smallest number of turns
func AssignPaths() {
	for i := range finalPath {
		a := strings.Split(finalPath[i], ",")
		roomLength = append(roomLength, (len(a) - 2))
	}
	for n := range ants {
		ants[n].Path = finalPath[MinPath(roomLength)]
		roomLength[MinPath(roomLength)]++
		a := strings.Split(ants[n].Path, ",")
		for _, ele := range Start.NextRoom {
			if ele.Name == a[1] {
				ants[n].Room = ele
			}
		}
	}
}

var AntPath string

func TraversePath(r *Room) {
	endAnts := 0
	for endAnts != End.NumAnts {
		for i := range ants {
			if i < len(ants)-1 {
				if ants[i].Room.Visited == 0 && !ants[i].Room.End {
					ants[i].Room.Visited = 1
					AntPath += string("L"+ants[i].Name+"-"+ants[i].Room.Name) + " "
					ants[i].PrevRoom = ants[i].Room
					ants[i].Room = ants[i].Room.NextRoom[0]
				} else if ants[i].Room.Visited == 0 && ants[i].Room.End {
					if !strings.Contains(AntPath, string("L"+ants[i].Name)+"-"+ants[i].Room.Name) {
						AntPath += "L" + ants[i].Name + "-" + ants[i].Room.Name + " "
						ants[i].PrevRoom = ants[i].Room
						endAnts++
					}
				}
			} else if i == len(ants)-1 {
				if ants[i].Room.End {
					AntPath += string("L"+ants[i].Name+"-"+ants[i].Room.Name) + " "
					endAnts++
				}
				if ants[i].Room.Visited == 0 && !ants[i].Room.End {
					ants[i].Room.Visited = 1
					AntPath += "L" + ants[i].Name + "-" + ants[i].Room.Name + " "
					ants[i].PrevRoom = ants[i].Room
					ants[i].Room = ants[i].Room.NextRoom[0]
				}
				for j := range ants {
					if ants[j].PrevRoom != nil {
						ants[j].PrevRoom.Visited = 0
					}
				}
			}
		}

		AntPath += "\n"
	}

	AntPath = AntPath[:len(AntPath)-1]
}
