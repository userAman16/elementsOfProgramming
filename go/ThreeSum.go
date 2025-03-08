package main

import (
	"fmt"
	"sort"
)

type threeComb struct {
	a, b, c int
}

func ThreeSum(input []int) {

	var pairSlice []Pair
	var out []threeComb
	var outArr [][]int
	for i, val := range input {
		pairSlice = append(pairSlice, Pair{a: i, b: val})
	}

	sort.Slice(pairSlice, func(i, j int) bool { return (pairSlice[i].b).(int) < (pairSlice[j].b).(int) })
	fmt.Println(pairSlice)

	n := len(pairSlice)
	for i := 0; i < n-2; i++ {
		if i > 0 && pairSlice[i].b.(int) == pairSlice[i-1].b.(int) {
			continue
		}
		fmt.Println(pairSlice[i].b.(int), i)

		x := i + 1
		y := n - 1
		for x < y {
			sum := pairSlice[i].b.(int) + +pairSlice[x].b.(int) + pairSlice[y].b.(int)

			if sum == 0 {
				out = append(out, threeComb{a: pairSlice[i].b.(int), b: pairSlice[x].b.(int), c: pairSlice[y].b.(int)})
				tempSlice := []int{pairSlice[i].a.(int), pairSlice[x].a.(int), pairSlice[y].a.(int)}
				sort.Ints(tempSlice)
				outArr = append(outArr, tempSlice)
				x++
				y--

				for x < y && pairSlice[x].b == pairSlice[x-1].b {
					x = x + 1
				}

				for x < y && pairSlice[y].b == pairSlice[y+1].b {
					y = y - 1
				}
			} else if sum < 0 {
				x = x + 1
			} else {
				y = y - 1
			}
		}

	}
	fmt.Println(out)

	fmt.Println(outArr)
}

func ThreeSumMain() {
	ThreeSum([]int{0, -1, 2, 4, -4, -1, 2, 0, 0})
	ThreeSum([]int{0, -1, 2, -3, 1})

}
