package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	student := Student{"asd", 1}
	studnet2 := Student{"wefji", 3}
	studnet2 = student
	student.Age = 5
	fmt.Println(student)
	fmt.Println(studnet2)
}