package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Tile struct {
	x, y int
}

func flipTiles(flip []string, visitedTiles map[Tile]int) {
	// nw (-1, 1) ne (1, 1) e (1, 0) w (-1, 0) se (1, -1) sw (-1, -1)
	x, y := 0, 0
	for _, move := range flip {
		switch move {
		case "nw":
			x--
			y++
		case "ne":
			x++
			y++
		case "e":
			x++
		case "w":
			x--
		case "se":
			x++
			y--
		case "sw":
			x--
			y--
		}
	}
	visitedTiles[Tile{x, y}]++
}

func main() {
	input, _ := ioutil.ReadFile("day24_ex.in")
	flips := [][]string{}

	for _, in := range strings.Split(string(input), "\n") {
		flip := []string{}
		i := 0
		for i < len(in) {
			if in[i] == 's' && in[i+1] == 'e' {
				flip = append(flip, "se")
				i += 2
			} else if in[i] == 's' && in[i+1] == 'w' {
				flip = append(flip, "sw")
				i += 2
			} else if in[i] == 'n' && in[i+1] == 'e' {
				flip = append(flip, "ne")
				i += 2
			} else if in[i] == 'n' && in[i+1] == 'w' {
				flip = append(flip, "nw")
				i += 2
			} else if in[i] == 'w' {
				flip = append(flip, "w")
				i++
			} else if in[i] == 'e' {
				flip = append(flip, "e")
				i++
			}
		}
		flips = append(flips, flip)
	}

	visitedTiles := map[Tile]int{}
	for _, flip := range flips {
		flipTiles(flip, visitedTiles)
	}
	fmt.Println(visitedTiles)

	p1 := 0
	for _, tile := range visitedTiles {
		if !(tile%2 == 0) {
			p1++
		}
	}
	fmt.Println("P1", p1)
}
