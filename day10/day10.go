package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(array []int) (int, int) {
	ind := 0
	min := array[0]
	for i, value := range array {
		if min > value {
			min = value
			ind = i
		}
	}
	return min, ind
}

func sort(adapters []int) []int {
	ordered := make([]int, 0)

	for len(adapters) > 0 {
		min, ind := min(adapters)
		ordered = append(ordered, min)
		adapters = append(adapters[:ind], adapters[ind+1:]...)
	}

	return ordered
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findDifferences(sortedAdapters []int) (int, int) {
	diff1 := 0
	diff3 := 0

	for i := 0; i < len(sortedAdapters)-1; i++ {
		diff := abs(sortedAdapters[i] - sortedAdapters[i+1])
		if diff == 1 {
			diff1++
		} else if diff == 3 {
			diff3++
		}
	}
	return diff1, diff3
}

func countCombinations(mem []int, sortedAdapters []int, current int) int {
	if mem[current] != 0 {
		return mem[current]
	}

	count := 0
	if sortedAdapters[current] == sortedAdapters[len(sortedAdapters)-1] {
		count++
		mem[current] = count
	}

	for i := 1; i <= 3; i++ {
		if current+i >= len(sortedAdapters) {
			return count
		}

		diff := abs(sortedAdapters[current] - sortedAdapters[current+i])
		if diff == 1 {
			count += countCombinations(mem, sortedAdapters, current+i)
		} else if diff == 2 {
			count += countCombinations(mem, sortedAdapters, current+i)
		} else if diff == 3 {
			count += countCombinations(mem, sortedAdapters, current+i)
		}
	}

	return count
}

func main() {
	adapters := make([]int, 0)
	file, _ := os.Open("day10.in")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		a, _ := strconv.Atoi(line)
		adapters = append(adapters, int(a))
	}
	sortedAdapters := sort(adapters)
	fmt.Println(sortedAdapters)
	diff1, diff3 := findDifferences(sortedAdapters)
	fmt.Println("p1", diff1*diff3)
	// Memoization, store computations outside recursive function
	mem := make([]int, len(sortedAdapters))
	c := countCombinations(mem, sortedAdapters[1:], 0)
	fmt.Println(mem)
	fmt.Println("p2", c)
}
