package main

import (
	"fmt"
	"math"
)

type Pair struct {
	a, b any
}

func TwoSum() {
	var outArr []Pair
	inputArr := []int{2, 6, 5, 8, 11}
	target := 14
	var targetMap map[int]int
	targetMap = make(map[int]int)

	for x := range inputArr {
		targetMap[inputArr[x]] = -1
	}

	for key, _ := range targetMap {
		val, ok := targetMap[int(math.Abs(float64(target-key)))]
		if ok && val == -1 {
			targetMap[key] = int(math.Abs(float64(target - key)))
		}
	}

	for key, val := range targetMap {
		if val != -1 {
			outArr = append(outArr, Pair{key, val})
		}
	}

	fmt.Println(outArr)
}
