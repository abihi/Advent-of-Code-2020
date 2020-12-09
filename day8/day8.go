package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	argument  int
	count     int
}

func checkIfNotLoop(q []instruction, i int) bool {
	for i < len(q) {
		if q[i].operation == "nop" {
			if q[i].count == 1 {
				return false
			}
			q[i].count++
			i++
		} else if q[i].operation == "jmp" {
			if q[i].count == 1 {
				return false
			}
			q[i].count++
			i += q[i].argument
		} else {
			if q[i].count == 1 {
				return false
			}
			q[i].count++
			i++
		}
	}
	return true
}

func fixCorruptedOperation(q []instruction) []instruction {
	i := 0
	tmp := make([]instruction, len(q))
	for i < len(q) {
		copy(tmp, q)
		if q[i].operation == "nop" {
			tmp[i].operation = "jmp"
			if checkIfNotLoop(tmp, i) {
				q[i].operation = "jmp"
				return q
			}
			i++
		} else if q[i].operation == "jmp" {
			tmp[i].operation = "nop"
			if checkIfNotLoop(tmp, i) {
				q[i].operation = "nop"
				return q
			}
			i += q[i].argument
		} else {
			i++
		}
	}
	return q
}

func readInstructions(q []instruction) int {
	i := 0
	acc := 0
	for i < len(q) {
		if q[i].operation == "nop" {
			i++
		} else if q[i].operation == "jmp" {
			i += q[i].argument
		} else {
			acc += q[i].argument
			i++
		}
	}
	return acc
}

func main() {
	q := make([]instruction, 0)
	file, _ := os.Open("day8.in")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inst := strings.Split(line, " ")
		op := inst[0]
		arg, _ := strconv.Atoi(strings.TrimRight(inst[1], "\n"))
		q = append(q, instruction{operation: op, argument: arg, count: 0})
	}
	fixCorruptedOperation(q)
	fmt.Println(readInstructions(q))
}
