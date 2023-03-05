package main

import "fmt"

func caesar(offset int, input string) string {
	output := []rune{}

	for _, character := range input {
		character = 'a' + (character-'a'+rune(offset))%26
		output = append(output, character)
	}

	return string(output)
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
