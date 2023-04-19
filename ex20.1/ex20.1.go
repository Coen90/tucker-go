package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해!", s.Age, s.Name)
}

func (s Student) GetAge() int {
	return s.Age
}

func main() {
	s := Student{"Hyuntae", 34}
	var stringer Stringer

	stringer = s
	fmt.Printf("%s\n", stringer.String())
}