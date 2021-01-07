package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type Cup struct {
	label, pos int
}

func in(i int, s []int) bool {
	for _, num := range s {
		if i == num {
			return true
		}
	}
	return false
}

func findDestination(curr int, circle []int) int {
	for i, label := range circle {
		if curr == label {
			return i
		}
	}
	return -1
}

func main() {
	input, _ := ioutil.ReadFile("day23_ex.in")
	circle := []int{}
	cup := Cup{}
	min := 10000
	max := 0
	for i, in := range input {
		num, _ := strconv.Atoi(string(in))
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		if i == 0 {
			cup = Cup{num, i}
			circle = append(circle, num)
		} else {
			circle = append(circle, num)
		}
	}

	moves := 2
	for moves > 0 {
		current := cup.label
		currentPos := cup.pos
		puPos := []int{(cup.pos) % len(circle), (cup.pos + 1) % len(circle), (cup.pos + 2) % len(circle)}
		pickedUp := []int{circle[puPos[0]], circle[puPos[1]], circle[puPos[2]]}
		destinationPos := -1

		for true {
			current--
			if current < min {
				current = max
			}
			destinationPos = findDestination(current, circle)
			if destinationPos != -1 && !in(circle[destinationPos], pickedUp) {
				cup.label = circle[destinationPos]
				cup.pos = currentPos
				break
			}
		}

		for _, cup := range circle[destinationPos:] {
			pickedUp = append(pickedUp, cup)
		}

		circle = circle[:destinationPos]

		for _, cup := range pickedUp {
			circle = append(circle, cup)
		}

		moves--
	}

	fmt.Println(circle)
}
