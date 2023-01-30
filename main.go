package main

import (
	"fmt"
	"lem-in-project/control"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	inputfile, err := os.ReadFile(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}
	splitfile := strings.Split(string(inputfile), "\n")
	farmstruct := new(control.Antfarm)

	for i, v := range splitfile {
		switch i {
		case 0:
			farmstruct.Numants, err = strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}

		case i:

			if v == "##start" && len(v) == 7 {
				startroom := strings.Split(splitfile[i+1], " ")
				farmstruct.Roomnames = append(farmstruct.Roomnames, startroom[0])
				farmstruct.Roomtype = append(farmstruct.Roomtype, "start")
			} else if v == "##end" && len(v) == 5 {
				endroom := strings.Split(splitfile[i+1], " ")
				farmstruct.Roomnames = append(farmstruct.Roomnames, endroom[0])
				farmstruct.Roomtype = append(farmstruct.Roomtype, "end")
			} else if v != "##start" && v != "##end" {
				midroom := strings.Split(splitfile[i], " ")
				farmstruct.Roomnames = append(farmstruct.Roomnames, midroom[0])
				farmstruct.Roomtype = append(farmstruct.Roomtype, "middle")

			} else {
				sep := strings.Split(splitfile[i], "")
				join := strings.Join(sep, "")
				for _, x := range sep {
					if x == "-" {

						farmstruct.Adjacent = append(farmstruct.Adjacent, sep)

					}

					if x == " " {
						rooms := strings.Split(join, " ")

						farmstruct.Roomnames = append(farmstruct.Roomnames, rooms[0])
						break
					}

				}

			}
		}

	}
	fmt.Println("Number of ants:", farmstruct.Numants)
	fmt.Println("Start room:", farmstruct.Roomnames)
	fmt.Println("Room types:", farmstruct.Roomtype)

	test := control.Graph{}

	for _, v := range farmstruct.Roomnames {

		test.AddVertex(v)

	}
	var from []string
	var to []string
	adj := farmstruct.Adjacent
	for range adj {

		from = adj[0]
		to = adj[1]
	}
	fmt.Println(adj, "x")

	test.AddEdge(from, to)
	test.Print()
}
