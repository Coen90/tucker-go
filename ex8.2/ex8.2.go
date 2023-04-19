package main

import "fmt"

func main() {
	const pi1 float64 = 3.141592653589793228 // 상수는 값을 변경할 수 없음 / 사용하지 않아도 에러 발생하지 않음
	var pi2 float64 = 3.141592653589793228
	// pi1 = 3
	pi2 = 4
	fmt.Println(pi2)
}