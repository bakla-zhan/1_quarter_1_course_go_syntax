package main

import (
	"fmt"
)

func even(num int) {
	if num%3 == 0 {
		fmt.Println("Число ", num, " делится без остатка на 3")
	} else {
		fmt.Println("Число ", num, " НЕ делится без остатка на 3")
	}
}

func main() {
	var num int
	fmt.Println("Введите число")
	fmt.Scanln(&num)
	even(num)
}
