package main

import "fmt"

type Database interface {
	Get() string
	Set(data string)
}

type CDatabase struct {
	data string
}

func (c *CDatabase) GetData() string {
	return c.data
}

func (c *CDatabase) SetData(data string) {
	c.data = data
}

type Wrapper struct {
	CDatabase
}

func (w *Wrapper) Get() string {
	return w.GetData()
}

func (w *Wrapper) Set(data string) {
	w.SetData(data)
}

func main() {
	var cdatabase CDatabase = CDatabase{""}
	var database Database
	database = &Wrapper{cdatabase}
	fmt.Println(database.Get())
	database.Set("hello")
	fmt.Println(database.Get())
}