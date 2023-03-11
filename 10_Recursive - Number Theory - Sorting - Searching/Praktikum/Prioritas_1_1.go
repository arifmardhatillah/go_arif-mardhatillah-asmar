package main

import "fmt"

func Fibonacci(nomor int) int {
	if nomor == 0 {
		return 0
	}
	if nomor == 1 {
		return 1
	}
	return Fibonacci(nomor-1) + Fibonacci(nomor-2)
}

func main() {
	fmt.Println(Fibonacci(0))  // 0
	fmt.Println(Fibonacci(2))  // 1
	fmt.Println(Fibonacci(9))  // 34
	fmt.Println(Fibonacci(10)) // 55
	fmt.Println(Fibonacci(12)) // 144
}
