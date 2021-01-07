package main

import (
	"fmt"
)

func transform(subject_number int, loop_size int) int {
	value := 1
	for i := 0; i < loop_size; i++ {
		value *= subject_number
		value %= 20201227
	}
	return value
}

func calc_encryption_key(door_pubkey int, card_loop_size int, card_pubkey int, door_loop_size int) (int, int) {
	return transform(door_pubkey, card_loop_size), transform(card_pubkey, door_loop_size)
}

func main() {
	subject_number := 7
	card_pubkey := 6270530
	door_pubkey := 14540258
	var card_loop_size int
	var door_loop_size int

	for i := 0; i < 100; i++ {
		value := transform(subject_number, i)
		if value == card_pubkey {
			card_loop_size = i
			fmt.Println("Card loop size", i)
		}
		if value == door_pubkey {
			door_loop_size = i
		}
	}

	encryption_key1, encryption_key2 := calc_encryption_key(door_pubkey, card_loop_size, card_pubkey, door_loop_size)

	fmt.Println(encryption_key1, encryption_key2)
}
