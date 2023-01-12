package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
			split := strings.Split(v, " ")
			farmstruct.Roomnames = split[0]
			farmstruct.Xcoords, err = strconv.Atoi(split[1])
			if err != nil {
				log.Fatal(err)
			}
			farmstruct.Ycoords, err = strconv.Atoi(split[2])
			if err != nil {
				log.Fatal(err)
			}

			if v == "##start" {
				fmt.Println(len(v))
				farmstruct.Start = v
			} else if v == "##end" {
				farmstruct.End = v
			}

		}

	}
	fmt.Println(farmstruct.Numants)
	fmt.Println(farmstruct.Start)
	fmt.Println(farmstruct.Roomnames)
	fmt.Println(farmstruct.Xcoords)
	fmt.Println(farmstruct.Ycoords)
	fmt.Println(farmstruct.End)
	fmt.Println(farmstruct)
}
