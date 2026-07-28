package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	writeIndex := 0

	index := 0

	currentCharStartIndex := 0
	currentChar := chars[index]

	for index < len(chars) {
		index++

		fmt.Printf("Index updated to %v \n", index)

		endOfStringReached := index >= len(chars)

		if endOfStringReached || (chars[index] != currentChar) { // char changed
			if endOfStringReached {
				fmt.Println()
				fmt.Printf("Reached the end of the string. \n")
			} else {
				fmt.Println()
				fmt.Printf("Char changed from %c to %c \n", currentChar, chars[index])
			}

			count := index - currentCharStartIndex

			// append the previous char
			chars[writeIndex] = currentChar

			if count > 1 { // if count is 1, we don't append it
				countString := strconv.Itoa(count)

				countBytes := []byte(countString)

				for _, v := range countBytes {
					writeIndex++
					chars[writeIndex] = v
				}
			}

			fmt.Printf("Appended %v of char %c \n", count, currentChar)

			if !endOfStringReached {
				fmt.Printf("New chars: %v \n", string(chars))

				// change the current character
				writeIndex++
				currentChar = chars[index]
				currentCharStartIndex = index

				fmt.Printf("New writeIndex: %v \n", writeIndex)
				fmt.Printf("New char: %c \n", currentChar)
				fmt.Printf("New current char %c starts at %v \n", currentChar, currentCharStartIndex)
			}
		}
	}

	return writeIndex + 1 // length is 1 more than the index
}

func test(s string, expectedResult string) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Original string: %v \n", s)

	chars := []byte(s)
	compressedLength := compress(chars)

	compressedChars := chars[:compressedLength]
	compressedString := string(compressedChars)

	fmt.Println()
	fmt.Printf("String \"%v\" compressed to \"%v\". \n", s, compressedString)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if compressedString != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, compressedString)
	}
}

func test1() {
	s := "aabbccc"
	expectedResult := "a2b2c3"

	test(s, expectedResult)
}

func test2() {
	s := "a"
	expectedResult := "a"

	test(s, expectedResult)
}

func test3() {
	s := "aaabba"
	expectedResult := "a3b2a"

	test(s, expectedResult)
}

func main() {
	// 443. String Compression
	test1()
	test2()
	test3()
}
