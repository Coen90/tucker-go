package main

import (
	"fmt"

	"github.com/tuckersGo/musthaveGo/ch16/expkg"

	"github.com/guptarohit/asciigraph"

	"go_project/usepkg/custompkg"

	"go_project/usepkg/exinit"
)

func main() {
	custompkg.PrintCustom()
	custompkg.Test()

	expkg.PrintSample()

	data  := []float64{3,4,5,6,7,8,7,5,8,5,2,2,4,7,4,5,6}

	graph := asciigraph.Plot(data)
	fmt.Println(graph)

	var student custompkg.Student = custompkg.Test3()
	fmt.Println(student)
	custompkg.Test3()

	exinit.PrintD()
}

