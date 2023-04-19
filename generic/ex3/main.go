package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Stringer2 interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
	String() string
}

type Integer interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

func PrintMin[T Stringer2](a, b T) {
	if a < b {
		fmt.Println(a.String())
	} else {
		fmt.Println(b.String())
	}
}

type MyInt int

func (m MyInt) String() string {
	return strconv.Itoa(int(m))
}

func Print1(a Stringer) { // Stringer Type
	fmt.Println(a.String())
}

func Print2[T Stringer](a T) { // 넘어온 타입
	fmt.Println(a.String())
}

type MyString struct {
	name string
}

func (m MyString) String() string {
	return m.name
}

func main() {
	m := MyString{"name"}
	Print1(m)
	Print2(m)

	var m1 MyInt = 10
	var m2 MyInt = 20
	PrintMin(m1, m2)
}