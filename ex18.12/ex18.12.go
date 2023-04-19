package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{4,3,2,6,1}
	sort.Ints(slice)
	fmt.Println(slice)
}