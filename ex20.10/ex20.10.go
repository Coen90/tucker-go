package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {

}

func (f *File) Close() {

}

func ReadFile(reader Reader) {
	c, ok := reader.(Closer)
	fmt.Println(c, ok)
	if ok {
		c.Close()
	}
}

func main() {
	file := &File{}
	ReadFile(file)
}