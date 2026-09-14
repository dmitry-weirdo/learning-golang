package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func convertToBase7(num int) string {
	return convertToBase7_builtIn(num)

	//return convertToBase7_manual(num)
}

func convertToBase7_builtIn(num int) string {
	return strconv.FormatInt(int64(num), 7)
}

func convertToBase7_manual(num int) string {
	if num == 0 {
		return "0"
	}

	var sb strings.Builder

	// to not have any problem with mod operation on negative numbers, work on the absolute values.
	n := abs(num)

	for n > 0 {
		lastBit := n % 7
		n = n / 7
		sb.WriteString(strconv.Itoa(lastBit))
	}

	// append "-" sign at the end
	if num < 0 {
		sb.WriteByte('-')
	}

	// we collected from smallest bit up, so we need to reverse the string
	s := sb.String()
	return reverseString(s)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func reverseString(s string) string {
	stringAsSlice := []byte(s)
	slices.Reverse(stringAsSlice) // reverses in place

	return string(stringAsSlice)
}

func test(x int, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x)

	result := convertToBase7(x)

	fmt.Printf("Number %v in base-7: %v \n", x, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(0, "0")
}

func test2() {
	test(100, "202") // 2 * 49 + 0 * 7 + 2 * 1
}

func test3() {
	test(-7, "-10") // 1 * 7 + 0 * 1
}

func main() {
	// 504. Base 7
	test1()
	test2()
	test3()
}
