package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Grammar problem: https://en.wikipedia.org/wiki/Context-sensitive_grammar
func consume(x string, rule_number int, rules_map map[int]string) []int {
	rule := rules_map[rule_number]

	// Terminal symbol
	if rule[0] == '"' {
		rule = strings.Trim(rule, "\"")
		if strings.HasPrefix(x, rule) {
			return []int{len(rule)}
		}
		return []int{}
	}

	acc_chains := []int{}
	for _, opt := range strings.Split(rule, " | ") {
		acc_chain := []int{0}
		for _, rnStr := range strings.Split(opt, " ") {
			acc := []int{}
			rn, _ := strconv.Atoi(rnStr)
			for _, ap := range acc_chain {
				ret := consume(x[ap:], rn, rules_map)
				for _, c := range ret {
					acc = append(acc, c+ap)
				}
			}
			acc_chain = acc
		}
		for _, ap := range acc_chain {
			acc_chains = append(acc_chains, ap)
		}
	}
	return acc_chains
}

func main() {
	dat, _ := ioutil.ReadFile("day19.in")
	rules_tests := strings.Split(strings.Trim(string(dat), " "), "\n\n")
	rules := rules_tests[0]
	tests := rules_tests[1]

	rules_map := map[int]string{}
	for _, r := range strings.Split(rules, "\n") {
		rn_val := strings.Split(r, ": ")
		rule_number, _ := strconv.Atoi(rn_val[0])
		if rule_number == 8 {
			rn_val[1] = "42 | 42 8"
		}
		if rule_number == 11 {
			rn_val[1] = "42 31 | 42 11 31"
		}
		rules_map[rule_number] = rn_val[1]
	}

	p2 := 0
	for _, str := range strings.Split(tests, "\n") {
		for _, acc_chain := range consume(str, 0, rules_map) {
			if acc_chain == len(str) {
				p2++
				break
			}
		}
	}
	fmt.Println("P2:", p2)
}
