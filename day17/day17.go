package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	x, y, z, w int
}

func neighbors(active map[Coord]bool, x int, y int, z int, w int) int {
	nbrs := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if dx != 0 || dy != 0 || dz != 0 || dw != 0 {
						if _, ok := active[Coord{x + dx, y + dy, z + dz, w + dw}]; ok {
							nbrs++
						}
					}
				}
			}
		}
	}
	return nbrs
}

func simulation(active map[Coord]bool) map[Coord]bool {
	newActive := map[Coord]bool{}
	// Bounding volume
	for x := -15; x < 15; x++ {
		for y := -15; y < 15; y++ {
			for z := -8; z < 8; z++ {
				for w := -8; w < 8; w++ {
					nbrs := neighbors(active, x, y, z, w)
					if _, ok := active[Coord{x, y, z, w}]; ok && nbrs == 3 {
						newActive[Coord{x, y, z, w}] = true
					}
					if _, ok := active[Coord{x, y, z, w}]; ok && nbrs == 2 || nbrs == 3 {
						newActive[Coord{x, y, z, w}] = true
					}
				}
			}
		}
	}
	return newActive
}

func main() {
	file, _ := os.Open("day17.in")
	scanner := bufio.NewScanner(file)

	active := map[Coord]bool{}
	x := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for y, ch := range line {
			if ch == "#" {
				active[Coord{x, y, 0, 0}] = true
			}
		}
		x++
	}

	cycles := 6
	for i := 0; i < cycles; i++ {
		active = simulation(active)
	}

	fmt.Println("Ans:", len(active))
}
