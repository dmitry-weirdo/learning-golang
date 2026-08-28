package main

import (
	"fmt"
	"slices"
	"strings"
)

func lexPalindromicPermutation(s string, target string) string {
	// The search idea is similar to "3720. Lexicographically Smallest Permutation Greater Than Target"
	// However, we're only searching the first half of the palindrome, subtracting the character usage 2 times.

	// To be permutable to a palindrome, the string must have at most 1 character that has an odd frequency.
	// The single (non-paired) appearance of this character will appear in the center of the palindrome.
	// Palindrome is <left> + <optionalCenter> + <right> = <left> + <optionalCenter> + <reversedLeft>.

	// collect frequencies
	freq := make([]int, 26)
	for i := range len(s) {
		freq[s[i]-'a']++
	}

	// check whether S can be permutated to a palindrome
	center := ""
	for i, frequency := range freq {
		if frequency%2 == 1 { // odd number
			if center != "" { // there was already a odd-frequency character -> cannot permute S into a palindrome -> no solution
				return ""
			}

			center = string(byte('a' + i))
			freq[i]-- // subtract the center usage from the frequencies
		}
	}

	halfLength := len(s) / 2 // will exclude the center in case of odd length

	// reduce the counts of the target prefix characters of the first half by 2
	// (again similar to "3720. Lexicographically Smallest Permutation Greater Than Target")
	for i := range halfLength {
		char := target[i] - 'a'
		freq[char] -= 2
	}

	// if we already found the target palindrome with exact target first half -> no need to update the last character
	// but we need to check whether this palindrome is bigger than the target
	if canUseExactPrefix(freq) {
		firstHalf := target[:halfLength]
		secondHalf := reverseString(firstHalf)

		palindrome := firstHalf + center + secondHalf

		if palindrome > target {
			return palindrome
		}
	}

	// iterate the half of the target from right to left
	for i := halfLength - 1; i >= 0; i-- {
		// remove the usage of this character
		char := target[i] - 'a'
		freq[char] += 2

		// check whether the remaining characters (i.e. the prefix before the current character)
		// have counts >= 0, i.e. whether we can match the prefix
		if !canUseExactPrefix(freq) {
			// cannot use the exact prefix -> get to the previous character in target
			continue
		}

		// We can match the prefix.
		// But we need to find the smallest still available (counts[i] > 0) character
		// greater than current, i.e. that can be used at the current position instead of the current characters
		for j := char + 1; j < 26; j++ {
			if freq[j] > 0 { // we found the character for the current position -> generate the result
				freq[j] -= 2 // use this character at the current position

				// prefix characters will be used to 0 counts.
				// If counts[k] > 0, we still can use these characters in the suffix after the current pos.
				var sb = strings.Builder{}
				sb.WriteString(target[:i]) // write prefix before the current position
				sb.WriteByte('a' + j)      // append the found smallest greater char of the current position

				// append the remaining characters of the first half in lexicographical order.
				// every character is appended count[k] times.

				for k, v := range freq {
					for range v / 2 { // append count[k] / 2 times, it's just the first half
						sb.WriteByte(byte('a' + k))
					}
				}

				firstHalf := sb.String()
				secondHalf := reverseString(firstHalf)

				sb.WriteString(center)
				sb.WriteString(secondHalf)

				return sb.String()
			}
		}

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

func reverseString(s string) string {
	stringAsSlice := []byte(s)
	slices.Reverse(stringAsSlice) // reverses in place

	return string(stringAsSlice)
}

func test(s, target string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)
	fmt.Printf("Target string: %v \n", target)

	result := lexPalindromicPermutation(s, target)

	fmt.Printf("Smallest palindromic permutation of \"%v\" that is greater than \"%v\": %v \n", s, target, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("baba", "abba", "baab")
}

func test2() {
	test("baba", "bbaa", "") // baab < bbaa
}

func test3() {
	test("abc", "abb", "") // no palindromic permutations
}

func test4() {
	test("aac", "abb", "aca")
}

func main() {
	// 3734. Lexicographically Smallest Palindromic Permutation Greater Than Target
	test1()
	test2()
	test3()
	test4()
}
