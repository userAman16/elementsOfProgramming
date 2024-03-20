package main

import (
	"fmt"
	"strconv"
)

func Parity() {
	var x uint64 = 125
	fmt.Println(strconv.FormatUint(x, 2))
}
