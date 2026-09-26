package main

import (
	"fmt"
	"strings"
)

func evaluate(s string, knowledge [][]string) string {
	// fill paramName -> paramValue map
	m := make(map[string]string)

	for _, v := range knowledge {
		m[v[0]] = v[1]
	}

	//fmt.Printf("Dictionary map: %v \n", m)

	var sb strings.Builder

	i := 0

	for i < len(s) {
		if s[i] == '(' {
			start := i

			for s[i] != ')' {
				i++
			}

			paramName := s[start+1 : i] // exclude the brackets
			//fmt.Printf("Parameter found: \"%v\" \n", paramName)

			if paramValue, ok := m[paramName]; ok { // param value found in dictionary
				sb.WriteString(paramValue)
			} else { // param value NOT found in dictionary -> append "?"
				sb.WriteByte('?')
			}
		} else { // not within brackets -> append to result as is
			sb.WriteByte(s[i])
		}

		i++
	}

	return sb.String()
}

func test(s string, m [][]string, expectedResult string) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)
	fmt.Printf("Parameters dictionary: %v \n", m)

	result := evaluate(s, m)

	fmt.Printf("String with replaced parameters: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		"(name)is(age)yearsold",
		[][]string{
			{"name", "bob"},
			{"age", "two"},
		},
		"bobistwoyearsold",
	)
}

func test2() {
	test(
		"hi(name)",
		[][]string{
			{"a", "b"},
		},
		"hi?", // parameter (name) not found in the dictionary
	)
}

func test3() {
	test(
		"(a)(a)(a)aaa",
		[][]string{
			{"a", "yes"},
		},
		"yesyesyesaaa",
	)
}

func main() {
	// 1807. Evaluate the Bracket Pairs of a String
	test1()
	test2()
	test3()
}
