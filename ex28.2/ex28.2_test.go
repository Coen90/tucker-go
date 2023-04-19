package main

import "testing"

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

func BenchmarkFibonaccifibonacciMemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciMemo(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
} 