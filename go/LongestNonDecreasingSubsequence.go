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

func LongestNonDecreasingSubsequenceMain() {
	var input = &[]int{10, 1, 9, 12, 2, 7, 13}
	fmt.Println(LongestNonDecreasingSubsequence(input, len(*input)))
	input = &[]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 10}
	fmt.Println(LongestNonDecreasingSubsequence(input, len(*input)))
	input = &[]int{8, 4, 12, 2, 10, 6, 14, 15}
	fmt.Println(LongestNonDecreasingSubsequence(input, len(*input)))
}
