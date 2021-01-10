package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/tebeka/deque"
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

type Decks struct {
	D1, D2 *deque.Deque
}

func recursiveGame(D1 *deque.Deque, D2 *deque.Deque) (bool, *deque.Deque) {
	SEEN := map[Decks]bool{}
	for D1.Len() > 0 && D2.Len() > 0 {
		decks := Decks{D1, D2}
		if _, ok := SEEN[decks]; ok {
			return true, D1
		}
		SEEN[decks] = true
		c1, _ := D1.PopLeft()
		c2, _ := D2.PopLeft()
		var p1_wins bool
		if D1.Len() >= c1.(int) && D2.Len() >= c2.(int) {
			NewD1, NewD2 := deque.New(), deque.New()
			for x := 0; x < c1.(int); x++ {
				c, _ := D1.Get(x)
				NewD1.Append(c.(int))
			}
			for x := 0; x < c2.(int); x++ {
				c, _ := D2.Get(x)
				NewD2.Append(c.(int))
			}
			p1_wins, _ = recursiveGame(NewD1, NewD2)
		}

		if p1_wins {
			D1.Append(c1)
			D1.Append(c2)
		} else {
			D2.Append(c2)
			D2.Append(c1)
		}
	}

	if D1.Len() > 0 {
		return true, D1
	}
	return false, D2
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

	D1, D2 := deque.New(), deque.New()
	for i := range player1 {
		D1.Append(player1[i])
		D2.Append(player2[i])
	}

	player1, player2 = game(player1, player2)
	if len(player1) == 0 {
		fmt.Println("P1", score(player2))
	} else {
		fmt.Println("P1", score(player1))
	}

	_, deck := recursiveGame(D1, D2)
	winner_deck := []int{}
	for deck.Len() > 0 {
		c, _ := deck.PopLeft()
		winner_deck = append(winner_deck, c.(int))
	}
	var score int = 0
	for i, c := range winner_deck {
		score += (len(winner_deck) - i) * c
	}
	fmt.Println("p2", score)
}
