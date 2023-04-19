package main

import "fmt"

type StringHeader struct {
	Data uintptr
	Len int
}

func main() {
	var a string = "asdf"
	fmt.Println(&a)
	fmt.Println("Hello" > "Hell")
}