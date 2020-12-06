package main

import (
	"bufio"
	"fmt"
	"os"
)

func groupScoreP1(group string) int {
	scores := make([]int, 26)
	score := 0
	for _, c := range group {
		asciiC := int(c) - 97
		if scores[asciiC] == 0 {
			scores[asciiC] = 1
			score++
		}
	}
	return score
}

func groupScoreP2(group string, people int) int {
	scores := make([]int, 26)
	score := 0
	for _, c := range group {
		scores[int(c)-97]++
	}
	for i := 0; i < len(scores); i++ {
		if scores[i] == people {
			score++
		}
	}
	return score
}

func main() {
	totalScoreP1 := 0
	totalScoreP2 := 0

	file, _ := os.Open("day6.in")
	scanner := bufio.NewScanner(file)
	group := ""
	people := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			group += line
			people++
		} else {
			totalScoreP1 += groupScoreP1(group)
			totalScoreP2 += groupScoreP2(group, people)
			group = ""
			people = 0
		}
	}
	fmt.Println("Answer P1:", totalScoreP1)
	fmt.Println("Answer P2:", totalScoreP2)
}
