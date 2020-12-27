package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numbers_outside_range(limits map[string][4]int, nearby_ticket []int) []int {
	outside_range := []int{}
	for _, num := range nearby_ticket {
		inside := false
		for _, m := range limits {
			if (num >= m[0] && num <= m[1]) || (num >= m[2] && num <= m[3]) {
				inside = true
				break
			}
		}
		if !inside {
			outside_range = append(outside_range, num)
		}
	}
	return outside_range
}

func error_rates(limits map[string][4]int, nearby_tickets [][]int) (int, [][]int) {
	error_rate := 0

	valid_tickets := [][]int{}
	for _, nearby_ticket := range nearby_tickets {
		outside_range := numbers_outside_range(limits, nearby_ticket)
		if len(outside_range) == 0 {
			valid_tickets = append(valid_tickets, nearby_ticket)
		}
		for _, num := range outside_range {
			error_rate += num
		}
	}

	return error_rate, valid_tickets
}

const SIZE = 20

func find_fields(valid_tickets [][]int, limits map[string][4]int, keys []string) [SIZE][SIZE]bool {
	fields := [SIZE][SIZE]bool{}

	for x := 0; x < len(fields); x++ {
		for y := 0; y < len(fields); y++ {
			fields[x][y] = true
		}
	}

	for _, valid_ticket := range valid_tickets {
		for x, num := range valid_ticket {
			m := limits[keys[y]]
			fmt.Println("x:", x, "y:", y, num, keys[y])
			if !((m[0] <= num && num <= m[1]) || (m[2] <= num && num <= m[3])) {

				fields[x][y] = false
			}
		}
	}

	M := [SIZE]int{}
	used := [SIZE]bool{}
	found := 0
	for found >= SIZE {
		for i := 0; i < SIZE; i++ {
			valid_j := []int{}
			for j := 0; j < SIZE; j++ {
				if fields[i][j] && !used[j] {
					valid_j = append(valid_j, j)
				}
			}
			if len(valid_j) == 1 {
				fmt.Println(i, valid_j)
				M[i] = valid_j[0]
				used[valid_j[0]] = true
				found++
			}
			if found >= SIZE {
				break
			}
		}
	}

	fmt.Println("M", M)

	return fields
}

func main() {
	file, _ := os.Open("day16.in")
	scanner := bufio.NewScanner(file)

	your_ticket := false
	nearby_tickets := false
	limits := map[string][4]int{}
	your_ticket_numbers := []int{}
	nearby_tickets_numbers := [][]int{}
	keys := []string{}

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
			lim1 := strings.Split(split[1], "-")
			lim2 := strings.Split(split[3], "-")
			min1, _ := strconv.Atoi(lim1[0])
			max1, _ := strconv.Atoi(lim1[1])
			min2, _ := strconv.Atoi(lim2[0])
			max2, _ := strconv.Atoi(lim2[1])
			keys = append(keys, k)
			limits[k] = [4]int{min1, max1, min2, max2}
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

	ticket_scanning_error_rate, valid_tickets := error_rates(limits, nearby_tickets_numbers)

	fields := find_fields(valid_tickets, limits, keys)

	fmt.Println(ticket_scanning_error_rate)
	fmt.Println(valid_tickets)
	for _, field := range fields {
		fmt.Println(field)
	}

}
