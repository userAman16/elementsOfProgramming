package main

import (
	"fmt"
	"sort"
)

func LongestNonDecreasingSubsequence(input *[]int, l int) int {
	var dp = []int{}
	for i := 0; i < l; i++ {
		dp = append(dp, 1)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if (*input)[i] > (*input)[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	sort.Ints(dp)
	return dp[l-1]
}

func LongestNonDecreasingSubsequenceBinarySearch(input *[]int, l int) int {
	var tail = []int{}
	tail = append(tail, (*input)[0])
	for i := 1; i < l; i++ {
		if ((*input)[i]) > tail[len(tail)-1] {
			tail = append(tail, (*input)[i])
		} else {
			// search smallest elemnt in tail that is greater than or euql to (*input[i]) and replace with it
			start := 0
			end := len(tail) - 1
			for start < end {
				mid := start + (end-start)/2
				if tail[mid] < (*input)[i] {
					start = mid + 1
				} else {
					end = mid
				}
			}
			tail[start] = (*input)[i]
		}

	}
	return len(tail)
}

func LongestNonDecreasingSubsequenceMain() {
	var input = &[]int{10, 1, 9, 12, 2, 7, 13}
	fmt.Println(LongestNonDecreasingSubsequence(input, len(*input)))
	input = &[]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 10}
	fmt.Println(LongestNonDecreasingSubsequence(input, len(*input)))
	input = &[]int{8, 4, 12, 2, 10, 6, 14, 15}
	fmt.Println(LongestNonDecreasingSubsequence(input, len(*input)))

	fmt.Println("--------- Via Binar Search --------")
	input = &[]int{10, 1, 9, 12, 2, 7, 13}
	fmt.Println(LongestNonDecreasingSubsequenceBinarySearch(input, len(*input)))
	input = &[]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 10}
	fmt.Println(LongestNonDecreasingSubsequenceBinarySearch(input, len(*input)))
	input = &[]int{8, 4, 12, 2, 10, 6, 14, 15}
	fmt.Println(LongestNonDecreasingSubsequenceBinarySearch(input, len(*input)))
}
