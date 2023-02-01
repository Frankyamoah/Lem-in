package lemin

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var ants []*Ant

// This is to obtain and create an address for n number of ants.
func GetAnts() {
	data, err1 := os.Open(os.Args[1])
	if err1 != nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("File Error")
		log.Fatal()
	}
	getants := bufio.NewScanner(data)
	line := 0

	for getants.Scan() {
		if getants.Text() == "" {
			fmt.Println("ERROR: invalid data format")
			log.Fatal()
		}
		line++
		if line == 1 {
			a, err2 := strconv.Atoi(getants.Text())
			if a == 0 {
				fmt.Println("ERROR: invalid data format")
				fmt.Println("No ants found")
				log.Fatal()
			}
			if err2 != nil {
				fmt.Println("ERROR: invalid data format")
				fmt.Println("No ants found")
				log.Fatal()
			}
			ants = make([]*Ant, a)
			for i := 0; i < a; i++ {
				antName := &Ant{Name: strconv.Itoa(i + 1)}
				ants[i] = antName
			}
		}
	}
}
