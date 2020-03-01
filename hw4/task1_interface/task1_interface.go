package main

import "fmt"

type vehicle interface {
	move()
}

func drive(vehicle vehicle) {
	vehicle.move()
}

type car struct{}
type aircraft struct{}

func (c car) move() {
	fmt.Println("Автомобиль едет")
}
func (a aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {
	tesla := car{}
	boeing := aircraft{}
	drive(tesla)
	drive(boeing)
}
