package main

import (
	"fmt"

	"main.com/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "Hello"

	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)

	dictionary.Delete(baseWord)

	word, err := dictionary.Search(baseWord)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}
