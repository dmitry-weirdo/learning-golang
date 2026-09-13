package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(sentence string) string {
	words := strings.Split(sentence, " ")

	var sb strings.Builder

	for i, v := range words {
		if startsWithVowel(v) { // starts with vowel -> append "ma" to the word
			sb.WriteString(v)
			sb.WriteString("ma")
		} else { // starts with consonant -> move first character to the end, then append "ma"
			sb.WriteString(v[1:])
			sb.WriteByte(v[0])
			sb.WriteString("ma")
		}

		// append 'a' letters (i + 1) times
		for range i + 1 {
			sb.WriteByte('a')
		}

		sb.WriteByte(' ')
	}

	s := sb.String()

	// remove the trailing space
	return s[:len(s)-1]
}

func startsWithVowel(s string) bool {
	switch s[0] {
	case 'a', 'e', 'o', 'u', 'i', 'A', 'E', 'O', 'U', 'I':
		return true
	default:
		return false
	}
}

func test(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := toGoatLatin(s)

	fmt.Printf("String in Goat Latin: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		"I speak Goat Latin",
		"Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
	)
}

func test2() {
	test(
		"The quick brown fox jumped over the lazy dog",
		"heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
	)
}

func main() {
	// 824. Goat Latin
	test1()
	test2()
}
