package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func game(input []string) int {
	numbers := make([]int, len(input))
	for _, num := range input {
		n, _ := strconv.Atoi(num)
		numbers = append(numbers, n)
	}
	numberMap := map[int][]int{}
	nn := make([]int, 0)
	nn = append(nn, 2)
	numberMap[numbers[0]] = nn
	fmt.Println(numberMap[numbers[0]])
	for i := 0; i < 2020; i++ {

	}
	return 0
}

func main(){
	file, _ := os.Open("day15_ex.in")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		fmt.Println("input:", input, "2020th value:",game(input))
	}
}