package main

import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	split := strings.Split(sentence, " ")

	for i, s := range split {
		if strings.HasPrefix(s, searchWord) {
			return i + 1
		}
	}

	return -1
}

func test(sentence string, searchWord string, expected int) {
	fmt.Println()
	fmt.Println("========================")

	result := isPrefixOfWord(sentence, searchWord)

	fmt.Printf("sentence: %v \n", sentence)
	fmt.Printf("search word: %v \n", searchWord)
	fmt.Printf("result: %v \n", result)
	fmt.Printf("expected result: %v \n", expected)
}

func test1() {
	sentence := "i love eating burger"
	searchWord := "burg"

	test(sentence, searchWord, 4)
}

func test2() {
	sentence := " "
	searchWord := "test"

	test(sentence, searchWord, -1)
}

func test3() {
	sentence := "i am tired"
	searchWord := "you"

	test(sentence, searchWord, -1)
}

func main() {
	test1()
	test2()
	test3()
}
