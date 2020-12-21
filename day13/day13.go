package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Constraint struct {
	b int
	i int
}

func gcd(x int, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}

func modPow(b int, e int, mod int) int {
	if e == 0 {
		return 1
	}
	return (b * modPow(b, e-1, mod)) % mod
}

func modInverse(a int, m int) int {
	return modPow(a%m, m-2, m)
}

func p1(startTime int, b string) {
	busesSplit := strings.Split(strings.Replace(b, "x", "", -1), ",")
	buses := make([]int, 0)
	for _, bus := range busesSplit {
		ib, _ := strconv.Atoi(bus)
		if ib != 0 {
			buses = append(buses, ib)
		}
	}

	found := true
	currentTime := startTime
	for found {
		for _, bus := range buses {
			if currentTime%bus == 0 {
				waitTime := currentTime - startTime
				fmt.Println("P1:", bus*waitTime)
				found = false
			}
		}
		currentTime++
	}
}

// https://en.wikipedia.org/wiki/Chinese_remainder_theorem#Existence_(direct_construction)
func p2(b string) {
	bs := strings.Split(b, ",")
	constraints := make([]Constraint, 0)
	N := 1
	for i, bus := range bs {
		if bus != "x" {
			b, _ := strconv.Atoi(bus)
			i %= b
			constraints = append(constraints, Constraint{b: b, i: (b - i) % b})
			N *= b
		}
	}

	x := 0
	for _, a := range constraints {
		// NI product of all other bus IDs
		NI := N / a.b
		MI := modInverse(NI, a.b)
		x += a.i * MI * NI
	}
	fmt.Println("P2:", x%N)
}

func main() {
	file, _ := os.Open("day13.in")

	reader := bufio.NewReader(file)

	t, _ := reader.ReadString('\n')
	startTime, _ := strconv.Atoi(t[:len(t)-2])
	b, _ := reader.ReadString('\n')

	p1(startTime, b)
	p2(b)
}
