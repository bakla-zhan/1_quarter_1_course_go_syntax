package main

import (
	"fmt"
)

func even(num int) {
	if num%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}
}

func main() {
	var num int
	fmt.Println("Введите число")
	fmt.Scanln(&num)
	even(num)
}
