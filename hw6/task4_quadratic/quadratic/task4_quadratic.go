package quadratic

import (
	"fmt"
	"math"
)

// Quadratic функция для расчёта корней квадратного уравнения
func Quadratic(a, b, c float64) (x1, x2 float64, err error) {
	D := math.Pow(b, 2) - 4*a*c

	if D < 0 {
		return -1, -1, fmt.Errorf("квадратное уравнение не имеет корней")
	}

	x1 = (-b + math.Sqrt(D)) / 2 * a
	x2 = (-b - math.Sqrt(D)) / 2 * a

	return x1, x2, nil
}
