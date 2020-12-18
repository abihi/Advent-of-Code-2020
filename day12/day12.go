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

type Waypoint struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calcFacing(facing int, value int, action string) int {
	if action == "L" {
		leftRotations := value / 90
		for i := 0; i < leftRotations; i++ {
			if facing == 0 {
				facing = 360
			}
			facing -= 90
		}
		return facing
	}
	return abs(facing+value) % 360
}

func moveShipP1(ship Ship, action string, value int) Ship {
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
		ship.facing = calcFacing(ship.facing, value, action)
	case "L":
		ship.facing = calcFacing(ship.facing, value, action)
	case "F":
		direction := ship.facing / 90
		if direction < 1 {
			ship = moveShipP1(ship, "N", value)
		} else if direction < 2 {
			ship = moveShipP1(ship, "E", value)
		} else if direction < 3 {
			ship = moveShipP1(ship, "S", value)
		} else {
			ship = moveShipP1(ship, "W", value)
		}
	default:
		fmt.Println(action, "<-- Invalid action")
	}
	return ship
}

func moveWaypointP2(ship Ship, waypoint Waypoint, action string, value int) (Ship, Waypoint) {
	switch action {
	case "N":
		waypoint.y += value
	case "S":
		waypoint.y -= value
	case "E":
		waypoint.x += value
	case "W":
		waypoint.x -= value
	case "R":
		ship.facing = calcFacing(ship.facing, value, action)
	case "L":
		ship.facing = calcFacing(ship.facing, value, action)
	case "F":
		ship.x = value * waypoint.x
		ship.y = value * waypoint.y
		waypoint.x = ship.x + waypoint.x
		waypoint.y = ship.y + waypoint.y
	default:
		fmt.Println(action, "<-- Invalid action")
	}
	return ship, waypoint
}

func main() {
	file, _ := os.Open("day12.in")
	scanner := bufio.NewScanner(file)

	shipP1 := Ship{x: 0, y: 0, facing: 90}
	shipP2 := Ship{x: 0, y: 0, facing: 90}
	waypoint := Waypoint{x: 10, y: 1}

	for scanner.Scan() {
		line := scanner.Text()
		action := string(line[0])
		value, _ := strconv.Atoi(line[1:])
		shipP1 = moveShipP1(shipP1, action, value)
		shipP2, waypoint = moveWaypointP2(shipP2, waypoint, action, value)
		//fmt.Println(action, value)
		//fmt.Println("x:", ship.x, "y:", ship.y, "facing:", ship.facing)
	}

	fmt.Println("P1:", abs(shipP1.x)+abs(shipP1.y))
}
