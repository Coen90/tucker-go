package main

import (
	"fmt"
	"os"
)

type Wirter func(string)

func WriteHello(writer Wirter) {
	writer("Heelo World")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	WriteHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}