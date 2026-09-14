package main

import "fmt"

func titleToNumber(columnTitle string) int {
	x := 0
	base := 1

	for i := len(columnTitle) - 1; i >= 0; i-- {
		v := int(columnTitle[i] - 'A' + 1)
		x += base * v

		base *= 26
	}

	return x
}

func test(s string, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Excel column title: %v \n", s)

	result := titleToNumber(s)

	fmt.Printf("Column \"%v\" as 1-based number: %v \n", s, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("A", 1)
}

func test2() {
	test("Z", 26)
}

func test3() {
	test("AB", 28)
}

func test4() {
	test("ZY", 701)
}

func main() {
	// 171. Excel Sheet Column Number
	test1()
	test2()
	test3()
	test4()
}
