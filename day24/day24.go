package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Tile struct {
	x, y, z int
}

func flipTiles(flip []string, visitedTiles map[Tile]bool) {
	x, y, z := 0, 0, 0
	for _, move := range flip {
		switch move {
		case "e":
			x++
			y--
		case "se":
			y--
			z++
		case "sw":
			x--
			z++
		case "w":
			x--
			y++
		case "nw":
			y++
			z--
		case "ne":
			x++
			z--
		}
	}
	tile := Tile{x, y, z}
	if visitedTiles[tile] {
		delete(visitedTiles, tile)
	} else {
		visitedTiles[tile] = true
	}

}

func main() {
	input, _ := ioutil.ReadFile("day24.in")
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

	visitedTiles := map[Tile]bool{}
	for _, flip := range flips {
		flipTiles(flip, visitedTiles)
	}

	fmt.Println("P1", len(visitedTiles))

	nbrPos := [][]int{{1, -1, 0}, {0, -1, 1}, {-1, 0, 1}, {-1, 1, 0}, {0, 1, -1}, {1, 0, -1}}

	for day := 0; day < 100; day++ {
		tiles := map[Tile]bool{}
		neighbors := map[Tile]bool{}

		for tile := range visitedTiles {
			x, y, z := tile.x, tile.y, tile.z
			neighbors[Tile{x, y, z}] = true
			for _, pos := range nbrPos {
				dx, dy, dz := pos[0], pos[1], pos[2]
				neighbors[Tile{x + dx, y + dy, z + dz}] = true
			}
		}

		for tile := range neighbors {
			nbr := 0
			x, y, z := tile.x, tile.y, tile.z
			for _, pos := range nbrPos {
				dx, dy, dz := pos[0], pos[1], pos[2]
				if visitedTiles[Tile{x + dx, y + dy, z + dz}] {
					nbr++
				}
			}
			T := Tile{x, y, z}
			if visitedTiles[T] && (nbr == 1 || nbr == 2) {
				tiles[T] = true
			}
			if !visitedTiles[T] && nbr == 2 {
				tiles[T] = true
			}
		}
		visitedTiles = map[Tile]bool{}
		for k, v := range tiles {
			visitedTiles[k] = v
		}
	}

	fmt.Println("P2", len(visitedTiles))
}
