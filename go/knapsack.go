package main

import (
	"fmt"
	"sort"
)

var maxWeight int

type pair struct {
	value  int
	weight int
	ratio  int
}

type node struct {
	pairs    []pair
	value    int
	index    int
	estimate int
	weight   int
}

func getEstimate(i int, input []pair, weight int) int {
	l := len(input)

	estimateEval := 0
	tempMaxWeight := maxWeight - weight
	for ; i < l; i++ {
		p := input[i]
		fmt.Println(i, p.value, p.weight, p.ratio)
		if p.weight <= tempMaxWeight {
			tempMaxWeight = tempMaxWeight - p.weight
			estimateEval = estimateEval + p.value
		} else {
			for tempMaxWeight > 0 {
				estimateEval = estimateEval + p.ratio
				tempMaxWeight = tempMaxWeight - 1
			}
		}

	}
	return estimateEval

}

func Knapsack(estimate int, input []pair) {
	var maxNode node
	maxValue := 0
	s := NewStack[node]()

	s.Push(node{index: -1, estimate: estimate})

	for !s.IsEmpty() {
		ele, _ := s.Pop()
		if ele.estimate <= maxValue {
			continue
		}
		i := ele.index + 1
		if i >= len(input) {
			continue
		}

		// do not choose
		tempEle := ele
		tempEle.index = i
		tempEle.estimate = tempEle.value + getEstimate(tempEle.index+1, input, tempEle.weight)
		s.Push(tempEle)

		//choose next at i
		if ele.weight+input[i].weight <= maxWeight {
			tempNode := node{pairs: append(ele.pairs, input[i]), value: ele.value + input[i].value, weight: ele.weight + input[i].weight, index: i, estimate: ele.estimate}
			if ele.value+input[i].value > maxValue {
				maxNode = tempNode
				maxValue = ele.value + input[i].value
			}
			s.Push(tempNode)
		}

	}
	fmt.Println(maxNode)
}

func KnapsackMain() {

	/*var input map[int]int
	input = make(map[int]int)
	input[45] = 5
	input[48] = 8
	input[35] = 3*/
	input := map[int]int{
		60:  10, // Item 1
		100: 20, // Item 2
		120: 30, // Item 3
		50:  5,  // Item 4
		90:  8,  // Item 5
		70:  7,  // Item 6
	}

	maxWeight = 15

	pairs := []pair{}

	for k, v := range input {
		pairs = append(pairs, pair{value: k, weight: v, ratio: int(k / v)})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].ratio > pairs[j].ratio
	})
	estimateEval := 0
	tempMaxWeight := maxWeight
	for i, p := range pairs {
		fmt.Println(i, p.value, p.weight, p.ratio)
		if p.weight <= tempMaxWeight {
			tempMaxWeight = tempMaxWeight - p.weight
			estimateEval = estimateEval + p.value
		} else {
			for tempMaxWeight > 0 {
				estimateEval = estimateEval + p.ratio
				tempMaxWeight = tempMaxWeight - 1
			}
		}

	}
	fmt.Println(estimateEval)
	Knapsack(estimateEval, pairs)

}
