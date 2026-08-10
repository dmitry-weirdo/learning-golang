package main

import "fmt"

func magicalString(n int) int {
	// n >= 1
	if n <= 3 { // first 122 -> nothing to append
		return 1
	}

	// todo: we can use a queue with removal from the beginning
	// we can also use an array of boolean, but bool also take 1 byte, so this won't be a memory win
	a := make([]byte, n+1) // since we can append 2 values, add 1 additional position after N
	a[0] = 1
	a[1] = 2
	a[2] = 2

	definingPos := 2 // we start from the last 2 of 122

	writePos := 3 // to where we add the result

	currentDigit := byte(1) // it switches after each append, after 122 we're starting with 1

	totalOnes := 1 // first one we got from the starting "122"

	for writePos < n {
		// define so many currentDigit values as a[definingPos]
		for range a[definingPos] {
			a[writePos] = currentDigit
			writePos++
		}

		if currentDigit == 1 {
			totalOnes += int(a[definingPos])
		}

		// update to the next position
		definingPos++

		currentDigit = 3 - currentDigit // 1 -> 2, 2 -> 1
	}

	if (writePos > n) && (a[writePos-1] == 1) {
		// we appended 2 ones and 2nd of them is after N
		totalOnes--
	}

	//fmt.Printf("N: %v, writePos: %v \n", n, writePos)
	//fmt.Printf("Current sequence (first 10 values): \n%v\n", a[:10])

	return totalOnes
}

func test(n int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number of digits: %v \n", n)

	result := magicalString(n)

	fmt.Printf("Count of 1 in %v digits of the sequence: %v \n", n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(6, 3) // 122112
}

func test2() {
	test(1, 1) // 1
}

func test3() {
	test(4, 2) // 1221
}

func test4() {
	test(5, 3) // 12211
}

func test5() {
	test(6, 3) // 122112
}

func test6() {
	test(7, 4) // 1221121
}

func main() {
	// 481. Magical String
	// This is a Kolakoski sequence starting with "122", and the last 2 of these "122" is the starting defining digit.
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
