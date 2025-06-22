package main

import "fmt"

func InsertionSort(input []int) []int {
	n := len(input)
	for i := 1; i < n; i++ {
		if input[i] < input[i-1] {
			y := i - 1
			for input[i] < input[y] && y >= 0 {
				y = y - 1
			}
			y = y + 1
			x := i
			for y < x {
				input[x], input[x-1] = input[x-1], input[x]
				x = x - 1
			}
		}

	}
	return input
}

func InsertionSortMain() {
	fmt.Println(InsertionSort([]int{2, 8, 5, 3, 9, 4}))
}
