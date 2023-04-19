package main

import "fmt"

func main() {
	var t [5]int = [5]int{1, 2, 3, 4, 5}
	for i, v := range t {
		fmt.Println(i, v)
	}

	for i, v := range t {
		fmt.Println(i, v)
	}

}