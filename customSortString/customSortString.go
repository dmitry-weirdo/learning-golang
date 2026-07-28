package main

import (
	"fmt"
	"strings"
)

func customSortString(order string, s string) string {
	var m = make(map[byte]int) // char to count

	for i := 0; i < len(s); i++ {
		ch := s[i]
		_, ok := m[ch]

		if !ok { // 1 occurrence -> put 1
			m[ch] = 1
		} else { // not 1st occurrence -> increase
			m[ch] = m[ch] + 1
		}
	}

	var charFoundInOrder = make(map[byte]bool)

	for k, _ := range m {
		charFoundInOrder[k] = false
	}

	var b strings.Builder

	for i := 0; i < len(order); i++ {
		ch := order[i]

		count, ok := m[ch]
		if ok {
			stringToAppend := strings.Repeat(string(ch), count)

			fmt.Printf("Char %v found %v times in the string. Appending string: %v \n", ch, count, stringToAppend)

			b.WriteString(stringToAppend)

			charFoundInOrder[ch] = true
		}
	}

	// append not found characters
	for ch, count := range m {
		if charFoundInOrder[ch] {
			continue
		}

		stringToAppend := strings.Repeat(string(ch), count)

		fmt.Printf("Char %v found %v times in the string and NOT found in ordered. Appending string: %v \n", ch, count, stringToAppend)

		b.WriteString(stringToAppend)
	}

	return b.String()
}

func main() {
	// 791. Custom Sort String
	order := "cba"
	s := "abcd"
	expected := "cbad"

	result := customSortString(order, s)

	fmt.Printf("Order: %v \n", order)
	fmt.Printf("String: %v \n", s)
	fmt.Printf("Result: %v \n", result)
	fmt.Printf("Expected result: %v \n", expected)
}
