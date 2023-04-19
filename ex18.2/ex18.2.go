package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}

	slice2 := append(slice, 4, 5, 6)

	fmt.Println(slice)
	fmt.Println(slice2)

}