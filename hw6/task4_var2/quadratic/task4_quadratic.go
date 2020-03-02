package quadratic

import (
	"fmt"
	"math"
)

// Results результаты вычисления функции
type Results struct {
	x1  float64
	x2  float64
	err error
}

// Quadratic функция для расчёта корней квадратного уравнения
func Quadratic(a, b, c float64) (res Results) {
	D := math.Pow(b, 2) - 4*a*c

	if D < 0 {
		res.x1 = -1
		res.x2 = -1
		res.err = fmt.Errorf("квадратное уравнение не имеет корней")
		return
	}

	res.x1 = (-b + math.Sqrt(D)) / 2 * a
	res.x2 = (-b - math.Sqrt(D)) / 2 * a
	res.err = nil
	return
}
