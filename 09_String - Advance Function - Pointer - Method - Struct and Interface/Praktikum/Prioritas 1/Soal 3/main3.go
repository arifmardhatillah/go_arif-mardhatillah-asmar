package main

import (
	"fmt"
	"strings"
)

func compare(i, j string) string {

	if len(i) <= len(j) {
		if strings.Contains(j, i) {
			return i
		}
	} else {
		if strings.Contains(i, j) {
			return j
		}
	}
	return "INVALID"
}
func main() {
	fmt.Println(compare("AKA", "AKASHI"))
	fmt.Println(compare("KANGGOORO", "KANG"))
	fmt.Println(compare("KI", "KIJANG"))
	fmt.Println(compare("KUPU-KUPU", "KUPU"))
	fmt.Println(compare("ILALANG", "ILA"))
}
