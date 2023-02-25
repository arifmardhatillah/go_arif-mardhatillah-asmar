// Diberi array angka yang diurutkan dan jumlah target, temukan pasangan dalam array yang jumlahnya sama dengan target yang diberikan.
// Tulis fungsi untuk mengembalikan indeks dari dua angka (yaitu pasangan)
// sehingga jika value pada index tersebut dijumlahkan sesuai target yang diberikan.

package main

import "fmt"

func pairSum(array []int, target int) (int, int) {
	i := 0
	j := len(array) - 1

	for i < j {
		sum := array[i] + array[j]
		if sum == target {
			return i, j
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

	return -1, -1
}

func main() {
	fmt.Println(pairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(pairSum([]int{2, 5, 9, 11}, 11))
}

// Arif Mardhatillah Asmar
