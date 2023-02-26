// Sample Test Case
// Input: 1000000007
// Output: Bilangan Prima
// Input: 1500450271
// Output: Bilangan Prima

package main

import "fmt"

func primeNumber(num int) bool {

	if num <= 1 {
		return false
	}
	for v := 2; v < num; v++ {
		if num%v == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(1500450271))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))

}

// Arif Mardhatillah Asmar
