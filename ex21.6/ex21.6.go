package main

import "fmt"

func CaptureLoop1() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i :=0; i < 3; i ++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f:= make([]func(), 3)
	fmt.Println("ValueLoop3")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {fmt.Println(v)}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop1()
	CaptureLoop2()
}