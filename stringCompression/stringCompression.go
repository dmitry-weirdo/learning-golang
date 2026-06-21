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

		if index >= len(chars) { // reached the end of the string
			fmt.Println()
			fmt.Printf("Reached the end of the string. \n")

			count := index - currentCharStartIndex

			// append the previous char
			chars[writeIndex] = currentChar

			if count > 1 { // if count is 1, we don't append it
				countString := strconv.Itoa(count)

				countBytes := []byte(countString)

				for j := 0; j < len(countBytes); j++ {
					writeIndex++
					chars[writeIndex] = countBytes[j]
				}
			}

			fmt.Printf("Appended %v of char %c \n", count, currentChar)

			break
		}

		if chars[index] != currentChar { // char changed
			fmt.Println()
			fmt.Printf("Char changed from %c to %c \n", currentChar, chars[index])

			count := index - currentCharStartIndex

			// append the previous char
			chars[writeIndex] = currentChar

			if count > 1 { // if count is 1, we don't append it
				countString := strconv.Itoa(count)

				countBytes := []byte(countString)

				for j := 0; j < len(countBytes); j++ {
					writeIndex++
					chars[writeIndex] = countBytes[j]
				}
			}

			fmt.Printf("Appended %v of char %c \n", count, currentChar)
			fmt.Printf("New chars: %v \n", string(chars))

			// change the current character
			writeIndex++
			currentChar = chars[index]
			currentCharStartIndex = index

			fmt.Printf("New writeIndex: %v \n", writeIndex)
			fmt.Printf("New char: %c \n", currentChar)
			fmt.Printf("New current char %c starts at %v \n", currentChar, currentCharStartIndex)
		} else {

		}
	}

	return writeIndex + 1 // length is 1 more than the index
}

func main() {
	var originalString string

	originalString = "aabbccc"
	originalString = "aabbccc"
	originalString = "a"
	originalString = "aaabba"

	fmt.Printf("Original string: \n%v \n", originalString)

	chars := []byte(originalString)
	compressedLength := compress(chars)

	compressedChars := chars[:compressedLength]
	compressedString := string(compressedChars)

	fmt.Printf("String \n%v \ncompressed to \n%v \n", originalString, compressedString)
}
