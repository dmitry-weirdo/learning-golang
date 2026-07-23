package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	// todo: we can also solve this with a sliding window and count zeroes within
	// todo: a nice trick is just moving left by 1 and not while (zeroes > 1)
	//if len(nums) < 2 { // if array 0 or 1 element -> we can replace this single element
	//	return len(nums)
	//}

	// at the beginning of the array, onesBefore is 0, we calculate the current ones as onesAfter
	onesBefore := 0
	onesAfter := 0
	maxResult := 0
	zeroesEncountered := 0 // will switch to 1 if any zero found and should be counted in the end

	for _, v := range nums {
		if v == 1 {
			onesAfter++
		}

		if v == 0 {
			zeroesEncountered = 1
			currentSum := onesBefore + onesAfter + zeroesEncountered // + 1 is for middle 0 between ones, if there was a 0
			maxResult = max(maxResult, currentSum)

			// after switching 0, "after" of the previous 0 become "before" for the new 0
			onesBefore = onesAfter
			onesAfter = 0
		}
	}

	// after the end of the array, we need to compare the results
	currentSum := onesBefore + onesAfter + zeroesEncountered // + 1 is for middle 0 between ones, if there was a 0
	maxResult = max(maxResult, currentSum)

	return maxResult
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := findMaxConsecutiveOnes(arr)

	fmt.Printf("Max consecutive ones with one 0 replacement: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{1, 0, 1, 1, 0}
	expected := 4

	test(arr, expected)
}

func test2() {
	arr := []int{1, 0, 1, 1, 0, 1}
	expected := 4

	test(arr, expected)
}

func test3() {
	arr := []int{1}
	expected := 1

	test(arr, expected)
}

func test4() {
	arr := []int{0}
	expected := 1

	test(arr, expected)
}

func test5() {
	arr := []int{1, 0, 1, 1, 0, 1, 1, 1, 1, 1}
	expected := 8

	test(arr, expected)
}

func test6() {
	arr := []int{0, 0, 0}
	expected := 1

	test(arr, expected)
}

func test7() {
	arr := []int{1, 0, 1}
	expected := 3

	test(arr, expected)
}

func test8() {
	arr := []int{1, 1}
	expected := 2

	test(arr, expected)
}

func main() {
	// 487. Max Consecutive Ones II
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
}
