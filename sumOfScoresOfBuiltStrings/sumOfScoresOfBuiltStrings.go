package main

import "fmt"

func sumScores(s string) int64 {

	// Z-array sum is calculated when calculating Z-array itself
	// O(n) time
	// runs _a bit_ faster than O(2 * n) version - I was able to execute it in 2-5 ms
	return sumScores_zArray_optimized(s)

	// todo: we can calculate the sums in the same iteration as calculating the Z-array, but it will require modification of the standard Z-algorithm
	// O(2*n) time - O(n) to calculate the z-array and O(n) to sum the sums
	// It passed in 5-6 ms
	//return sumScores_zArray(s)

	// Brute-force O(n^2) like in "214. Shortest Palindrome".
	// However, since we don't break fast on finding 1 solution,
	// it's falling on TLE.
	// Test-case 43 / 150 testcases passed
	// String length: 39239
	//return sumScores_bruteForce(s)
}

func sumScores_zArray_optimized(s string) int64 {
	if s == "" { // corner-case
		return 0
	}

	n := len(s)
	fmt.Printf("String length: %v \n", n)

	// modified version that returns just the sum of values in the Z-array
	zArraySum := calculateZArraySum(s)
	fmt.Printf("Z-array sum for string \"%v\": \n%v \n", s, zArraySum)

	return int64(len(s)) + zArraySum // complete string is always a match in this problem, but it is not calculated within the Z-array
}

func calculateZArraySum(s string) int64 {
	// see https://www.youtube.com/watch?v=CpZh4eF8QBw

	// For every s[i], we calculate z[i] = the length of the longest string starting at s[i]
	// that coincides with the prefix s[0:j] of s.

	// s[0] is not important and set to 0.

	// basically we calculate with a sliding window

	sum := int64(0)

	n := len(s)

	z := make([]int, n)
	z[0] = 0

	// i - index in the string for that we're calculating z[i]

	// sliding window borders
	// window can be called z-box
	left := 0
	right := 0

	for i := 1; i < n; i++ { // we calculate from [1] since z[0] is not important (it will be the whole string match) and set to 0
		if i > right { // [i] not within the window -> start to expand window from [i] while window matches the start of the string (i.e. matches the prefix)
			left = i
			right = i

			// s[right - left] is the character of the prefix
			for (right < n) && (s[right] == s[right-left]) {
				right++
			}

			// right is now the first non-matching character
			right--                 // move to the first matching character, can be (left - 1) if there were no matching characters
			z[i] = right - left + 1 // size of the matched window, can be 0
			sum += int64(z[i])
		} else {
			// [i] is within the window

			// how many prefix characters we currently matched at the [i] position
			prefixCharactersCount := i - left

			// Since the current window has matches with prefix,
			// we're trying to copy z[j] from the prefix if possible.
			// It is possible to copy if adding z[j] characters from s[i]
			// will end earlier that the end of the window (end of the window is right).

			// So we're trying to append z[j] characters starting from s[i], s[i] inclusive.
			zOfPrefixCharacter := z[prefixCharactersCount]

			if (i + zOfPrefixCharacter - 1) < right {
				// we can just copy z[i] = z[j] from the prefix
				z[i] = zOfPrefixCharacter
				sum += int64(z[i])
			} else {
				// We CANNOT copy from the prefix character, since the match touches or overflows the right border.
				// It means that we should try to match further after the right,
				// since this part can (is?) different from the furrther prefix characters.

				// try to expand the window starting from s[i]
				left = i

				// s[right - left] is the character of the prefix
				for (right < n) && (s[right] == s[right-left]) {
					right++
				}

				// right is now the first non-matching character
				right--                 // move to the first matching character, can be (left - 1) if there were no matching characters
				z[i] = right - left + 1 // size of the matched window, can be 0
				sum += int64(z[i])
			}
		}
	}

	return sum
}

