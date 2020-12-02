package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func find2020(numberList []int) int {
	for _, x := range numberList {
		for _, y := range numberList {
			for _, z := range numberList {
				if x+y+z == 2020 {
					return x * y * z
				}
			}
		}
	}
	return 0
}

func preprocesInput(filename string) []int {
	numberList := make([]int, 0)
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numberList = append(numberList, num)
	}

	return numberList
}

func main() {
	var numberList = preprocesInput("day1_input.txt")
	var answer = find2020(numberList)
	fmt.Println(answer)
}
