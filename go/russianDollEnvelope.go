package main

import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {

	//sort the envelopes
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	var dp []int
	for i := 0; i < len(envelopes); i++ {
		dp = append(dp, 1)
	}
	for i := 0; i < len(envelopes); i++ {
		for j := 0; j < len(envelopes); j++ {
			if i == j {
				continue
			}
			if (envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1]) && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	sort.Ints(dp)
	return dp[len(envelopes)-1]
}

func RussianDollEnvelope() {
	fmt.Println("RussianDollEnvelope")
	envelopes := [][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}}
	fmt.Println(maxEnvelopes(envelopes))
}
