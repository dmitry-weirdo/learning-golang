package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 { // empty string or string of length 1 are themselves non-repeating
		return len(s)
	}

	maxLength := 1

	// character to its last index
	m := make(map[byte]int)

	left := 0
	right := 0

	for right < len(s) {
		fmt.Println()
		fmt.Printf("Left: %v, right: %v. Substring [%v; %v] = \"%v\". \n", left, right, left, right, s[left:right+1])

		chRight := s[right]

		if v, ok := m[chRight]; ok {
			if v < left { // old index was before the current range -> update this index to right
				fmt.Printf("Character '%c' was before the current range [%v; %v]. Updating its lastIndex to right = %v.\n", chRight, left, right, right)
				m[chRight] = right
			}
		} else { // value doesn't yet exist -> add its index
			m[chRight] = right
		}

		if m[chRight] == right {
			// no duplicate chars -> check whether current string is longer than max
			currentLength := right - left + 1

			if currentLength > maxLength {
				maxLength = currentLength

				fmt.Printf("Non-repeating string [%v, %v] = \"%v\" of is a new max length %v. \n", left, right, s[left:right+1], currentLength)
			}

			right++
		} else {
			fmt.Printf("Found repeating character '%c' at right of [%v, %v] = \"%v\". Moving left from [%v] to last index of '%c' + 1 = [%v]. \n", chRight, left, right, s[left:right+1], left, chRight, m[chRight]+1)

			// duplicate char added -> move left after the lastIndex of this char
			left = m[chRight] + 1

			fmt.Printf("Moved left to [%v]. Current non-repeating substring: [%v, %v] = \"%v\". \n", left, left, right, s[left:right+1])

			// last index of chRight is now right
			m[chRight] = right

			// we still move to the next position after jumping left
			right++
		}
	}

	return maxLength
}

func lengthOfLongestSubstringWithFrequencyMap(s string) int {
	if len(s) < 2 { // empty string or string of length 1 are themselves non-repeating
		return len(s)
	}

	maxLength := 1

	// character to its frequency
	m := make(map[byte]int)

	left := 0
	right := 0

	for right < len(s) {
		chRight := s[right]

		if v, ok := m[chRight]; ok { // value exists -> increase existing frequency
			m[chRight] = v + 1
		} else { // value doesn't yet exist -> add frequency 1
			m[chRight] = 1
		}

		if m[chRight] == 1 {
			// no duplicate chars -> check whether current string is longer then max
			currentLength := right - left + 1

			if currentLength > maxLength {
				maxLength = currentLength

				fmt.Printf("Non-repeating string [%v, %v] = \"%v\" of is a new max length %v. \n", left, right, s[left:right+1], currentLength)
			}

			right++
		} else {
			fmt.Printf("Found repeating character '%c' at right of [%v, %v] = \"%v\". Moving left from [%v] until there are no duplicates of '%c'... \n", chRight, left, right, s[left:right+1], left, chRight)

			// duplicate char added -> move left until this character is decreased to 1
			for m[chRight] > 1 {
				chLeft := s[left]
				m[chLeft]--

				left++
			}

			fmt.Printf("Moved left to [%v]. Current non-repeating substring: [%v, %v] = \"%v\". \n", left, left, right, s[left:right+1])

			// we still move to the next position after jumping left
			right++
		}
	}

	return maxLength
}

func test(s string, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("string: %v \n", s)

	result := lengthOfLongestSubstring(s)

	fmt.Printf("Result: %v  \n", result)
	fmt.Printf("Expected result: %v  \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "abcabcbb"
	expected := 3

	test(s, expected)
}

func test2() {
	s := "pwwkew"
	expected := 3

	test(s, expected)
}

func test3() {
	s := "pwwkewqwertyuiop"
	expected := 10

	test(s, expected)
}

func main() {
	test1()
	test2()
	test3()
}
