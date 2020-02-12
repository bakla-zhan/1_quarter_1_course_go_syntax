package main

import (
	"fmt"
	"math"
)

func main() {
	var cathetus1 float64
	var cathetus2 float64
	var hypotenuse float64
	var perimeter float64
	var area float64
	fmt.Println("Укажите длину катета 1 в сантиметрах")
	fmt.Scanln(&cathetus1)
	fmt.Println("Укажите длину катета 2 в сантиметрах")
	fmt.Scanln(&cathetus2)
	hypotenuse = math.Hypot(cathetus1, cathetus2)
	perimeter = cathetus1 + cathetus2 + hypotenuse
	area = cathetus1 * cathetus2 / 2
	fmt.Printf("Длина гипотенузы прямоугольного треугольника равна %.2f см\n", hypotenuse)
	fmt.Printf("Длина периметра прямоугольного треугольника равна %.2f см\n", perimeter)
	fmt.Printf("Площадь прямоугольного треугольника равна %.2f см\u00B2\n", area)
}
