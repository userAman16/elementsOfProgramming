/*
design an algorithm that determines the maximum profit that
could have been made by buying and then selling a single share over a given day
range, subject to the constraint that the buy and the sell have to take place at the start
of the day. (This algorithm may be needed to backtest a trading strategy.)
*/

/*
Iterate through S, keeping track of the minimum element m seen thus far. If the
difference of the current element and m is greater than the maximum profit recorded
so far, update the maximum profit. This algorithm performs a constant amount of
work per array element, leading to an O(n) time complexity.
*/

package main

import "fmt"

func SharePrice() {
	inputArr := []int{8, 7, 10, 2, 6, 4, 3, 10}
	i, canMin := 0, 0
	min, max := inputArr[i], inputArr[i]
	for ; i < len(inputArr)-1; i++ {
		if (inputArr[i+1] - inputArr[i]) > (max - min) {
			max = inputArr[i+1]
			min = inputArr[i]
		}
		if min < inputArr[canMin] {
			canMin = i
		}
		if (inputArr[i+1] - inputArr[canMin]) > (max - min) {
			min = inputArr[canMin]
		}
	}

	fmt.Println("min : ", min, "max :", max)

}
