// Input merupakan variable string berisi kumpulan angka. Output merupakan list / array berisi angka yang hanya muncul 1 kali pada input.

package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	unik := make(map[int]bool)
	hasil := make([]int, 0)

	for _, i := range angka {
		num := int(i) - '0'
		if _, j := unik[num]; !j {
			unik[num] = true
		} else {
			unik[num] = false
		}
	}

	for num, nomorUnik := range unik {
		if nomorUnik {
			hasil = append(hasil, num)
		}
	}

	return hasil
}

func main() {
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("1122"))

}

// Arif Mardhatillah Asmar
