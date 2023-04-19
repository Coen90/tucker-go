package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297
	fmt.Println("a:", a, "b:" ,b)
	fmt.Print("a:", a, "b:", b, "f:", f, "\n")
	fmt.Printf("a: %d b: %d f: %f\n", a, b, f)
	fmt.Printf("%v %v %v %T %T %T\n", a, b, f, a, b, f)
	fmt.Printf("%q\n", "Hello \"World\"")
	fmt.Printf("%v\n", "Hello \"World\"")
}