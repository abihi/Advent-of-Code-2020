package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func game(input []string, turns int) int {
	numbers := make([]int, 0)
	for _, num := range input {
		n, _ := strconv.Atoi(num)
		numbers = append(numbers, n)
	}

	turn := 1
	numberMap := map[int]int{}
	for _, num := range numbers {
		numberMap[num] = turn
		turn++
	}

	spoken := 0
	for turn < turns {
		if _, ok := numberMap[spoken]; ok {
			prevTurn := numberMap[spoken]
			numberMap[spoken] = turn
			spoken = turn - prevTurn
		} else {
			numberMap[spoken] = turn
			spoken = 0
		}
		turn++
	}
	return spoken
}

func main() {
	file, _ := os.Open("day15.in")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")
		fmt.Println("2020th value:", game(input, 2020))
		fmt.Println("30000000th value:", game(input, 30000000))
	}
}
