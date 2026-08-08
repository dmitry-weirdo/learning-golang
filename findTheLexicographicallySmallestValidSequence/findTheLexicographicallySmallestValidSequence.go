package main

import "fmt"

func validSequence(word1 string, word2 string) []int {
	// Find last indexes of every character of word2 in word1, including next correct sequences.
	// If not found -> set -1
	lastIndexes := make([]int, len(word2))

	//for i := range lastIndexes {
	//	lastIndexes[i] = -1
	//}

	i := len(word1) - 1 // position in word1
	j := len(word2) - 1 // position in word2

	for j = len(word2) - 1; j >= 0; j-- {
		for (i >= 0) && (word1[i] != word2[j]) {
			i--
		}

		if i < 0 { // character no found and reached the start of word1
			break
		}

		lastIndexes[j] = i
		i-- // go to prev character in word1
	}

	if j > 0 { // small optimization -> set -1 only before the first index that wasn't found, not the complete array
		for j >= 0 {
			lastIndexes[j] = -1
			j--
		}
	}

	fmt.Printf("Indexes of latest char occurrences of \"%v\" in \"%v\": %v \n", word2, word1, lastIndexes)

	// Greedily select characters of word2 when going from left in word1
	// Apply the 1-letter-change early as possible (this will minimize the indexes),
	// but we only use the change if:
	// - Change is not yet used
	// AND
	// - This is the last character of word2
	// OR
	// - The next character in word2 has a lastIndex[j] that is AFTER the current position in word1
	changeUsed := false

	result := make([]int, len(word2))

	j = 0 // position in word2

	for i = 0; i < len(word1); i++ {
		if j >= len(word2) { // reached the end of word2 -> success
			break
		}

		if word1[i] == word2[j] { // character matches -> just use it
			result[j] = i // save the index
			j++

			continue
		}

		// character is different -> apply the change if possible
		if !changeUsed &&
			((j == len(word2)-1) || // last character in word2
				(lastIndexes[j+1] > i)) { // we will meet all next characters later in word1
			fmt.Printf("Used change on word1[%v] = %c instead of word2[%v] = %c. \n", i, word1[i], j, word2[j])

			changeUsed = true
			result[j] = i
			j++
		}

		// if using change is not possible -> we will just skip the character in word[i]
	}

	if j >= len(word2) {
		// we processed all chars in word2 -> success, return found indexes
		return result
	}

	// cannot find subsequence in word1 -> return an empty array
	return []int{}
}

func test(a, b string, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("String 1 (containing string): %v \n", a)
	fmt.Printf("String 2 (substring to search): %v \n", b)

	result := validSequence(a, b)

	fmt.Printf("Indexes of minimum subsequence (max 1 change): %v \n", result)
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

func test1() {
	s1 := "vbcca"
	s2 := "abc"
	expected := []int{0, 1, 2}

	test(s1, s2, expected)
}

func test2() {
	s1 := "bacdc"
	s2 := "abc"
	expected := []int{1, 2, 4}

	test(s1, s2, expected)
}

func test3() {
	s1 := "aaaaaa"
	s2 := "aaabc"
	expected := []int{}

	test(s1, s2, expected)
}

func test4() {
	s1 := "abc"
	s2 := "ab"
	expected := []int{0, 1}

	test(s1, s2, expected)
}

func test5() {
	// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14
	// b b e i g i i b h j a  f  j  i  g

	// lastIndexes: [-1, -1, -1, 8, 12]

	s1 := "bbeigiibhjafjig"
	s2 := "iihhj"
	expected := []int{3, 5, 6, 8, 9}

	test(s1, s2, expected)
}

func main() {
	// 3302. Find the Lexicographically Smallest Valid Sequence
	test1()
	test2()
	test3()
	test4()
	test5()
}
