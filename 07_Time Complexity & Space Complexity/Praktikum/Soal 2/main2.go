// **Sample Test Cases**
// Input : x = 2, n = 3
// Output : 8
// Input : x = 7, n = 2
// Output : 49

package main

import "fmt"

func powMod(x, n int) int {
	hasil := 1
	for n > 0 {
		if n&1 == 1 {
			hasil = (hasil * x)
		}
		x = (x * x)
		n >>= 1
	}
	return hasil
}

func main() {

	fmt.Println(powMod(2, 3))
	fmt.Println(powMod(7, 2))
	fmt.Println(powMod(5, 3))
	fmt.Println(powMod(10, 2))
	fmt.Println(powMod(2, 5))
	fmt.Println(powMod(7, 3))
}

// Arif Mardhatillah Asmar
