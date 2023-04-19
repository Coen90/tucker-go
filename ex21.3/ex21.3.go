package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

// func getOperator(op string) OpFn {
// 	if op == "+" {
// 		return add
// 	} else if op == "*" {
// 		return mul
// 	} else {
// 		return nil
// 	}
// }

func getOperator(op string) OpFn {
	if op == "+" {
		return func (a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func (a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

type OpFn func(int, int) int

func main() {
	var operator OpFn
	operator = getOperator("+")
	var result = operator(1, 2)
	fmt.Println(result)

}