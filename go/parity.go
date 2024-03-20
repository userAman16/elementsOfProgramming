package main

import (
	"fmt"
	"strconv"
)

func Parity() {
	var xx [5]uint64 = [5]uint64{125, 11, 20, 24, 13}
	for _, x := range xx {
		fmt.Println(strconv.FormatUint(x, 2))
		x ^= x >> 32
		x ^= x >> 16
		x ^= x >> 8
		x ^= x >> 4
		x ^= x >> 2
		x ^= x >> 1

		fmt.Println(x & 1)
		if x&1 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}

}
