package quadratic

import (
	"fmt"
	"testing"
)

/*type testpair struct {
	values []float64
	results []float64
}

var tests = []testpair{
	{[]float64{2.5, -5.5, 3}, {7.5, 6.25, nil}},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}*/

func TestQuadratic(t *testing.T) {
	x1, x2, err := Quadratic(2.5, -5.5, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		if x1 != 7.5 && x2 != 6.25 {
			t.Error(
				"For x1 & x2 expected 7.5 & 6.25, got", x1, x2)
		}
	}
}
