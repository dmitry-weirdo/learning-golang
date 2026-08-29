package main

import "fmt"

func confusingNumber(n int) bool {
	reversed := 0

	current := n

	for current > 0 {
		digit := current % 10
		current = current / 10

		if !isValid(digit) { // any invalid digit makes the value invalid
			return false
		}

		reversed = reversed*10 + reverse(digit)
	}

	//fmt.Printf("Reversed number: %v \n", reversed)

	return reversed != n
}

func reverse(x int) int {
	switch x {
	case 0:
		return 0
	case 1:
		return 1
	case 6:
		return 9
	case 8:
		return 8
	case 9:
		return 6

	default:
		panic(fmt.Sprintf("Invalid value to reverse: %v.", x))
	}
}

func isValid(x int) bool {
	switch x {
	case 0, 1, 6, 8, 9:
		return true
	default:
		return false // invalid or out of range -> false
	}
}

func test(x int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x)

	result := confusingNumber(x)

	fmt.Printf("Number is confusing: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(6, true) // 6 -> 9
}

func test2() {
	test(89, true) // 89 -> 68
}

func test3() {
	test(11, false) // 11 -> 11, same number
}

func test4() {
	test(916, false) // 916 -> 916, same number
}

func test5() {
	test(979, false) // 7 does not revert valid
}

func main() {
	// 1056. Confusing Number
	test1()
	test2()
	test3()
	test4()
	test5()
}
