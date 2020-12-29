package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calc(x int, op string, y int) int {
	if op == "+" {
		return x + y
	}
	return x * y
}

func parseTerm(equation []string, i int) (int, int) {
	if equation[i] == "(" {
		v, i2 := evalExpr(equation, i+1)
		return v, i2 + 1
	}
	v, _ := strconv.Atoi(equation[i])
	return v, i + 1
}

func evalExpr(equation []string, i int) (int, int) {
	t1, i2 := parseTerm(equation, i)
	ans := t1

	for true {
		if i2 == len(equation) || equation[i2] == ")" {
			return ans, i2
		}
		t2, i3 := parseTerm(equation, i2+1)
		ans = calc(ans, equation[i2], t2)
		i2 = i3
	}
	return t1, i2
}

func main() {
	file, _ := os.Open("day18.in")
	scanner := bufio.NewScanner(file)
	p1 := 0

	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), "(", "( ")
		line = strings.ReplaceAll(line, ")", " )")
		equation := strings.Split(line, " ")
		ans, _ := evalExpr(equation, 0)
		p1 += ans
	}
	fmt.Println("P1:", p1)
}
