package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for i := 0; i < n+1; i++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}
	fmt.Println()

	for i := 0; i < n+1; i++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}
}