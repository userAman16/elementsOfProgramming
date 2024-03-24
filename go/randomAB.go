/*How would you implement a random number generator that generates
a random integer i in [a, b], given a random number generator that produces either zero or
one with equal probability? All generated values should have equal probability. What is the
run time of your algorithm?*/

//  i is the number of bits required to represent all possible values in the range [0, t-1]

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func zeroOneRandom() int {
	if rand.Intn(2) == 0 {
		return 0
	}
	return 1
}

func uniformRandomAB(a, b int) int {
	var t int = b - a + 1
	var res int = 0

	for i := 0; (1 << i) < t; i++ {
		res = (res << 1) | zeroOneRandom()
	}

	for res >= t {
		for i := 0; (1 << i) < t; i++ {
			res = (res << 1) | zeroOneRandom()
		}
	}

	return res + a
}

func randomAB() {
	rand.Seed(time.Now().UnixNano())
	a := 3
	b := 8
	result := uniformRandomAB(a, b)
	fmt.Printf("Random integer between %d and %d is: %d\n", a, b, result)
}
