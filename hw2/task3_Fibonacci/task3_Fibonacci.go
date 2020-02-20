package main

import (
	"fmt"
	"math/big"
)

func fibonacci() {
	f1 := big.NewInt(0)
	f2 := big.NewInt(1)
	for i := 1; i <= 100; i++ {
		f2.Add(f2, f1)
		f2, f1 = f1, f2
		fmt.Printf("%v-е число в последовательности Фибоначчи: %v\n", i, f2)
	}
}

func main() {
	fibonacci()
}
