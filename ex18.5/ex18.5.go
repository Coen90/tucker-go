package main

import "fmt"

func addNum(slice []int) []int {
	slice = append(slice, 4)
	return slice
}

func main() {
	slice := []int{1, 2, 3}
	slice2 := addNum(slice)
	fmt.Println(slice2)
}