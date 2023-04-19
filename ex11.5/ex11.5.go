package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	for i := 0; i < 5; i++ {
		for j := 5; j > i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	for i := 0; i < 5; i++ {
		for j := 0; j < 4 - i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i * 2 + 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}