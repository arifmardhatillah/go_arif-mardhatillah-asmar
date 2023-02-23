// Buatlah sebuah program untuk memeriksa apakah sebuah kata adalah palindrome atau bukan,
// serta coba terapkan scanner untuk menangkap inputan dari console.

package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Print("Masukkan sebuah kata: ")
	var kata string
	fmt.Scanln(&kata)

	kata = strings.ToLower(strings.ReplaceAll(kata, " ", ""))

	var balikKata string
	for i := len(kata) - 1; i >= 0; i-- {
		balikKata += string(kata[i])
	}

	if kata == balikKata {
		fmt.Printf("%s adalah palindrome", kata)
	} else {
		fmt.Printf("%s bukan palindrome", kata)
	}
}

// Arif Mardhatillah Asmar
