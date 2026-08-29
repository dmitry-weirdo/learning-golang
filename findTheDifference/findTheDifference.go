package main

import "fmt"

func findTheDifference(s string, t string) byte {
	xor := byte(0)

	// do one cycle instead of 2 on separate strings
	for i := range s {
		xor ^= s[i]
		xor ^= t[i]
	}

	// append the last character of t
	xor ^= t[len(t)-1]

	return xor
}

func test(s, t string, expectedResult byte) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String S: %v \n", s)
	fmt.Printf("String T (S with one added character): %v \n", t)

	result := findTheDifference(s, t)

	fmt.Printf("Character in T that was added to S: %c \n", result)
	fmt.Printf("Expected result: %c \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %c, actual result = %c \n", expectedResult, result)
	}
}

func test1() {
	test("abcd", "abcde", 'e')
}

func test2() {
	test("", "y", 'y')
}

func test3() {
	test("aa", "aaa", 'a')
}

func test4() {
	test("abc", "ceba", 'e')
}

func main() {
	// 389. Find the Difference
	// The same as "136. Single Number", just on characters
	test1()
	test2()
	test3()
	test4()
}
