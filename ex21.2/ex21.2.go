package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Fail to create a file", err)
	}

	defer fmt.Println("반드시 호출됩니다.") // defer는 스택구조라 FILO
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다.")

	fmt.Println("파일에 Hello World를 씁니다.")
	fmt.Fprintln(f, "Hello World11")
}