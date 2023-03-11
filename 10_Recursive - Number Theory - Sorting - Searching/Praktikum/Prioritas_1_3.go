package main

import "fmt"

func isPrime(nomor int, i int) bool {
	if i == 1 {
		return true
	}

	if nomor%i == 0 {
		return false
	}

	return isPrime(nomor, i-1)
}

func getNthPrime(nomor int, candidate int) int {
	if isPrime(candidate, candidate-1) {
		if nomor == 1 {
			return candidate
		}
		return getNthPrime(nomor-1, candidate+1)
	}
	return getNthPrime(nomor, candidate+1)

}

func primeX(nomor int) int {
	return getNthPrime(nomor, 2)
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
