package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("%s, %d", u.Name, u.Age)
}

type VIPUser struct {
	User // Embeded field
	VIPLevel int
}

func(v VIPUser) vipLevel() int {
	return v.VIPLevel
}

func main() {
	var v VIPUser = VIPUser{User{"Hyuntae", 34}, 5}
	fmt.Println(v.String())
}