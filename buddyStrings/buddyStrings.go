package main

import "fmt"

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) { // different lengths
		return false
	}

	if s == goal { // strings are same -> we need a repeating characters to swap it in 2 positions
		freq := make([]int, 26)

		for _, ch := range s {
			freq[ch-'a']++

			if freq[ch-'a'] > 1 {
				return true
			}
		}

		// not a single repeating character found -> fail
		return false
	}

	// we need exactly 2 different indices and the characters at these indices should be vice versa in S and T
	index0 := -1
	index1 := -1

	for i := range s {
		if s[i] == goal[i] { // skip positions with same characters
			continue
		}

		if index0 == -1 { // 1st difference found
			index0 = i
		} else if index1 == -1 { // 2nd difference found
			index1 = i
		} else { // 3rd difference found -> fail
			return false
		}
	}

	if (index0 == -1) || (index1 == -1) { // No 2 different positions found -> fail
		return false
	}

	// at difference positions, the characters should be opposite in S and T
	return (s[index0] == goal[index1]) && (s[index1] == goal[index0])
}

func test(s, t string, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String S: %v \n", s)
	fmt.Printf("String T: %v \n", t)

	result := buddyStrings(s, t)

	fmt.Printf("Strings can be converted by swapping exactly 2 characters: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("ab", "ba", true)
}

func test2() {
	test("ab", "ab", false)
}

func test3() {
	test("aa", "aa", true) // we can swap [0] and [1]
}

func test4() {
	test("ab", "bac", false) // different lengths
}

func test5() {
	test("abd", "bae", false) // cannot swap 2 indices
}

func main() {
	// 859. Buddy Strings
	test1()
	test2()
	test3()
	test4()
	test5()
}
