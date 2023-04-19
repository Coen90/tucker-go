package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Float interface {
	float32 | float64
}

type Integer interface {
	int | int8 | int16 | int32 | int64
}

func print[T any](a T) {
	fmt.Println(a)
}

func min[T constraints.Ordered](a, b T) T {
	if a < b{
		return a
	}
	return b
}

func main() {
	var a int = 10
	fmt.Println(a)
	var b float32 = 3.14
	fmt.Println(b)
	var c string = "Hello"
	fmt.Println(c)

	var a1 int = 10
	var a2 int = 20
	fmt.Println(min(a1, a2))

	var a3 float32 = 10.123
	var a4 float32 = 10.1241
	fmt.Println(min(a3, a4))
}