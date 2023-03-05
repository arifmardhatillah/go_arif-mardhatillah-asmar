package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Score []int
}

func main() {
	students := make([]Student, 0)

	for i := 1; i <= 5; i++ {
		var name string
		fmt.Printf("Masukkan nama siswa ke-%d: ", i)
		fmt.Scanln(&name)

		var scores []int
		for j := 1; j <= 1; j++ {
			var score int
			fmt.Printf("Masukkan skor siswa %s: ", name)
			fmt.Scanln(&score)
			scores = append(scores, score)
		}

		student := Student{Name: name, Score: scores}
		students = append(students, student)
	}

	var totalScore int
	var minScore, maxScore int = 100, 0
	var minStudent, maxStudent string

	for _, student := range students {
		var sum int
		for _, score := range student.Score {
			sum += score
		}
		avg := sum / len(student.Score)

		totalScore += avg

		if avg < minScore {
			minScore = avg
			minStudent = student.Name
		}

		if avg > maxScore {
			maxScore = avg
			maxStudent = student.Name
		}
	}

	avgScore := totalScore / len(students)

	fmt.Printf("Average Score: %d\n", avgScore)
	fmt.Printf("Min Score of Students: %s (%d)\n", minStudent, minScore)
	fmt.Printf("Max Score of Students: %s (%d)\n", maxStudent, maxScore)
}
