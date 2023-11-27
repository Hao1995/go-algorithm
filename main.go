package main

import "fmt"

func main() {
	idxMap := make(map[rune]int)
	idxMap['a'] = 1

	s := "abcde"
	fmt.Println(s[0:])
	fmt.Println(s[4:5])

	// print byte
	fmt.Println(s[0])
	// print rune
	fmt.Println(rune(s[0]))
}
