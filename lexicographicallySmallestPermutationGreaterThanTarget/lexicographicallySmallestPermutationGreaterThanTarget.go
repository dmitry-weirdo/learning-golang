package main

import (
	"fmt"
	"strings"
)

func lexGreaterPermutation(s string, target string) string {
	// Logically, we want the longest possible prefix to be equal to target.
	// After at some positions the permutation characters is greater than the target character,
	// we arrange the remaining characters in the ascending order to get the minimal string.

	// count the diff of counts of every character between s and target.
	counts := make([]int, 26)

	for i := range len(s) { // lengths of s and target are equal
		counts[s[i]-'a']++
		counts[target[i]-'a']--
	}

	//fmt.Printf("Counts of characters (S - T): %v \n", counts)

	// iterate the target from right to left
	for i := len(target) - 1; i >= 0; i-- {
		// remove the usage of this character
		char := target[i] - 'a'
		counts[char]++

		// check whether the remaining characters (i.e. the prefix before the current character)
		// have counts >= 0, i.e. whether we can match the prefix

		// Note that for the characters from current position to the end of target,
		// we already removed their usages.
		if !canUseExactPrefix(counts) {
			// cannot use the exact prefix -> get to the previous character in target
			continue
		}

		// We can match the prefix.
		// But we need to find the smallest still available (counts[i] > 0) character
		// greater than current, i.e. that can be used at the current position instead of the current characters
		for j := char + 1; j < 26; j++ {
			if counts[j] > 0 { // we found the character for the current position -> generate the result
				counts[j]-- // use this character at the current position

				// prefix characters will be used to 0 counts.
				// If counts[k] > 0, we still can use these characters in the suffix after the current pos.
				var sb = strings.Builder{}
				sb.WriteString(target[:i]) // write prefix before the current position
				sb.WriteByte('a' + j)      // append the found smallest greater char of the current position

				// append the remaining characters in lexicographical order.
				// every character is appended count[k] times.

				for k, v := range counts {
					for range v { // append count[k] times
						sb.WriteByte(byte('a' + k))
					}
				}

				return sb.String()
			}
		}

		// If we could not find the character for the current position -> go to the previous position in target.
	}

	// No match found -> s cannot not be permutated to a string > target.
	return ""
}

func canUseExactPrefix(counts []int) bool {
	for _, v := range counts {
		if v < 0 { // at least for one char, there are not enough characters in the initial string
			return false
		}
	}

	return true
}

func test(s, target string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)
	fmt.Printf("Target string: %v \n", target)

	result := lexGreaterPermutation(s, target)

	fmt.Printf("Lexicographically smallest permutation of \"%v\" that is greater than \"%v\": %v \n", s, target, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("abc", "bba", "bca")
}

func test2() {
	test("leet", "code", "eelt")
}

func test3() {
	test("baba", "bbaa", "")
}

func test4() {
	test("aaa", "aaa", "")
}

func main() {
	// 3720. Lexicographically Smallest Permutation Greater Than Target
	test1()
	test2()
	test3()
	test4()
}
