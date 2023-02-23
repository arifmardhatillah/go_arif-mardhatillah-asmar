// Buatlah sebuah program untuk menentukan apakah sebuah Bilang adalah Bilang Ganjil atau Genap
package main

import "fmt"

func main() {
	var bil int

	fmt.Print("Bilangan: ")
	fmt.Scanln(&bil)

	if bil%2 == 0 {
		fmt.Println(bil, "Adalah Bilangan Genap")
	} else {
		fmt.Println(bil, "Adalah Bilangan Ganjil")
	}
}

// Arif Mardhatillah Asmar
