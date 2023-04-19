package main

import "fmt"

type MyInt int

func(m MyInt) Add(a int) MyInt {
	rst := int(m) + a
	return MyInt(rst)
}

func main() {
	var a MyInt
	a = a.Add(10)
	fmt.Printf("%d\n", a)
}