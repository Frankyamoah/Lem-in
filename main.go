package main

import (
	"fmt"
	"lem-in-project/control"
	"log"
	"os"
	"strconv"
	"strings"
)

type Antfarm struct {
	Numants   int
	Start     string
	End       string
	Roomnames []string
	Xcoords   []int
	Ycoords   []int
	Links     []string
}

func main() {

	inputfile, err := os.ReadFile(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}
	splitfile := strings.Split(string(inputfile), "\n")
	farmstruct := new(Antfarm)

	for i, v := range splitfile {
		switch i {
		case 0:
			farmstruct.Numants, err = strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}

		case i:

			if v == "##start" && len(v) == 7 {
				farmstruct.Start = v
			} else if v == "##end" && len(v) == 5 {
				farmstruct.End = v
			} else {
				sep := strings.Split(splitfile[i], "")
				join := strings.Join(sep, "")
				for _, x := range sep {
					if x == "-" {
						farmstruct.Links = append(farmstruct.Links, control.GetLinks(join))

					}
					if x == " " {
						rooms := strings.Split(join, " ")

						farmstruct.Roomnames = append(farmstruct.Roomnames, rooms[0])

						x, err := strconv.Atoi(rooms[1])
						if err != nil {
							log.Fatal(err)
						}
						farmstruct.Xcoords = append(farmstruct.Xcoords, x)

						y, err := strconv.Atoi(rooms[2])
						if err != nil {
							log.Fatal(err)
						}
						farmstruct.Ycoords = append(farmstruct.Ycoords, y)
						break
					}

				}

			}
		}

	}
	fmt.Println(farmstruct.Numants)
	fmt.Println(farmstruct.Start)
	fmt.Println(farmstruct.Roomnames)
	fmt.Println(farmstruct.Xcoords)
	fmt.Println(farmstruct.Ycoords)
	fmt.Println(farmstruct.End)
	fmt.Println(farmstruct.Links)
}
