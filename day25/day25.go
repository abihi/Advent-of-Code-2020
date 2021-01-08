package main

import (
	"fmt"
	"math"
)

func transform(x, lz int64) int64 {
	return int64(math.Remainder(math.Pow(float64(x), float64(lz)), 20201227))
}

func main() {
	// var card_pubkey int64 = 6270530
	// var door_pubkey int64 = 14540258
	var card_pubkey int64 = 5764801
	var door_pubkey int64 = 17807724

	var clz int64 = 0
	var value int64 = 0
	for transform(7, clz) != card_pubkey {
		value = transform(7, clz)
		fmt.Println(value)
		clz++
	}

	encryption_key := transform(door_pubkey, clz)
	fmt.Println("Encryption key:", encryption_key)
	//16311885
}
