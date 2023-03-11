package main

import "fmt"

func generateTriangel(numRows int) [][]int {
	triAngle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = triAngle[i-1][j-1] + triAngle[i-1][j]
		}
		triAngle[i] = row
	}
	return triAngle
}

func main() {
	fmt.Println(generateTriangel(5)) //[1], [1 1], [1 2 1], [1 3 3 1], [1 4 6 4 1]
}
