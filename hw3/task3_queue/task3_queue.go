package main

import (
	"fmt"
)

var x []string

// функция добавления новых элементов
func push(str string) {
	x = append(x, str)
}

// функция возврата первого элемента
func pop() string {
	numOfElements := len(x)
	if numOfElements == 0 {
		return "очередь пуста"
	}
	popElem := x[0]
	x = append(x[1:])
	return popElem
}

func main() {
	push("1_первый пошёл")
	push("2_второй пошёл")
	push("3_третий пошёл")

	fmt.Println(pop())
	fmt.Println(pop())

	push("4_четвёртый пошёл")

	fmt.Println(pop())
	fmt.Println(pop())

	fmt.Println(pop())
}
