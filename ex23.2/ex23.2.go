package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, fmt.Errorf("제곱근은 양수여야 합니다 f:%g", f)
		return 0, errors.New("제곱근은 양수여야 합니다")
	}
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Sqrt(-2)= %v\n", sqrt)
}