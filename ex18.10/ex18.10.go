package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	// for문으로 중간 요소 삭제
	// for i := idx + 1; i < len(slice); i++ {
	// 	slice[i-1] = slice[i]
	// }
	// slice = slice[:len(slice)-1]

	// append로 중간 요소 삭제
	// slice = append(slice[:idx], slice[idx+1:]...)
	
	copy(slice[idx:], slice[idx+1:])


	fmt.Println(slice)
}