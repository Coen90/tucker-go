package exinit

import "fmt"

var (
	b = f("b")
	c = f("c")
	d = 3
)

func init() {
	d++
	fmt.Println("exinit.init function", d)
}

func f(arg string) int {
	d++
	fmt.Println("f(" + arg + ") d:", d)
	return d
}

func PrintD() {
	fmt.Println("d:", d)
}