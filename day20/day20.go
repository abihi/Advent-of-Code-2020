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

func reverse(str string) string {
	chars := []rune(str)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

func matchTiles(tilesMap map[int]*Tile, tileIDs []int) {
	for _, id := range tileIDs {
		for _, iid := range tileIDs {
			if id == iid {
				continue
			}
			for _, border := range tilesMap[id].borders {
				for _, nb := range tilesMap[iid].borders {
					if border == nb || reverse(border) == nb {
						tilesMap[id].neighbors++
					}
				}
			}
		}
	}
}

func main() {
	dat, _ := ioutil.ReadFile("day20.in")
	tiles := strings.Split(string(dat), "\n\n")

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

	P1 := 1
	for k, v := range tilesMap {
		fmt.Println("key", k, "value", v)
		if v.neighbors == 2 {
			P1 *= k
		}
	}
	fmt.Println("P1:", P1)
}
