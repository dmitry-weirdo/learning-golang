package main

import "fmt"

func detectCapitalUse(word string) bool {
	if len(word) < 2 {
		return true
	}

	firstCapital := isCapital(word[0])
	secondCapital := isCapital(word[1])

	// first non-capital, second capital -> incorrect
	if !firstCapital && secondCapital {
		return false
	}

	// first and second capital -> all should be capital
	// else all should be non-capital
	otherMustBeCapital := firstCapital && secondCapital

	// todo: this can be optimized in if (otherMustBeCapital) around and 2 different for-s. But it passes in 0ms even with this code
	for i := 2; i < len(word); i++ {
		if otherMustBeCapital && !isCapital(word[i]) {
			return false
		} else if !otherMustBeCapital && isCapital(word[i]) {
			return false
		}
	}

	return true
}

func isCapital(b byte) bool {
	return ('A' <= b) && (b <= 'Z')
}

func test(s string, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := detectCapitalUse(s)

	fmt.Printf("Valid capital letters: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("USA", true)
}

func test2() {
	test("leetcode", true)
}

func test3() {
	test("Google", true)
}

func test4() {
	test("FlaG", false)
}

func test5() {
	test("FLAg", false)
}

func test6() {
	test("flaG", false)
}

func test7() {
	test("aB", false)
}

func test8() {
	test("a", true)
}

func test9() {
	test("A", true)
}

func main() {
	// 520. Detect Capital
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
	test9()
}
