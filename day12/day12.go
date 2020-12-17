package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Ship struct {
	x      int
	y      int
	facing int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func moveShip(ship Ship, action string, value int) Ship {
	switch action {
	case "N":
		ship.y += value
	case "S":
		ship.y -= value
	case "E":
		ship.x += value
	case "W":
		ship.x -= value
	case "R":
		ship.facing = abs(ship.facing+value) % 360
	case "L":
		ship.facing = abs(ship.facing-value) % 360
	case "F":
		direction := ship.facing / 90
		if direction < 1 {
			ship = moveShip(ship, "N", value)
		} else if direction < 2 {
			ship = moveShip(ship, "E", value)
		} else if direction < 3 {
			ship = moveShip(ship, "S", value)
		} else {
			ship = moveShip(ship, "W", value)
		}
	default:
		fmt.Println(action, "<-- Invalid action")
	}
	return ship
}

func main() {
	file, _ := os.Open("day12_ex.in")
	scanner := bufio.NewScanner(file)

	ship := Ship{x: 0, y: 0, facing: 90}

	for scanner.Scan() {
		line := scanner.Text()
		action := string(line[0])
		value, _ := strconv.Atoi(line[1:])
		ship = moveShip(ship, action, value)
		fmt.Println(action, value)
		fmt.Println("x:", ship.x, "y:", ship.y, "facing:", ship.facing)
	}

	fmt.Println("P1:", abs(ship.x)+abs(ship.y))
}
