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

func solve(input string, isP2 bool) string {
	var n int
	if isP2 {
		n = int(1e6)
	} else {
		n = len(input)
	}
	N := make([]int, n+1)
	X := []int{}
	for _, in := range input {
		num, _ := strconv.Atoi(string(in))
		X = append(X, num)
	}

	for i := 0; i < len(X); i++ {
		N[X[i]] = X[(i+1)%len(X)]
	}

	if isP2 {
		N[X[len(X)-1]] = len(X) + 1
		for i := 0; i < len(X)+1; i = i + (n + 1) {
			N[i] = i + 1
		}
		N[len(N)-1] = X[0]
	}

	t := 0
	current := X[0]
	nMoves := 100
	var dest int
	for moves := nMoves; 0 < moves; moves-- {
		t++
		pickup := N[current]
		N[current] = N[N[N[pickup]]]

		if current == 1 {
			dest = n
		} else {
			dest = current - 1
		}

		slice := []int{pickup, N[pickup], N[N[pickup]]}
		for in(dest, slice) {
			fmt.Println(dest, slice)
			if dest == 1 {
				dest = n
			} else {
				dest--
			}
		}
		N[N[N[pickup]]] = N[dest]
		N[dest] = pickup
		current = N[current]
	}

	if isP2 {
		fmt.Println("P2:", N[1]*N[N[1]])
	}

	tmp := []int{}
	x := N[1]
	for x != 1 {
		tmp = append(tmp, x)
		x = N[x]
	}

	ans := ""
	for _, i := range tmp {
		ans += fmt.Sprint(i)
	}
	return ans
}

func main() {
	input, _ := ioutil.ReadFile("day23_ex.in")

	fmt.Println(solve(string(input), false))
	//fmt.Println(solve(string(input), true))
}
