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
	fmt.Println(Fibonacci(0))
	fmt.Println(Fibonacci(1))
	fmt.Println(Fibonacci(2))
	fmt.Println(Fibonacci(3))
	fmt.Println(Fibonacci(5))
	fmt.Println(Fibonacci(6))
	fmt.Println(Fibonacci(7))
	fmt.Println(Fibonacci(9))
	fmt.Println(Fibonacci(10))
}
