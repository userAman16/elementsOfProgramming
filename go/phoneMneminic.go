package main

import (
	"fmt"
)

func PhoneMnemonicHelper(num string, mapNum *[]string, d int, out string) {
	if d == len(num) {
		fmt.Println(out)
		return
	}
	//fmt.Println(reflect.TypeOf(num[d] - '0'))
	//fmt.Println(reflect.TypeOf(num[d] - '0').Kind())
	for _, c := range (*mapNum)[num[d]-'0'] {
		PhoneMnemonicHelper(num, mapNum, d+1, out+string(c))
	}
}

func PhoneMnemonic() {
	out := ""
	num := "07345"
	mapNum := []string{"0", "1", "ABC", "DEF", "GHI", "JKL", "MNO", "PQRS", "TUV", "WXYZ"}
	d := 0
	PhoneMnemonicHelper(num, &mapNum, d, out)

}
