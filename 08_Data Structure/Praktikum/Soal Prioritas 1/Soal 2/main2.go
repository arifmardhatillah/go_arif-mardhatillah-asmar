// Buatlah sebuah program yang dapat menghitung berapa banyak sebuah string yang sama didalam sebuah slice!

package main

import "fmt"

func menghitungString(i string, slice []string) int {
	count := 0
	for _, j := range slice {
		if j == i {
			count++
		}
	}

	return count
}

func main() {
	j := []string{"Arif", "Amek", "ikhsan", "Ridho", "Amek", "Oka", "Amek", "Rafli", "Amek"}
	i := "Amek"
	hitung := menghitungString(i, j)

	fmt.Printf("%s Muncul Sebanyak %d Kali Dalam Slice\n", i, hitung)
}

// Arif Mardhatillah Asmar
