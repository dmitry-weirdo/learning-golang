package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	return decodeString_recursive(s)
}

func decodeString_recursive(s string) string {
	var dfs func(t string) string

	dfs = func(t string) string {
		// string is k[internal_string]

		//fmt.Println()
		//fmt.Printf("Evaluating string \"%v\"... \n", t)

		var sb strings.Builder

		// first we parse k
		i := 0
		start := 0

		for i < len(t) {
			// skip letters, just add them to the result
			for i < len(t) && isLetter(t[i]) {
				sb.WriteByte(t[i])
				i++
			}

			if i >= len(t)-1 { // letters reached the end of the string -> nothing to search anymore
				break
			}

			// not the end of the string -> it must be the number now
			start = i

			for isDigit(t[i]) {
				i++
			}

			numberString := t[start:i]
			number, _ := strconv.Atoi(numberString)

			//fmt.Printf("Number found: %v. Searching for the next [substring] \n", number)

			start = i

			// after the number, it must be an expression within '[]' brackets. We're searching for the most outer brackets.
			if t[i] != '[' {
				panic(fmt.Sprintf("Next character after number %v must be '[', but is '%c'.", number, t[i]))
			}

			openingBracketsCount := 0
			matchingClosingBracketFound := false

			// search for the closing ']', but there can be more opening '[' brackets, so we need to close all external brackets first and then get to the closing of the external
			for !matchingClosingBracketFound {
				if t[i] == '[' {
					openingBracketsCount++
				} else if t[i] == ']' {
					openingBracketsCount--

					if openingBracketsCount == 0 {
						matchingClosingBracketFound = true
					}
				}

				i++
			}

			// i is after the closing ']' now -> skip it
			substringInBrackets := t[start+1 : i-1] // skip starting and closing []
			//fmt.Printf("Substring to evaluate: %v \n", substringInBrackets)

			evaluatedSubstring := dfs(substringInBrackets)

			//fmt.Printf("Substring \"%v\" evaluated to \"%v\". Appending it %v times. \n", substringInBrackets, evaluatedSubstring, number)

			for range number { // append the evaluated substring <number> times
				sb.WriteString(evaluatedSubstring)
			}
		}

		return sb.String()
	}

	return dfs(s)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isLetter(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func test(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String to evaluate: %v \n", s)

	result := decodeString(s)

	fmt.Printf("String \"%v\" evaluated to: %v \n", s, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("3[a]2[bc]", "aaabcbc")
}

func test2() {
	test("3[a2[c]]", "accaccacc")
}

func test3() {
	test("2[abc]3[cd]ef", "abcabccdcdcdef")
}

func test4() {
	test("aa2[bb]def", "aabbbbdef")
}

func main() {
	// 394. Decode String
	test1()
	test2()
	test3()
	test4()
}
