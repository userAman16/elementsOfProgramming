package main

import (
	"fmt"
	"strconv"
)

func ConvertBase() {
	var inputStr string = "11101"
	var outStr string = ""
	var b1 int = 2
	var b2 int = 16
	var x int = 0
	for i := 0; i < len(inputStr); i++ {
		x *= b1
		x += int(inputStr[i]) - int('0')
	}
	fmt.Println(x)
	for x > 0 {
		var r int = x % b2
		if r > 10 {
			outStr += string(rune(int('A') + (r - 10)))
		} else {
			outStr += strconv.Itoa(r)
		}
		x = x / b2

	}
	fmt.Println(outStr)

}
