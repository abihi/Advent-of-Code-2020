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

func main() {
	mem := make(map[string]int64)
	file, _ := os.Open("day14.in")
	scanner := bufio.NewScanner(file)
	mask := ""
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if line[0] == "mask" {
			mask = line[2]
			// fmt.Println(mask)
		} else {
			key := line[0][4 : len(line[0])-1]
			value, _ := strconv.Atoi(line[2])
			bits := fmt.Sprintf("%036b", value)
			mem[key] = bitMask(mask, bits)
			// fmt.Println("Key", key)
			// fmt.Println("Value", value)
			// fmt.Println("Bits", bits)
			// fmt.Println("mem[key]", mem[key], "key", key)
		}
	}

	var p1 int64 = 0
	for _, v := range mem {
		p1 += v
	}
	fmt.Println("P1:", p1)
}
