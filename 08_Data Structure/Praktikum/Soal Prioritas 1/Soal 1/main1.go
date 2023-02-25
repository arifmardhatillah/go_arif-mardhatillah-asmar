// Buatlah sebuah program menggabungkan 2 array yang diberikan,
// dan jangan sampai terdapat nama yang sama di data yang sudah tergabung tadi!

package main

import "fmt"

func mergeArrays(array1 []string, array2 []string) []string {

	merged := append(array1, array2...)
	namaMap := make(map[string]bool)
	hasil := make([]string, 0)

	for _, nama := range merged {
		if _, i := namaMap[nama]; !i {
			namaMap[nama] = true
			hasil = append(hasil, nama)
		}
	}

	return hasil
}

func main() {
	array1 := []string{"Arif", "Ridho", "Amek", "Oka", "Dimas"}
	array2 := []string{"Ikhsan", "Amek", "Rafli", "Oka", "Aqsha"}
	merged := mergeArrays(array1, array2)

	fmt.Println(merged)
}

// Arif Mardhatillah Asmar
