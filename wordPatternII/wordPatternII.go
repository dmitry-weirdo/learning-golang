package main

import "fmt"

func wordPatternMatch(pattern string, s string) bool {
	m := len(pattern)
	n := len(s)

	if n < m {
		return false
	}

	// pattern to substring mapping
	p := make(map[string]string)

	// used substrings - to control that we don't use the same substring for different pattern characters
	// we use the reverse (substring -> pattern character) map
	substringToPatternCharacter := make(map[string]string)

	var dfsInner func(i, j int) bool

	dfsInner = func(i, j int) bool {
		// i - start index in pattern
		// j - start index in string

		fmt.Println()
		fmt.Printf("DFS started: patternIndex = %v, stringIndex = %v \n", i, j)
		fmt.Printf("Pattern to substring map: %v \n", p)
		fmt.Printf("Used substrings: %v \n", substringToPatternCharacter)

		// base case - we reached the end of both substrings -> success
		if i >= m && j >= n {
			fmt.Printf("Successfully reached end of both pattern and string. Returning true. \n")
			return true
		}

		// base case for prune - if we reached the end of just one substring -> return false
		if i >= m {
			fmt.Printf("Reached end of the pattern but did not reach the end of the string. Returning false. \n")
			return false
		}

		if j >= n {
			fmt.Printf("Reached end of the string but did not reach the end of the pattern. Returning false. \n")
			return false
		}

		// base case for prune - remaining part of the string is shorter than the remaining characters in the pattern
		if (n - j) < (m - i) {
			fmt.Printf("Length of the remaining substring \"%v\" = %v is less than the remaining pattern characters of pattern \"%v\" = %v", s[j:], n-j, pattern[i:], m-i)
			return false
		}

		patternCharacter := pattern[i : i+1]
		fmt.Printf("Pattern[%v] character = '%v' \n", i, patternCharacter)

		// try all possible substrings from s[j; j+1) to s[j; n)
		for endIndex := j + 1; endIndex <= n; endIndex++ {
			substring := s[j:endIndex] // endIndex is non-inclusive, so we will pass endIndex to the next dfs iterations

			fmt.Printf("Possible substring: %v \n", substring)

			// There are 4 cases:
			// patternChar already present + substring match -> do NOT add the new mapping, continue DFS with next positions
			// patternChar already present + substring not match -> fail, prune this substring
			// patternChar not present + substring is mapped to another pattern character -> fail, prune this substring
			// patternChar not present  + substring is NOT mapped to another pattern character -> add new mapping, continue DFS with next position

			// !!!! we're not returning false, we're just skipping the handling of this substring
			if patternCharacterMappedTo, ok := p[patternCharacter]; ok {
				fmt.Printf("Pattern character '%v' is already mapped to \"%v\". \n", patternCharacter, patternCharacterMappedTo)

				if patternCharacterMappedTo == substring {
					// substring matches the current mapping for the current character -> continue dfs without changes
					fmt.Printf("Substring \"%v\" correctly matches to pattern character '%v' = \"%v\". Continuing DFS with patternIndex = %v, stringIndex = %v. \n", substring, patternCharacter, patternCharacterMappedTo, i+1, endIndex)

					nextDfsResult := dfsInner(i+1, endIndex) // endIndex is NOT inclusive, so do NOT pass endIndex + 1
					if nextDfsResult == true {               // !!! we do NOT return false, we will just prune handling the current substring
						return true
					} else {
						continue
					}
				} else {
					fmt.Printf("Substring \"%v\" does NOT match the pattern character '%v' = \"%v\". Pruning handling of this substring. \n", substring, patternCharacter, patternCharacterMappedTo)
					continue
				}
			} else { // pattern character is not yet present
				fmt.Printf("Pattern character '%v' is not yet mapped to any substring. \n", patternCharacter)

				if patternCharacterForThisSubstring, ok := substringToPatternCharacter[substring]; ok {
					// substring is already mapped to another pattern character -> prune handling of this substring
					fmt.Printf("Substring \"%v\" is already mapped to pattern character '%v'. It cannot be mapped again to a different new pattern character = '%v'. Pruning handling of this substring. \n", substring, patternCharacterForThisSubstring, patternCharacterMappedTo)

					continue
				} else { // successful substring (not yet mapped to another pattern character) -> add new mapping, continue DFS
					p[patternCharacter] = substring
					substringToPatternCharacter[substring] = patternCharacter

					fmt.Printf("Pattern character '%v' mapped to \"%v\". Continuing DFS with patternIndex = %v, stringIndex = %v. \n", patternCharacter, substring, i+1, endIndex)

					nextDfsResult := dfsInner(i+1, endIndex) // endIndex is NOT inclusive, so do NOT pass endIndex + 1
					if nextDfsResult == true {               // !!! we do NOT return false, we will just prune handling the current substring
						return true
					}

					fmt.Println()
					fmt.Printf("DFS(%v, %v). Returned from non-successful inner DFS(%v, %v). \n", i, j, i+1, endIndex)
					fmt.Printf("Removing p[%v] = \"%v\" and s[%v] = \"%v\". \n", patternCharacter, p[patternCharacter], substring, substringToPatternCharacter[substring])

					// remove the handled (and non-successful) mappings
					delete(p, patternCharacter)
					delete(substringToPatternCharacter, substring)
				}
			}
		}

		// no valid mapping from the current i and j positions
		return false
	}

	return dfsInner(0, 0)
}

func test(pattern, s string, expectedResult bool) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Pattern: %v \n", pattern)
	fmt.Printf("String: %v \n", s)

	result := wordPatternMatch(pattern, s)

	fmt.Println()
	fmt.Printf("String \"%v\" is convertible a pattern \"%v\": %v \n", s, pattern, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	pattern := "abab"
	s := "redblueredblue"
	expected := true

	test(pattern, s, expected)
}

func test2() {
	pattern := "aaaa"
	s := "asdasdasdasd"
	expected := true

	test(pattern, s, expected)
}

func test3() {
	pattern := "aabb"
	s := "xyzabcxzyabc"
	expected := false

	test(pattern, s, expected)
}

func test4() {
	pattern := "aabb"
	s := "xyxyzz"
	expected := true

	test(pattern, s, expected)
}

func test5() {
	pattern := "a"
	s := "x"
	expected := true

	test(pattern, s, expected)
}

func test6() {
	pattern := "ab"
	s := "redred"
	expected := true // will actually work with a = "r", b = "edred"

	test(pattern, s, expected)
}

func test7() {
	pattern := "aba"
	s := "redbluered"
	expected := true

	test(pattern, s, expected)
}

func test8() {
	pattern := "ab"
	s := "aa"
	expected := false

	test(pattern, s, expected)
}

func main() {
	// 291. Word Pattern II
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
}
