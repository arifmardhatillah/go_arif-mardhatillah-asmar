// Buatlah sebuah program untuk menghitung luas Trapesium
package main

import "fmt"

func main() {
	var a, b, t float64

	fmt.Print("Panjang Sisi Atas: ")
	fmt.Scanln(&a)

	fmt.Print("Panjang Sisi Bawah: ")
	fmt.Scanln(&b)

	fmt.Print("Masukkan Tinggi: ")
	fmt.Scanln(&t)

	luas := (a + b) * t / 2

	fmt.Println("Luas Trapesium:", luas)
}

// Arif Mardhatillah Asmar
