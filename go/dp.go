package main

import "fmt"

func fibonacciIterative(n int) {
	if n == 0 {
		fmt.Println(n)

	} else if n == 1 {
		fmt.Println(0, 1)
	} else {
		a := 0
		b := 1
		fmt.Println(a)
		fmt.Println(b)
		for i := 2; i < n; i++ {
			c := a + b
			fmt.Println(c)
			a = b
			b = c

		}
	}
}

func fibonacciRecurssive(n int, na *[10]int) int {
	if n == 0 {
		na[0] = 0
		return 0
	}
	if n == 1 {
		na[1] = 1
		return 1
	}
	if na[n] != 0 {
		return na[n]
	}
	x := fibonacciRecurssive(n-1, na) + fibonacciRecurssive(n-2, na)
	na[n] = x
	return x
}

func findMaxSubArray() {
	//input := [8]int{-2, -3, 4, -1, -2, 1, 5, -3}
	input := [8]int{-2, -3, -4, -1, -2, -1, -5, -3}
	maxSum, currSum := input[0], input[0]
	i := 1
	for i < len(input) {
		if input[i]+currSum < 0 {
			currSum = input[i]
		} else {
			if input[i]+currSum < input[i] {
				currSum = input[i]
			} else {
				currSum = currSum + input[i]
			}
		}
		if maxSum < currSum {
			maxSum = currSum
		}
		i = i + 1
	}
	fmt.Println(maxSum)
}

func CoinCombRoutine(input *[3]int, n int, sum int) int {
	if sum == 0 {
		return 1
	}
	if sum < 0 {
		return 0
	}
	if n <= 0 {
		return 0
	}
	total := CoinCombRoutine(input, n, sum-input[n-1]) + CoinCombRoutine(input, n-1, sum)
	return total

}

func CoinComb() {
	var input = [3]int{1, 2, 3}
	var finalSum int = 5
	var len int = len(input)
	total := CoinCombRoutine(&input, len, finalSum)
	fmt.Println(total)
}

func LCSHelper(L1 string, L2 string, i int, j int, dp *[][]int) int {
	if i < 0 || j < 0 {
		return 0
	}

	if L1[i] == L2[j] {

		(*dp)[i][j] = 1 + LCSHelper(L1, L2, i-1, j-1, dp)
		return (*dp)[i][j]
	}

	(*dp)[i][j] = max(LCSHelper(L1, L2, i-1, j, dp), LCSHelper(L1, L2, i, j-1, dp))
	return (*dp)[i][j]

}

func LCS() {
	var L1 = "agngmtab"
	var L2 = "gxtxayb"
	i := len(L1)
	j := len(L2)
	dp := make([][]int, i)
	for x := 0; x < i; x++ {
		dp[x] = make([]int, j)
		for m := 0; m < j; m++ {
			dp[x][m] = -1
		}

	}
	out := LCSHelper(L1, L2, i-1, j-1, &dp)

	fmt.Println(out)

}

func DP() {
	//fibonacciIterative(10)
	//na := [10]int{-1}
	//fibonacciRecurssive(9, &na)
	//fmt.Println(na)
	//findMaxSubArray()
	//CoinComb()
	LCS()
}
