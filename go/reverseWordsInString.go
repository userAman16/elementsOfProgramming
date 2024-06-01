package main

import (
	"fmt"
)

func ReverseWordsInString1() {
	input := "Alice messaged Bob"
	output := []rune{}
	i, j := len(input)-1, len(input)-1
	for j >= 0 {
		if input[j] == ' ' || j == 0 {
			x := i
			y := j + 1
			if j == 0 {
				y = j
			}
			for x >= y {
				output = append(output, rune(input[y]))
				y++
			}
			if j != 0 {
				output = append(output, ' ')
			}
			j--
			i = j
		}
		j--

	}
	fmt.Println(string(output))

}

func ReverseString(inp *[]rune, i, j int) {

	for ; i < j; i, j = i+1, j-1 {
		(*inp)[i], (*inp)[j] = (*inp)[j], (*inp)[i]
	}

}

func ReverseWordsInString2() {
	input := "Alice messaged Bob"
	inp := []rune(input)
	ReverseString(&inp, 0, len(inp)-1)
	i, j := 0, 0
	for ; j < len(inp); j++ {
		if inp[j] == ' ' {
			ReverseString(&inp, i, j-1)
			i = j + 1

		}
	}
	if j == len(inp) {
		ReverseString(&inp, i, j-1)
	}
	fmt.Println(string(inp))

}

func ReverseWordsInString() {
	//ReverseWordsInString1()
	ReverseWordsInString2()
}
