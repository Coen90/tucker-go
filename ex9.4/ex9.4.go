package main

import "fmt"

func UploadFile(result bool) (string, bool) {
	return "filename", result
}

func main() {
	if _, success := UploadFile(false); success {
		fmt.Println("성공!")
	} else {
		fmt.Println("실패!")
	}
}

