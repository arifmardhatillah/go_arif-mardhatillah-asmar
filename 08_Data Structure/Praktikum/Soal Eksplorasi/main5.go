// Diberi matriks persegi, buatlah program untuk menghitung selisih absolut antara jumlah diagonalnya
// Diagonal kiri ke kanan : 1+5+9 = 15 . Diagonal kanan ke kiri : 3+5+9 = 17 . Perbedaan mutlak adalah |15 - 17| = 2.

package main

import (
	"fmt"
	"math"
)

func main() {
	totalDiagonal1 := 0
	totalDiagonal2 := 0
	matriks := [3][3]int{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}

	for i := 0; i < len(matriks); i++ {
		totalDiagonal1 += matriks[i][i]
	}

	for i := 0; i < len(matriks); i++ {
		totalDiagonal2 += matriks[i][len(matriks)-1-i]
	}

	fmt.Println(totalDiagonal1)
	fmt.Println(totalDiagonal2)
	fmt.Println("Total Diagonal =", int(math.Abs(float64(totalDiagonal1-totalDiagonal2))))
}

// Arif Mardhatillah Asmar
