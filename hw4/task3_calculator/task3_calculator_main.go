package main

import (
	"1_Go_syntax/hw4/task3_calculator/calculator"
	"1_Go_syntax/hw4/task3_calculator/help"
	"fmt"
)

func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}

		// ДЗ задание 3. Вызов справки к калькулятору
		if input == "help" {
			help.Help()
			continue
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
