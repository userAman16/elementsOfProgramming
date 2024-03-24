package main

import (
	"fmt"
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

func ComputeXUponY() {
	//Primitive
	var dividend uint8 = 13
	var divisor uint8 = 7
	var quotient uint8 = 0
	var remainder uint8 = 0

	//Primitive(dividend, divisor, quotient, remainder)

	Advanced(dividend, divisor, quotient, remainder)
}
