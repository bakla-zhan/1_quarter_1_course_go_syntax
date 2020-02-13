package main

import (
	"fmt"
)

func main() {
	var deposit float64
	var percent float64
	maxYearCount := 5
	fmt.Println("Укажите сумму, которую Вы хотели бы разместить на вкладе")
	fmt.Scanln(&deposit)
	fmt.Println("Укажите процент (%), который будет начисляться ежегодно (вклад с капитализацией)")
	fmt.Scanln(&percent)
	for yearCount := 1; yearCount <= maxYearCount; yearCount++ {
		deposit = deposit + (deposit * percent / 100)
		fmt.Printf("Размер вклада через %v год(-а)/лет составит %.2f рубль(-ей)\n", yearCount, deposit)
	}

}
