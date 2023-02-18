//Buatlah sebuah program untuk mencetak faktor bilangan dari sebuah angka

package main

import "fmt"

func main() {
	var angka int = 26

	fmt.Printf("Faktor dari %d adalah: ", angka)

	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}

// Arif Mardhatillah Asmar
