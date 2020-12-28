package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func activeNeighbors(cube [][][]string, x int, y int, z int) int {
	active := 0
	for xx := -1; xx <= 1; xx++ {
		for yy := -1; yy <= 1; yy++ {
			for zz := -1; zz <= 1; zz++ {
				symbol := ""
				if xx == 0 && yy == 0 && zz == 0 {
					continue
				}
				// tb: top-bot, lr: left-right, br: back-forwad
				tb := xx + x
				lr := yy + y
				br := zz + z
				inBounds := 0 <= tb && tb < len(cube) && 0 <= lr && lr < len(cube) && 0 <= br && br < len(cube)
				if inBounds {
					symbol = cube[tb][lr][br]
					if symbol == "#" {
						active++
					}
				}
			}

		}
	}
	return active
}

func addPadding(cube [][][]string) [][][]string {
	paddedCube := [][][]string{}
	return paddedCube
}

func printCube(cube [][][]string) {
	z := -len(cube) / 2
	for _, square := range cube {
		fmt.Println("z=", z)
		for _, line := range square {
			fmt.Println(line)
		}
		z++
	}
}

func simulation(cube [][][]string, cycles int) [][][]string {
	newCube := make([][][]string, len(cube))
	copy(newCube, cube)

	for x := 0; x < len(cube); x++ {
		for y := 0; y < len(cube); y++ {
			for z := 0; z < len(cube); z++ {
				activeN := activeNeighbors(cube, x, y, z)
				if cube[x][y][z] == "#" && !(activeN == 2 || activeN == 3) {
					newCube[x][y][z] = "."
				}
				if cube[x][y][z] == "." && activeN == 3 {
					newCube[x][y][z] = "#"
				}
			}
		}
	}

	return newCube
}

func main() {
	file, _ := os.Open("day17_ex.in")
	scanner := bufio.NewScanner(file)

	square := [][]string{}
	cube := [][][]string{}
	//tesseract := [][][][]string{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		square = append(square, line)
	}

	fmt.Println(square)
	for i := 0; i < 3; i++ {
		cube = append(cube, square)
	}

	cycles := 1
	for i := 0; i < cycles; i++ {
		cube = simulation(cube, 1)
		//printCube(cube)
	}

}
