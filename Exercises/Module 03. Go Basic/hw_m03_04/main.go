package main

import "fmt"

func countCharacters(sentence string) map[rune]int {
	characters := make(map[rune]int)
	for _, c := range sentence {
		characters[c]++
	}
	return characters
}

func main() {
	sentence := "съешь ещё этих мягких французских булок, да выпей чаю"
	for k, v := range countCharacters(sentence) {
		fmt.Printf("%q - %d\n", k, v)
	}
}
