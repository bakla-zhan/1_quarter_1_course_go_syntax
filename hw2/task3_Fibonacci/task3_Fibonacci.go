package main

import (
	"fmt"
	//"math/big"
)

func fibonacci() {
	var f1 int
	var f2 int
	var f3 int
	for i := 1; i <= 100; i++ {
		if i == 1 {
			f1 = 0
			fmt.Printf("%v-е число в последовательности Фибоначчи: %v\n", i, f1)
		} else if i == 2 {
			f2 = 1
			fmt.Printf("%v-е число в последовательности Фибоначчи: %v\n", i, f2)
		} else if i < 94 {
			f3 = f2 + f1
			f1 = f2
			f2 = f3
			fmt.Printf("%v-е число в последовательности Фибоначчи: %v\n", i, f3)
		} /*else {
			f3 = big.Int(f3)
			f3 = f2 + f1; f1 = f2; f2 = f3; fmt.Printf("%v-е число в последовательности Фибоначчи: %v\n", i, f3)
		}*/
	}
}

func main() {
	fibonacci()
}
