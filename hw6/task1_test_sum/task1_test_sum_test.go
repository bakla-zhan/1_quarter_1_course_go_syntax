package statistic

import "testing"

type testpair struct {
    values []float64
    result float64
}

var testsA = []testpair {
    {[]float64{1,2}, 1.5},
    {[]float64{1,1,1,1,1,1}, 1},
    {[]float64{-1,1}, 0},
}

var testsS = []testpair {
    {[]float64{1,2}, 3},
    {[]float64{1,1,1,1,1,1}, 6},
    {[]float64{-1,1}, 0},
}

func TestAverageSet(t *testing.T) {
    for _, pair := range testsA {
        v := Average(pair.values)
        if v != pair.result {
            t.Error(
                "For", pair.values, 
                "expected", pair.result,
                "got", v,
            )
        }
    }
}

func TestSummaSet(t *testing.T) {
    for _, pair := range testsS {
        v := Summa(pair.values)
        if v != pair.result {
            t.Error(
                "For", pair.values, 
                "expected", pair.result,
                "got", v,
            )
        }
    }
}