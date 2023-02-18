// Buatlah sebuah program untuk menentukan Grade dari sebuah Nilai, dengan ketentuan sebagai berikut:
// - Nilai 80 - 100: A
// - Nilai 65 - 79: B
// - Nilai 50 - 64: C
// - Nilai 35 - 49: D
// - Nilai 0 - 34: E
// - Nilai kurang dari 0 atau lebih dari 100 maka tampilkan 'Nilai Invalid'

package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan Nilai: ")
	fmt.Scanln(&nilai)

	if nilai <= 100 && nilai >= 80 {
		fmt.Println("Nilai yang diperoleh = A")

	} else if nilai <= 79 && nilai > 65 {
		fmt.Println("Nilai yang diperoleh = B")

	} else if nilai <= 64 && nilai >= 50 {
		fmt.Println("Nilai yang diperoleh = C")

	} else if nilai <= 49 && nilai >= 35 {
		fmt.Println("Nilai yang diperoleh = D")

	} else if nilai <= 34 && nilai >= 0 {
		fmt.Println("Nilai yang diperoleh = E")

	} else {
		fmt.Println("Nilai Invalid")
	}
}

// Arif Mardhatillah Asmar
