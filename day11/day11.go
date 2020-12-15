package main

import (
	"bufio"
	"fmt"
	"os"
)

func paddingRow(length int, padSymbol string) []string {
	row := make([]string, length)
	for i := range row {
		row[i] = padSymbol
	}
	return row
}

func padLine(line string, padSymbol string) []string {
	row := []string{padSymbol}
	for _, c := range line {
		if c != '\n' {
			row = append(row, string(c))
		}
	}
	row = append(row, padSymbol)
	return row
}

func generateBoard(filename string) [][]string {
	board := make([][]string, 0)
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	first := true
	for scanner.Scan() {
		line := scanner.Text()
		if first {
			board = append(board, paddingRow(len(line)+2, "."))
			first = false
		}
		board = append(board, padLine(line, "."))
	}
	board = append(board, paddingRow(len(board[0]), "."))
	return board
}

func occupiedSeats(board [][]string) int {
	occupied := 0
	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board)-1; j++ {
			if board[i][j] == "#" {
				occupied++
			}
		}
	}
	return occupied
}

func getNeighbors(board [][]string, i int, j int) string {
	top := board[i-1][j-1] + board[i-1][j] + board[i-1][j+1]
	mid := board[i][j-1] + board[i][j+1]
	bot := board[i+1][j-1] + board[i+1][j] + board[i+1][j+1]
	return top + mid + bot
}

func countSymbol(neighbors string, symbol rune) int {
	count := 0
	for _, c := range neighbors {
		if c == symbol {
			count++
		}
	}
	return count
}

func symbolNotIn(neighbors string, symbol rune) bool {
	for _, c := range neighbors {
		if c == symbol {
			return false
		}
	}
	return true
}

func applyRules(board [][]string) ([][]string, bool) {
	changed := false
	nextBoard := make([][]string, len(board))

	for i := 0; i < len(board); i++ {
		nextBoard[i] = make([]string, len(board[i]))
		copy(nextBoard[i], board[i])
	}

	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board)-1; j++ {
			neighbors := getNeighbors(board, i, j)
			if board[i][j] == "L" && symbolNotIn(neighbors, '#') {
				nextBoard[i][j] = "#"
				changed = true
			} else if board[i][j] == "#" && countSymbol(neighbors, '#') >= 4 {
				nextBoard[i][j] = "L"
				changed = true
			}
		}
	}

	// for i := range nextBoard {
	// 	fmt.Println("#####")
	// 	fmt.Println(board[i])
	// 	fmt.Println(nextBoard[i])
	// }

	return nextBoard, changed
}

func main() {
	board := generateBoard("day11.in")
	nextBoard, changed := applyRules(board)

	i := 0
	for changed {
		nextBoard, changed = applyRules(nextBoard)
		i++
		fmt.Println(i)
	}

	fmt.Println("P1:", occupiedSeats(nextBoard))

	// fmt.Println(changed)
}
