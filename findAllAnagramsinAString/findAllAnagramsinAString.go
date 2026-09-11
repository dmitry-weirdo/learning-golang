package main

import "fmt"

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	// count required frequencies
	f := make([]int, 26)
	for _, ch := range p {
		f[ch-'a']++
	}

	freq := make([]int, 26)

	result := make([]int, 0)

	left := 0

	charactersMatched := 0

	for right := 0; right < len(s); right++ {
		chRight := s[right]

		chRightIndex := chRight - 'a'
		freq[chRightIndex]++

		if freq[chRightIndex] <= f[chRightIndex] {
			// we're still increasing the required char -> increase count
			charactersMatched++
		}

		if (right - left + 1) > len(p) { // shrink from left
			chLeft := s[left]
			left++

			chLeftIndex := chLeft - 'a'
			freq[chLeftIndex]--

			if freq[chLeftIndex] < f[chLeftIndex] {
				// decreased the required char below the required count -> decrease count
				charactersMatched--
			}
		}

		//fmt.Println()
		//fmt.Printf("Left: %v, right: %v, substring [%v; %v] = \"%v\". Chars matched: %v \n", left, right, left, right, s[left:right+1], charactersMatched)

		if charactersMatched == len(p) { // exact match of all characters in P -> add index to the result
			result = append(result, left)
		}
	}

	return result
}

func test(s, p string, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("S: %v \n", s)
	fmt.Printf("P: %v \n", p)

	result := findAnagrams(s, p)

	fmt.Printf("Indexes of P anagrams in S: %v \n", result)
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
	test(
		"cbaebabacd",
		"abc",
		[]int{0, 6},
	)
}

func test2() {
	test(
		"abab",
		"ab",
		[]int{0, 1, 2},
	)
}

func main() {
	// 438. Find All Anagrams in a String
	test1()
	test2()
}
