package main

import "fmt"

func main() {
	var word string = "palabra"
	var wordLen int = len(word)
	fmt.Printf("El largo de la palabra es %d\n", wordLen)
	for i := 0; i < wordLen; i++ {
		fmt.Printf("%c\n", word[i])
	}
}