func sumScores_zArray(s string) int64 {
	if s == "" { // corner-case
		return 0
	}

	n := len(s)
	fmt.Printf("String length: %v \n", n)

	// todo: we can calculate the sums in the same iteration as calculating the Z-array, but it will require modification of the standard Z-algorithm
	z := calculateZArray(s)
	fmt.Printf("Z-array for string \"%v\": \n%v \n", s, z)

	result := int64(len(s)) // complete string is always a match in this problem, but it is not calculated within the Z-array

	// We're summing up the Z-array values, since for every position, this is a start of a suffix string,
	// and we're searching for longest prefix length for every starting position.
	// This is exactly what is stored in the Z-array.
	for _, v := range z {
		result += int64(v)
	}

	return result
}

func sumScores_bruteForce(s string) int64 {
	if s == "" { // corner-case
		return 0
	}

	n := len(s)
	fmt.Printf("String length: %v \n", n)

	result := int64(0)

	for i := n - 1; i >= 0; i-- { // we include the whole string
		// this should not require copying of the substrings, so probably should be faster as reversing every prefix?
		suffix := s[n-i-1 : n] // build suffix substring

		// we're comparing all the prefix with length up to i - 1
		for j := len(suffix) - 1; j >= 0; j-- {
			suffixPrefix := suffix[0 : j+1]
			prefix := s[0 : j+1]

			if prefix == suffixPrefix {
				//fmt.Printf("Suffix string = \"%v\". Longest common prefix = \"%v\". Adding its length = %v to the result. \n", suffix, prefix, j + 1)

				result += int64(j + 1)
				break
			}
		}
	}

	return result
}

func allIndexesOf_zAlgorithm(s, p string, separatorChar string) []int { // separatorChar must be the char that is not present neither in string S nor in pattern P
	// Concat to "<pattern><separatorChar><string>",
	// So that the <pattern> string works as the prefix.

	// Then we calculate Z-indexes for this concatenated string.
	// Positions in the s[j] in the <string> part where z[j] = len(pattern)
	// will be the positions where the <pattern> is found.

	// To get the indexes in the original string (not the combined),
	// we have to subtract j - len(pattern) - len(separatorChar)

	combined := p + separatorChar + s

	z := calculateZArray(combined)

	fmt.Printf("Z-array of combined string \"%v\": \n%v \n", combined, z)

	matchIndexes := make([]int, 0)

	patternLength := len(p)
	subtractedLength := patternLength + len(separatorChar)

	for i := subtractedLength; i < len(combined); i++ { // we only need matches in the <string> part (although in pattern there will be no match with the complete length)
		if z[i] == patternLength {
			matchIndexes = append(matchIndexes, i-subtractedLength)
		}
	}

	return matchIndexes
}

func calculateZArray(s string) []int {
	// see https://www.youtube.com/watch?v=CpZh4eF8QBw

	// For every s[i], we calculate z[i] = the length of the longest string starting at s[i]
	// that coincides with the prefix s[0:j] of s.

	// s[0] is not important and set to 0.

	// basically we calculate with a sliding window

	n := len(s)

	z := make([]int, n)
	z[0] = 0

	// i - index in the string for that we're calculating z[i]

	// sliding window borders
	// window can be called z-box
	left := 0
	right := 0

	for i := 1; i < n; i++ { // we calculate from [1] since z[0] is not important (it will be the whole string match) and set to 0
		if i > right { // [i] not within the window -> start to expand window from [i] while window matches the start of the string (i.e. matches the prefix)
			left = i
			right = i

			// s[right - left] is the character of the prefix
			for (right < n) && (s[right] == s[right-left]) {
				right++
			}

			// right is now the first non-matching character
			right--                 // move to the first matching character, can be (left - 1) if there were no matching characters
			z[i] = right - left + 1 // size of the matched window, can be 0
		} else {
			// [i] is within the window

			// how many prefix characters we currently matched at the [i] position
			prefixCharactersCount := i - left

			// Since the current window has matches with prefix,
			// we're trying to copy z[j] from the prefix if possible.
			// It is possible to copy if adding z[j] characters from s[i]
			// will end earlier that the end of the window (end of the window is right).

			// So we're trying to append z[j] characters starting from s[i], s[i] inclusive.
			zOfPrefixCharacter := z[prefixCharactersCount]

			if (i + zOfPrefixCharacter - 1) < right {
				// we can just copy z[i] = z[j] from the prefix
				z[i] = zOfPrefixCharacter
			} else {
				// We CANNOT copy from the prefix character, since the match touches or overflows the right border.
				// It means that we should try to match further after the right,
				// since this part can (is?) different from the furrther prefix characters.

				// try to expand the window starting from s[i]
				left = i

				// s[right - left] is the character of the prefix
				for (right < n) && (s[right] == s[right-left]) {
					right++
				}

				// right is now the first non-matching character
				right--                 // move to the first matching character, can be (left - 1) if there were no matching characters
				z[i] = right - left + 1 // size of the matched window, can be 0
			}
		}
	}

	return z
}

