package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"박현태", "hyuntae", 34}
	vip1 := VIPUser{user, 1, 1000}
	vip2 := VIPUser{User{"남단미", "danmi", 34}, 5, 5000}

	fmt.Println(vip1)
	fmt.Println(vip2)
}