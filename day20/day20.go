package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Tile struct {
	borders   []string
	neighbors int
}

func matchTiles(tilesMap map[int]*Tile, tileIDs []int) {
	for _, id := range tileIDs {
		for _, iid := range tileIDs {
			if id == iid {
				continue
			}
			for _, border := range tilesMap[id].borders {
				for _, nb := range tilesMap[iid].borders {
					if border == nb {
						tilesMap[id].neighbors++
					}
				}
			}
		}
	}
}

func main() {
	dat, _ := ioutil.ReadFile("day20_ex.in")
	tiles := strings.Split(string(dat), "\n\n")
	fmt.Println(tiles)

	tilesMap := map[int]*Tile{}
	tileIDs := []int{}

	for _, tile := range tiles {
		tileID := 0
		borders := []string{}
		l, r := "", ""
		firstL := true
		for i, line := range strings.Split(tile, "\n") {
			if strings.HasPrefix(line, "Tile") {
				tileID, _ = strconv.Atoi(strings.Split(line[:len(line)-1], " ")[1])
				tileIDs = append(tileIDs, tileID)
				continue
			}
			if firstL {
				borders = append(borders, line)
				firstL = false
			}
			l += string(line[0])
			r += string(line[len(line)-1])
			if i == len(line) {
				borders = append(borders, line)
			}
		}

		borders = append(borders, l)
		borders = append(borders, r)
		tilesMap[tileID] = &Tile{borders: borders, neighbors: 0}
	}

	matchTiles(tilesMap, tileIDs)

	// for k, v := range tilesMap {
	// 	fmt.Println("key", k, "value", v)
	// }
}
