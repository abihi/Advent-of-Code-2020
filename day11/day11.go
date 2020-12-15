package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func generateBoard(filename string) [][]string {
	board := make([][]string, 0)
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, strings.Split(strings.TrimSuffix(line, "\n"), ""))
	}
	return board
}

func occupiedSeats(board [][]string) int {
	occupied := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == "#" {
				occupied++
			}
		}
	}
	return occupied
}

func countNeighbors(board [][]string, i int, j int) int {
	neighbors := 0
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}
			tb := i + r
			lr := j + c
			for (0 <= tb && tb < len(board)) && (0 <= lr && lr < len(board[0])) && board[tb][lr] == "." {
				tb += r
				lr += c
			}
			if (0 <= tb && tb < len(board)) && (0 <= lr && lr < len(board[0])) {
				if board[tb][lr] == "#" {
					neighbors++
				}
			}
		}
	}
	return neighbors
}

func applyRules(board [][]string) ([][]string, bool) {
	changed := false
	nextBoard := make([][]string, len(board))

	for i := 0; i < len(board); i++ {
		nextBoard[i] = make([]string, len(board[i]))
		copy(nextBoard[i], board[i])
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			neighbors := countNeighbors(board, i, j)
			if board[i][j] == "L" && neighbors == 0 {
				nextBoard[i][j] = "#"
				changed = true
			} else if board[i][j] == "#" && neighbors >= 5 {
				nextBoard[i][j] = "L"
				changed = true
			}
		}
	}

	return nextBoard, changed
}

func main() {
	changed := false
	board := generateBoard("day11.in")

	i := 0
	for true {
		board, changed = applyRules(board)
		if !changed {
			break
		}
		i++
	}

	fmt.Println("P2:", occupiedSeats(board))
	fmt.Println(i)
}
