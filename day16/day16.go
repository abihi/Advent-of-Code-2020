package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numbers_outside_range(ranges map[string][4]int, nearby_ticket []int) []int {
	outside_range := []int{}
	for _, num := range nearby_ticket {
		inside := false
		for _, m := range ranges {
			if (num >= m[0] && num <= m[1]) || (num >= m[2] && num <= m[3]) {
				inside = true
				break
			}
		}
		if !inside {
			outside_range = append(outside_range, num)
		}
	}
	fmt.Println(outside_range)
	return outside_range
}

func error_rates(ranges map[string][4]int, nearby_tickets [][]int) int {
	error_rate := 0

	for _, nearby_ticket := range nearby_tickets {
		outside_range := numbers_outside_range(ranges, nearby_ticket)
		for _, num := range outside_range {
			error_rate += num
		}
	}

	return error_rate
}

func main() {
	file, _ := os.Open("day16.in")
	scanner := bufio.NewScanner(file)

	your_ticket := false
	nearby_tickets := false
	ranges := map[string][4]int{}
	your_ticket_numbers := []int{}
	nearby_tickets_numbers := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if line == "your ticket:" {
			your_ticket = true
			continue
		}
		if line == "nearby tickets:" {
			your_ticket = false
			nearby_tickets = true
			continue
		}

		if !your_ticket && !nearby_tickets {
			split := strings.Split(line, " ")
			k := split[0][:len(split[0])-1]
			mm1 := strings.Split(split[1], "-")
			mm2 := strings.Split(split[3], "-")
			min1, _ := strconv.Atoi(mm1[0])
			max1, _ := strconv.Atoi(mm1[1])
			min2, _ := strconv.Atoi(mm2[0])
			max2, _ := strconv.Atoi(mm2[1])
			ranges[k] = [4]int{min1, max1, min2, max2}
			fmt.Println(k, ranges[k])
		} else if your_ticket {
			for _, num := range strings.Split(line, ",") {
				i, _ := strconv.Atoi(num)
				your_ticket_numbers = append(your_ticket_numbers, i)
			}
		} else if nearby_tickets {
			ints := []int{}
			for _, num := range strings.Split(line, ",") {
				i, _ := strconv.Atoi(num)
				ints = append(ints, i)
			}
			nearby_tickets_numbers = append(nearby_tickets_numbers, ints)
		}
	}

	ticket_scanning_error_rate := error_rates(ranges, nearby_tickets_numbers)

	fmt.Println(your_ticket_numbers)
	fmt.Println(nearby_tickets_numbers)
	fmt.Println(ticket_scanning_error_rate)
}
