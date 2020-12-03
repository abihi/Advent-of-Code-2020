package main

import (
	"bufio"
	"fmt"
	"os"
)

func countTrees(right int, down int) int {
	trees := 0
	file, _ := os.Open("day3.in")

	currentPos := 0
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		if down < 2 && i == 0 {
			continue
		}
		if i%down != 0 {
			continue
		}

		line := scanner.Text()
		currentPos = (currentPos + right) % len(line)
		if line[currentPos] == '#' {
			trees++
		}
	}
	return trees
}

func main() {
	trees11 := countTrees(1, 1)
	trees31 := countTrees(3, 1)
	trees51 := countTrees(5, 1)
	trees71 := countTrees(7, 1)
	trees12 := countTrees(1, 2)
	answer := trees11 * trees31 * trees51 * trees71 * trees12
	fmt.Println(trees12)
	fmt.Println(answer)
}
