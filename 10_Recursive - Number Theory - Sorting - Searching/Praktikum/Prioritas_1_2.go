package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	counts := make(map[string]int)
	for _, item := range items {
		counts[item]++
	}

	var result []pair
	for key, value := range counts {
		result = append(result, pair{key, value})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].count > result[j].count
	})

	return result
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
