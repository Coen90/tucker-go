package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	return s.Name
}

type User struct {
	Name string
}

func (u User) String() string {
	return u.Name
}

func main() {
	var s Stringer = &Student{"Hyuntae", 34}
	fmt.Println(s.String())
	var s1 Stringer = User{"username"}
	fmt.Println(s1.String())
}