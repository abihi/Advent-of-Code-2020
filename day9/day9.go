package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func hasSum(number int, numbers []int) bool {
	for _, x := range numbers {
		for _, y := range numbers {
			if x+y == number {
				return true
			}
		}
	}
	return false
}

func findFirstWeakness(numbers []int, preamble int) int {
	for i := preamble; i < len(numbers); i++ {
		if !hasSum(numbers[i], numbers[i-preamble:i]) {
			return numbers[i]
		}
	}
	return -1
}

func contiguousSet(p1 int, numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		largest := 0
		smallest := int(^uint(0) >> 1)
		setSum := 0
		cSet := make([]int, 0)
		for _, num := range numbers[i:] {
			if num > largest {
				largest = num
			}
			if num < smallest {
				smallest = num
			}
			setSum += num
			cSet = append(cSet, num)
			if setSum == p1 {
				return smallest + largest
			}
		}
	}
	return -1
}

func main() {
	numbers := make([]int, 0)
	file, _ := os.Open("day9.in")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, num)
	}
	p1 := findFirstWeakness(numbers, 25)
	fmt.Println("p1", p1)
	p2 := contiguousSet(p1, numbers)
	fmt.Println("p2", p2)
}
