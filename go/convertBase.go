package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func ReverStringInGolang(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func BaseToBaseConversion(input string, b1, b2 int) string {

	decimal := 0
	for i := 0; i < len(input); i++ {

		decimal *= b1
		if unicode.IsDigit(rune(input[i])) {
			decimal += int(rune(input[i])) - int('0')
		} else {
			decimal += int(rune(input[i])) - int(rune('A')) + 10
		}

	}

	out := ""
	for decimal > 0 {
		temp := decimal % b2
		if temp > 9 {
			char := rune(65 + (temp - 10))
			out += string(char)
		} else {
			out += strconv.Itoa(temp)
		}

		decimal = int(decimal / b2)
	}
	return ReverStringInGolang(out)

}

func ConvertBase() {
	fmt.Println(BaseToBaseConversion("125", 10, 16))
	fmt.Println(BaseToBaseConversion("5A", 16, 2))
	fmt.Println(BaseToBaseConversion("10101", 2, 16))
}
