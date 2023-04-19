package main

import "fmt"

type account struct {
	balance   int
	firstname string
	lastanme  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2  account) withdrawValue(amount int) {
	a2.balance -= amount
}

func main() {
	var mainA *account = &account{100, "joe", "park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)
	var mainB *account = &account{100, "joe", "park"}
	mainB.withdrawValue(30)
	fmt.Println(mainB.balance)
}