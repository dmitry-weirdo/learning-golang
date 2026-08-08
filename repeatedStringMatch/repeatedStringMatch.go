package main

import (
	"fmt"
	"strings"
)

func repeatedStringMatch(a string, b string) int {
	return repeatedStringMatch_trivial(a, b)
}

func repeatedStringMatch_trivial(a string, b string) int {
	ac := a

	count := 1

	// make first string not shorter than the 2nd string
	for len(ac) < len(b) {
		ac = ac + a
		count++
	}

	// b should be contained in max (ac + ac) -> check up to it
	limit := 2 * len(ac)

	//for len(ac) <= limit { // exactly 2 repeats should also be checked
	for len(ac) < limit { // exactly 2 repeats should also be checked
		if strings.Index(ac, b) >= 0 {
			return count
		}

		ac = ac + a
		count++
	}

	// exactly 2 repeats should also be checked, but we avoided one excessive concatenation in the cycle
	if strings.Index(ac, b) >= 0 {
		return count
	}

	// b cannot be a repetition of a
	return -1
}

func test(a, b string, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Original string: %v \n", a)
	fmt.Printf("Potential string within repeats of the first string: %v \n", b)

	result := repeatedStringMatch(a, b)

	fmt.Printf("Counts of \"%s\" repeats that will contain \"%v\": %v \n", a, b, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("abcd", "cdabcdab", 3) // ab_cdabcdab_cd
}

func test2() {
	test("a", "aa", 2) // aa
}

func test3() {
	test("a", "bb", -1) // impossible
}

func test4() {
	test("aaaaaaaaaaaaaaaaaaaaaab", "ba", 2) // impossible
}

func main() {
	// 686. Repeated String Match
	test1()
	test2()
	test3()
	test4()
}
