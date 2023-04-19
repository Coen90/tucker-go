package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 26.9, 12.21, 123.5, 42.4}
	for i := 0; i < len(t); i++ {
		fmt.Println(t[i])
	}
}