package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findRow(line string) int {
	binaryString := strings.Replace(strings.Replace(line, "B", "1", 7), "F", "0", 7)
	row, _ := strconv.ParseInt(binaryString, 2, 7)
	return int(row)
}

func findCol(line string) int {
	binaryString := strings.Replace(strings.Replace(line, "R", "1", 3), "L", "0", 3)
	col, _ := strconv.ParseInt(binaryString, 2, 3)
	return int(col)
}

func calculateSeatID(row int, col int) int {
	return row*8 + col
}

func main() {
	highestSeatID := 0
	maxSeatID := 127*8 + 7
	occupiedSeats := make([]int, maxSeatID)

	file, _ := os.Open("day5.in")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := findRow(line[:7])
		col := findCol(line[7:])
		seatID := calculateSeatID(row, col)
		occupiedSeats[seatID] = seatID
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	fmt.Println("Highest seat ID:", highestSeatID)
	for i := 1; i < len(occupiedSeats[:highestSeatID+1]); i++ {
		if occupiedSeats[i] == 0 && occupiedSeats[i-1] > 0 && occupiedSeats[i+1] > 0 {
			fmt.Println("My seat ID:", i)
		}
	}
}
