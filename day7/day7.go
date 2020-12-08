package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Holds struct {
	amount int
	name   string
}

type Bag struct {
	name  string
	holds []Holds
}

func contains(v string, slice []string) bool {
	for _, s := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func fillBagsMap(filename string) map[string]Bag {
	bagsMap := make(map[string]Bag)

	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(strings.Replace(line, ".", "", 1), "\n", "", 1)
		containSplit := strings.Split(line, "contain")
		capacitySplit := strings.Split(containSplit[1], ",")
		bagName := strings.TrimRight(containSplit[0], " ")
		holds := make([]Holds, 0)
		if contains("no other bags", capacitySplit) {
			holds = append(holds, Holds{amount: 0, name: "no other bags"})
			bag := Bag{name: bagName, holds: holds}
			bagsMap[bagName] = bag
			continue
		}
		for _, c := range capacitySplit {
			amount, _ := strconv.Atoi(string(c[1]))
			name := c[3:]
			holds = append(holds, Holds{amount: amount, name: name})
		}
		bag := Bag{name: bagName, holds: holds}
		bagsMap[bagName] = bag
	}

	return bagsMap
}

func holdsShinyBag(key string, bagsMap map[string]Bag) bool {
	for _, bag := range bagsMap[key].holds {
		if "shiny gold bags" == bag.name {
			return true
		}
		if holdsShinyBag(bag.name, bagsMap) {
			return true
		}
	}
	return false
}

func bagsInBag(key string, bagsMap map[string]Bag) int {
	count := 0
	for _, bag := range bagsMap[key].holds {
		if bag.amount > 0 {
			count += bag.amount + bag.amount*bagsInBag(bag.name, bagsMap)
		}
	}
	return count
}

func main() {
	answerP1 := 0
	bagsMap := fillBagsMap("day7.in")
	for key := range bagsMap {
		if holdsShinyBag(key, bagsMap) {
			answerP1++
		}
	}
	answerP2 := bagsInBag("shiny gold bags", bagsMap)
	fmt.Println(answerP1)
	fmt.Println(answerP2)
}
