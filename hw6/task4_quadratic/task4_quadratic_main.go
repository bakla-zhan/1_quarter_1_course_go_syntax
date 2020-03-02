package main

import (
	"1_Go_syntax/hw6/task4_quadratic/quadratic"
	"fmt"
)

func main() {
	x1, x2, err := quadratic.Quadratic(2.5, -5.5, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x1, x2)
	}
}
