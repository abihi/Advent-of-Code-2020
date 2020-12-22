package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bitMask(mask string, bits string) int64 {
	newBits := []byte(bits)
	for i := 0; i < len(mask); i++ {
		if mask[i] != 'X' {
			newBits[i] = mask[i]
		}
	}

	value, err := strconv.ParseInt(string(newBits), 2, 37)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return value
}

func addBit(c chan string, bits string, result string, i int) {
	if i >= len(result) {
		c <- bits
		return
	}
	
	if result[i] == 'X' {
		addBit(c, bits + "0", result, i+1)
		addBit(c, bits + "1", result, i+1)
	} else {
		addBit(c, bits + string(result[i]), result, i+1)
	}	
}

func generateCombinations(result string) <-chan string {
    c := make(chan string)

	// Starting a separate goroutine that will generate all binary strings with n set bits
    go func(c chan string) {
        defer close(c) // Close channel when iteration is done

		addBit(c, "", result, 0)
    }(c)

    return c
}

func decoder(mask string, address string, memP2 map[string]int, value int) map[string]int {
	result := []byte(address)
	for i := 0; i < len(mask); i++ {
		if mask[i] == 'X' {
			result[i] = mask[i]
		} else if mask[i] == '1'{
			result[i] = mask[i]
		} else {
			result[i] = address[i]
		}	
	}

	for combination := range generateCombinations(string(result)) {
		key, _ := strconv.ParseInt(combination, 2, 37)
		memP2[fmt.Sprint(key)] = value
	}

	return memP2
}

func main() {
	memP1 := make(map[string]int64)
	memP2 := make(map[string]int)
	file, _ := os.Open("day14.in")
	scanner := bufio.NewScanner(file)
	mask := ""
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if line[0] == "mask" {
			mask = line[2]
		} else {
			key := line[0][4 : len(line[0])-1]
			keyInt, _ := strconv.Atoi(key)
			value, _ := strconv.Atoi(line[2])
			bits := fmt.Sprintf("%036b", value)
			bitsP2 := fmt.Sprintf("%036b", keyInt)
			memP1[key] = bitMask(mask, bits)
			memP2 = decoder(mask, bitsP2, memP2, value)
		}
	}

	var p1 int64 = 0
	var p2 int = 0
	for _, v := range memP1 {
		p1 += v
	}
	for _, v := range memP2 {
		p2 += v
	}
	fmt.Println("P1:", p1)
	fmt.Println("P2:", p2)
}
