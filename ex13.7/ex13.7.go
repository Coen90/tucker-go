package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 // 8 byte
	B int // 1 byte
	C int8 // 8 byte
	D int // 1 byte
	E int8 // 8 byte
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}