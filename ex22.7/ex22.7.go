package main

import "fmt"

type MapTest struct {
	Name string
	Age int
}

func main() {
	m := make(map[int]int)
	m[1] = 0
	m[2] = 2
	m[3] = 3

	delete(m, 3)
	delete(m, 4)
	fmt.Println(m[3])
	fmt.Println(m[1])

	// v, ok := m[3]
	// fmt.Println(v, ok)
	// v, ok = m[1]
	// fmt.Println(v, ok)

	n := 1
	if v, ok := m[n]; ok {
		fmt.Println(v)
	} else {
		fmt.Println(n,"키는 비었습니다.")
	}

	m1 := MapTest{"hyuntae", 1}
	m2 := MapTest{"hyuntae", 2}
	mapTest := make(map[MapTest]int)
	mapTest[m1] = 1000
	mapTest[m2] = 2000

	fmt.Println(mapTest)
}