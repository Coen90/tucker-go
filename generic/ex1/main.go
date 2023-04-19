package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var a int = 10
	var b int = 20

	fmt.Println(min(a, b))

	var c int16 = 10
	var d int16 = 20
	fmt.Println(min(int(c), int(d)))

	var e float32 = 3.14
	var f float32 = 3.98
	fmt.Println(min(int(e), int(f)))
}