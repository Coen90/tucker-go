package main

import "fmt"

func main() {
	// var slice []int

	slice := []int{1, 2, 3}
	if len(slice) == 0 {
		fmt.Println("Slice is empty", slice)
	}

	slice[1] = 10
	fmt.Println(slice)

	var slice2 []int = make([]int, 3)
	slice2[1] = 5

}