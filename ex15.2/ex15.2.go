package main

import "fmt"

func main() {
	poet1 := "죽는 날까지 하늘을 우러러\n 한 점 부끄럼이 없기를"

	poet2 := `죽는 날까지 $a 하늘을 우러러
	한 점 부끄럼이 없기를`

	fmt.Println(poet1)
	fmt.Println(poet2)
}