package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Tile struct {
	top, left, right, bot         string
	topID, leftID, rightID, botID int
}

func matchTiles(tilesMap map[int]*Tile, tileIDs []int) {
	for _, id := range tileIDs {
		for _, iid := range tileIDs {
			if id == iid {
				continue
			}
			if tilesMap[id].top == tilesMap[iid].bot || tilesMap[id].top == tilesMap[iid].top {
				tile := tilesMap[id]
				tile.topID = iid
				tilesMap[id] = tile
			}
			if tilesMap[id].left == tilesMap[id].left || tilesMap[id].left == tilesMap[iid].right {
				tile := tilesMap[id]
				tile.leftID = iid
				tilesMap[id] = tile
			}
			if tilesMap[id].right == tilesMap[id].right || tilesMap[id].right == tilesMap[iid].left {
				tile := tilesMap[id]
				tile.rightID = iid
				tilesMap[id] = tile
			}
			if tilesMap[id].bot == tilesMap[iid].bot || tilesMap[id].bot == tilesMap[iid].top {
				tile := tilesMap[id]
				tile.botID = iid
				tilesMap[id] = tile
			}
		}
	}
}

func main() {
	dat, _ := ioutil.ReadFile("day20_ex.in")

	tiles := strings.Split(string(dat), "\n\n")

	tilesMap := map[int]*Tile{}
	tileIDs := []int{}

	for _, tile := range tiles {
		tileID := 0
		l, r, t, b := "", "", "", ""
		firstL := true
		for i, line := range strings.Split(tile, "\n") {
			if strings.HasPrefix(line, "Tile") {
				tileID, _ = strconv.Atoi(strings.Split(line[:len(line)-1], " ")[1])
				tileIDs = append(tileIDs, tileID)
				fmt.Println("Tile", tileID)
				continue
			}
			if firstL {
				t = line
				firstL = false
			}
			l += string(line[0])
			r += string(line[len(line)-1])
			if i == len(line) {
				b = line
			}
		}

		tilesMap[tileID] = &Tile{top: t, left: l, right: r, bot: b}
	}

	matchTiles(tilesMap, tileIDs)

	for k, v := range tilesMap {
		fmt.Println(k, v)
	}
}
