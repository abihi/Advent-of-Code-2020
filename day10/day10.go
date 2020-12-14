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

func find(slice []int, toFind int) int {

	for i := 0; i < len(slice); i++ {
		if slice[i] == toFind {
			return i
		}
	}

	return -1
}

func countCombinations(mem []int, adapters []int, current int) int {
	count := 0
	for i := 1; i <= 3; i++ {
		indx := find(adapters, adapters[current]-i)
		if indx != -1 && adapters[indx] == adapters[0] {
			count++
		}
		if indx != -1 {
			combinations := 0
			if mem[indx] == 0 {
				combinations = countCombinations(mem, adapters, indx)
				mem[indx] = combinations
			} else {
				combinations = mem[indx]
			}
			count += combinations
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
	diff1, diff3 := findDifferences(sortedAdapters)
	fmt.Println("p1", diff1*diff3)

	start := []int{0}
	end := sortedAdapters[len(sortedAdapters)-1] + 3
	sortedAdapters = append(start, sortedAdapters...)
	sortedAdapters = append(sortedAdapters, end)
	// Memoization, store computations outside recursive function
	mem := make([]int, len(sortedAdapters))
	c := countCombinations(mem, sortedAdapters, len(sortedAdapters)-1)
	fmt.Println("p2", c)
}
