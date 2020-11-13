package client

import (
	"ch8/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
	//t.Log(series.square(3)) // undefined: series.square
	t.Log(series.Square(3))
}
