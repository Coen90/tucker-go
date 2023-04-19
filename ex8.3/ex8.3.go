package main

import "fmt"

const Pig int = 0
const Cow int = 1
const Chicken int = 2

const (
	Red int = 1 << iota
	Blue 
	Green
)

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	}
}

func main() {
	PrintAnimal(Cow)
	PrintAnimal(Pig)
	PrintAnimal(0)
	PrintAnimal(7)
	fmt.Println(Red)
	fmt.Println(Blue)
	fmt.Println(Green)
}