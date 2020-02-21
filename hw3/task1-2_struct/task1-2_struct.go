package main

import "fmt"

// описание структур
type engine struct {
	capacity   float32
	engineType string
	turbo      bool
}

type body struct {
	grossweight float32
	curbWeight  float32
	doorsAmount int
}

type car struct {
	carType string
	brand   string
	model   string
	year    int
	engine  // встроенные структуры
	body    // встроенные структуры
}

func main() {
	// инициализация структур
	var firstCar = car{
		carType: "Passenger",
		brand:   "Ford",
		model:   "Focus",
		year:    2018,
		engine: engine{
			capacity:   1.5,
			engineType: "Gasoline",
			turbo:      true,
		},
		body: body{
			grossweight: 1900,
			curbWeight:  1364,
			doorsAmount: 4,
		},
	}

	var secondCar = car{
		carType: "Passenger",
		brand:   "Mazda",
		model:   "3",
		year:    2019,
		engine: engine{
			capacity:   1.5,
			engineType: "Gasoline",
			turbo:      false,
		},
		body: body{
			grossweight: 1835,
			curbWeight:  1342,
			doorsAmount: 4,
		},
	}

	var thirdCar = car{
		carType: "Truck",
		brand:   "Kamaz",
		model:   "53504-50",
		year:    2019,
		engine: engine{
			capacity:   11.76,
			engineType: "Diesel",
			turbo:      true,
		},
		body: body{
			grossweight: 21400,
			curbWeight:  9025,
			doorsAmount: 2,
		},
	}

	// действия над структурами
	fmt.Println("Mazda не Ford?", firstCar != secondCar)

	if thirdCar.grossweight > secondCar.grossweight {
		fmt.Println(thirdCar.brand, "тяжелее", secondCar.brand)
	}

	if firstCar.year > secondCar.year {
		fmt.Println(secondCar.brand, "старше", firstCar.brand)
	} else {
		fmt.Println(firstCar.brand, "старше", secondCar.brand)
	}

	firstCar.year = 2017
	secondCar.year = 2018
	thirdCar.year = 2020

	// вывод значений свойств экземпляров в консоль
	fmt.Printf("%v %+v\n", firstCar.brand, firstCar)
	fmt.Printf("%v %+v\n", secondCar.brand, secondCar)
	fmt.Printf("%v %+v\n", thirdCar.brand, thirdCar)
}
