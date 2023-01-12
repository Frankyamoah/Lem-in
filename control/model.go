package control

import (
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

func GetLinks(s string) string {

	var x []string

	for _, v := range s {
		if v == '-' {
			x = strings.Split(s, "-")

		}

	}
	new := strings.Join(x, ",")
	return new
}
