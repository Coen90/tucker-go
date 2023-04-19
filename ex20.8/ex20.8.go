package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
}

func (s *Student) String() string {
	return ""
}

func main() {
	var stringer Stringer
	fmt.Println(stringer.(*Student)) // 구현체가 없으면 에러 발생
}