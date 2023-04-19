package main

import "fmt"

// 구조체 포인터 초기화
type Data struct {
	key int
	value int
}
func main() {
	var data Data
	var p *Data = &data

	var po *Data = &Data{}

	fmt.Println(*p, *po)
}