func testCalculateZArray(s string, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("String: %v \n", s)

	result := calculateZArray(s) // todo: replace with your function

	fmt.Printf("Z-array of \"%v\": \n%v \n", s, result)
	fmt.Printf("Expected result: \n%v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func testZArray1() {
	s := "babab"
	expected := []int{0, 0, 3, 0, 1}

	testCalculateZArray(s, expected)
}

func testZArray2() {
	s := "abaxabab"
	expected := []int{0, 0, 1, 0, 3, 0, 2, 0}

	testCalculateZArray(s, expected)
}

func testZArray3() {
	s := "azbazbzaz"
	expected := []int{0, 0, 0, 3, 0, 0, 0, 2, 0}

	testCalculateZArray(s, expected)
}

func testZArray4() {
	s := "aabxaabxcaabxaabxay"
	expected := []int{0, 1, 0, 0, 4, 1, 0, 0, 0, 8, 1, 0, 0, 5, 1, 0, 0, 1, 0}

	testCalculateZArray(s, expected)
}

func testZArray5() {
	s := "abc$xabcabzabc"
	expected := []int{0, 0, 0, 0, 0, 3, 0, 0, 2, 0, 0, 3, 0, 0}

	testCalculateZArray(s, expected)
}

func testZArray6() {
	s := "aabxaayaab"
	expected := []int{0, 1, 0, 0, 2, 1, 0, 3, 1, 0}

	testCalculateZArray(s, expected)
}

func testZArray7() {
	s := "a"
	expected := []int{0}

	testCalculateZArray(s, expected)
}

func testZArraySuite() {
	testZArray1()
	testZArray2()
	testZArray3()
	testZArray4()
	testZArray5()
	testZArray6()
	testZArray7()
}

func testGetAllIndexes_zAlgorithm(s, p string, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("String: %v \n", s)
	fmt.Printf("Pattern to find in the string: %v \n", p)

	separatorChar := "#"
	result := allIndexesOf_zAlgorithm(s, p, separatorChar)

	fmt.Printf("All indexes of \"%v\" in \"%v\": %v \n", p, s, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func testZAlgorithm1() {
	//    012345678901
	s := "aabxaabaaaab"
	p := "aab"
	expected := []int{0, 4, 9}

	testGetAllIndexes_zAlgorithm(s, p, expected)
}

func testZAlgorithm2() {
	s := "aaa"
	p := "a"
	expected := []int{0, 1, 2}

	testGetAllIndexes_zAlgorithm(s, p, expected)
}

func testZAlgorithmSuite() {
	testZAlgorithm1()
	testZAlgorithm2()
}

func test(s string, expectedResult int64) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := sumScores(s)

	fmt.Printf("Sum of longest prefixes lengths between original string and all suffix substrings: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	// "b" -> prefix "b" -> 1
	// "bab" -> prefix "bab" -> 3
	// "babab" -> prefix "babab" -> 5
	test("babab", 9)
}

func test2() {
	// "az" -> prefix "az" -> 2
	// "azbzaz" -> prefix "azb" -> 3
	// "azbazbzaz" -> prefix "azbazbzaz" -> 9
	test("azbazbzaz", 14)
}

func main() {
	// 2223. Sum of Scores of Built Strings
	test1()
	test2()

	//testZArraySuite()
	//testZAlgorithmSuite()
}
