package quadratic

import (
	"reflect"
	"testing"
)

type Expected struct {
	x1  float64
	x2  float64
	err error
}

func TestQuadratic(t *testing.T) {

	exp := Expected{
		x1:  7.5,
		x2:  6.25,
		err: nil,
	}

	res := Quadratic(2.5, -5.5, 3)
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("got %v, expected %v", res, exp)
	}
}
