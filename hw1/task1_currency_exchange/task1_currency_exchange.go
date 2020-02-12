package main

import (
	"fmt"
)

func main() {
	const dollar = 63
	const euro = 69
	var rub float64
	fmt.Println("Сколько рублей Вы хотели бы обменять?")
	fmt.Scanln(&rub)
	var dollarAmount = rub / dollar
	var euroAmount = rub / euro
	fmt.Printf("За это количество рублей Вы получите %.2f \u0024 или %.2f \u20AC\n", dollarAmount, euroAmount)
}
