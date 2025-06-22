package main

import (
	"fmt"
	"math"
)

func Primitive(dividend, divisor, quotient, remainder uint8) {
	for dividend >= divisor {
		dividend = dividend - divisor
		quotient++
	}
	remainder = dividend
	fmt.Println("Quotient : ", quotient, "remainder : ", remainder)

}

func Advanced(dividend, divisor, quotient, res uint8) {
	for dividend >= divisor {
		var power int = 1
		for (divisor<<power) >= (divisor<<(power-1)) && ((divisor << power) <= dividend) {
			power++

		}
		res += 1 << (power - 1)
		dividend -= divisor << (power - 1)
	}

	fmt.Println("Quotient : ", res, "remainder : ", dividend)
}

func binarySearchLargestK(n int, divisor, dividend uint8) int {
	start := 0
	end := n
	result := 0
	for start <= end {
		mid := (end-start)/2 + start
		if (divisor << mid) <= dividend {
			result = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return result
}

func AdvancedBinarySearch(dividend uint8, divisor uint8) {
	quotient := 0
	for dividend >= divisor {
		// search largest k
		k := binarySearchLargestK(int(math.Log2(float64(dividend))), divisor, dividend)
		fmt.Println(k)

		quotient += 1 << k
		dividend -= divisor << k

	}
	fmt.Println("quotient", quotient, "remainder", dividend)
}

func ComputeXUponY() {
	//Primitive
	/* var dividend uint8 = 13
	var divisor uint8 = 7
	var quotient uint8 = 0
	var remainder uint8 = 0 */

	//Primitive(dividend, divisor, quotient, remainder)

	//Advanced(dividend, divisor, quotient, remainder)
	//AdvancedBinarySearch(dividend, divisor)
	AdvancedBinarySearch(39, 5)
}
