package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func toInt(s []string) []int {
	is := []int{}
	for _, str := range s {
		i, _ := strconv.Atoi(str)
		is = append(is, i)
	}
	return is
}

func score(player []int) int {
	score := 0
	j := 0
	for i := len(player); i > 0; i-- {
		score += player[j] * i
		j++
	}
	return score
}

func game(player1 []int, player2 []int) ([]int, []int) {
	for true {
		if len(player1) == 0 || len(player2) == 0 {
			break
		}
		p1 := player1[0]
		p2 := player2[0]
		player1 = player1[1:]
		player2 = player2[1:]

		if p1 > p2 {
			player1 = append(player1, p1)
			player1 = append(player1, p2)
		} else if p2 > p1 {
			player2 = append(player2, p2)
			player2 = append(player2, p1)
		}
	}
	return player1, player2
}

func main() {
	dat, _ := ioutil.ReadFile("day22.in")
	players := strings.Split(string(dat), "\n\n")
	player1 := toInt(strings.Split(players[0], "\n")[1:])
	player2 := toInt(strings.Split(players[1], "\n")[1:])

	player1, player2 = game(player1, player2)
	if len(player1) == 0 {
		fmt.Println(score(player2))
	} else {
		fmt.Println(score(player1))
	}
}
