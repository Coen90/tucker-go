package custompkg

import "fmt"

type Student struct {
	Name string
	Age int
	number int
}

func PrintCustom() {
	fmt.Println("This is custom package!")
}
func printCustom2() {
	fmt.Println("This is custom package!2222222222222")
}

func Test2() {
	student := new(Student)
	var name string = student.Name
	var age int = student.Age
	var number int = student.number
	fmt.Println(name, age, number)
}