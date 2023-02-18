//Buatlah sebuah program untuk mencetak segitiga asterik

package main

import "fmt"

func main() {
	var tinggi int = 5

	for i := 1; i <= tinggi; i++ {
		for j := 5; j >= i; j-- {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// Arif Mardhatillah Asmar
