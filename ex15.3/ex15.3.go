package main

import "fmt"

func main() {
	str := "Hello 월드"

	fmt.Println(len(str)) // 바이트 길이 반환

	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("타입: %T 값: %d 문자값: %c\n", str[i], str[i], str[i])
	// }

	// arr := []rune(str)
	// // rune -> int32 별칭타입 4byte

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("타입: %T 값: %d 문자값: %c\n", arr[i], arr[i], arr[i])
	// }

	for _, v := range str {
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", v, v, v)
	}
}