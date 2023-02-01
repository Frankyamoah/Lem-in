package lemin

import "strings"

// confirms whether a slice contains a certain value
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// removes a slice of string from an array at the given index
func Remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// removes an invalid path
func StartEnd(s string) string {
	a := strings.Split(s, ",")
	if a[0] == Start.Name && a[len(a)-1] == End.Name {
	} else {
		s = ""
	}
	return s
}

// removes empty paths from path's array
var finalPath []string

func Final() {
	for i := range roomPaths {
		if roomPaths[i] != "" {
			finalPath = append(finalPath, roomPaths[i])
		}
	}

	for k := range roomList {
		if roomList[k].NextRoom != nil {
			roomList[k].Visited = 0
			roomList[k].NextRoom = nil
		}
	}

	for i := range roomPaths {
		b := strings.Split(roomPaths[i], ",")
		for o := 0; o < len(b)-1; o++ {
			for k := range roomList {
				for l := range roomList {
					if b[o] == roomList[k].Name && b[o+1] == roomList[l].Name {
						roomList[k].NextRoom = append(roomList[k].NextRoom, roomList[l])
					}
				}
			}
		}
	}
}

// this sorts the elements of roomPaths in ascending order
func Sort() {
	for i := 0; i < len(finalPath)-1; i++ {
		if len(finalPath[i]) > len(finalPath[i+1]) {
			g := finalPath[i]
			h := finalPath[i+1]
			finalPath[i] = h
			finalPath[i+1] = g
		} else if len(finalPath[i+1]) < len(finalPath[i]) {
			g := finalPath[i+1]
			h := finalPath[i]
			finalPath[i] = h
			finalPath[i+1] = g
		}
	}
}
