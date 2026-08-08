package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	s := "1"

	var sb strings.Builder

	for i := 2; i <= n; i++ {
		fmt.Printf("i: %v, len(s): %v \n", i, len(s))

		sb.Reset()

		char := s[0]
		charCount := 1

		for j := 1; j < len(s); j++ {
			if s[j] != char {
				if charCount > 3 { // never happens?
					panic("Char count > 3: " + strconv.Itoa(charCount))
				}

				sb.WriteString(strconv.Itoa(charCount))
				sb.WriteByte(char)

				char = s[j]
				charCount = 1
			} else {
				charCount++
			}
		}

		// also append the last char
		if charCount > 3 { // never happens?
			panic("Char count > 3: " + strconv.Itoa(charCount))
		}

		sb.WriteString(strconv.Itoa(charCount))
		sb.WriteByte(char)

		s = sb.String()
	}

	return s
}

func test(n int, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N: %v \n", n)

	result := countAndSay(n)

	fmt.Printf("CountAndSay(%v): %v \n", n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(1, "1")
}

func test2() {
	test(4, "1211")
}

func test3() {
	test(7, "13112221")
}

func main() {
	// 38. Count and Say
	// todo: is it proven that we'll never use any digits above 3?

	test1()
	test2()
	test3()

	// the time is O(4^n/3), so it's growing exponentially!
	// see https://leetcode.com/problems/count-and-say/editorial/comments/1721760/ for explanation of complexity

	//test(80, "-1")
	// n = 70 takes some time
	// n = 80 takes A LOT of time - i: 80, len(s): 1953245418 at the last step
}
