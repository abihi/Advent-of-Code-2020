package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func preprocesLine(line string) (int, int, byte, string) {
	splitLine := strings.Split(line, " ")
	password := splitLine[2]
	letter := splitLine[1][0]

	splitMinMax := strings.Split(splitLine[0], "-")
	min, _ := strconv.Atoi(splitMinMax[0])
	max, _ := strconv.Atoi(splitMinMax[1])

	return min, max, letter, password
}

func isPasswordValidPart1(min int, max int, letter byte, password string) bool {
	count := strings.Count(password, string(letter))
	if count < min || count > max {
		return false
	}
	return true
}

func isPasswordValidPart2(min int, max int, letter byte, password string) bool {
	valid := false
	if password[min-1] == letter {
		valid = true
	}
	if password[max-1] == letter {
		valid = !valid
	}
	return valid
}

func main() {
	file, _ := os.Open("day2_input.txt")

	validPasswordsP1 := 0
	validPasswordsP2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		min, max, letter, password := preprocesLine(line)
		if isPasswordValidPart1(min, max, letter, password) {
			validPasswordsP1++
		}
		if isPasswordValidPart2(min, max, letter, password) {
			validPasswordsP2++
		}
	}
	fmt.Println(validPasswordsP1)
	fmt.Println(validPasswordsP2)
}
