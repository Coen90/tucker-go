package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Map[F, T any](s []F, f func(F) T) []T {
	rst := make([]T, len(s))
	for i, v := range s {
		rst[i] = f(v)
	}
	return rst
}

func main() {
	doubled := Map([]int{1, 2, 3}, func(v int) int {
		return v * 2
	})
	fmt.Println(doubled)

	uppered := Map([]string{"hello", "world", "abc"}, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Println(uppered)

	toString := Map([]int{1, 2, 3}, func (v int) string {
		return "str" + strconv.Itoa(v)
	})
	fmt.Println(toString)
}