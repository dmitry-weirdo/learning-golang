package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)

	condition := func(x byte) bool {
		return x > target
	}

	index := binarySearchGeneric(letters, 0, n, condition)

	if index == n { // if not found -> return the first character in array. Notably, -1 will be an overflow for byte in Go
		return letters[0]
	}

	return letters[index]
}

func binarySearchGeneric(
	arr []byte, // todo: we can also generalize the type in the array
	left int, // usually it starts with 0, if we search in the complete array
	right int, // set len(arr) - 1 if you want to be within array. Set len(arr) if index after the array can be returned.
	condition func(byte) bool, // we will find the leftmost index satisfying this condition within [left; right] range
) int {
	// todo: this method can return an incorrect value for the empty array

	// Using a template from:
	// https://leetcode.com/discuss/post/786126/python-powerful-ultimate-binary-search-t-rwv8/
	for left < right {
		mid := (left + right) / 2

		if condition(arr[mid]) { // target condition
			right = mid // in this template it is always mid, NOT mid - 1
		} else {
			left = mid + 1
		}
	}

	// after exiting the while loop, left is the minimal k satisfying the condition function;
	return left
}

func test(arr []byte, target byte, expectedResult byte) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %s \n", arr)
	fmt.Printf("Target: %c \n", target)

	result := nextGreatestLetter(arr, target)

	fmt.Printf("First letter greater than %c: %c \n", target, result)
	fmt.Printf("Expected result: %c \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %c, actual result = %c \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]byte{'c', 'f', 'j'},
		'a',
		'c',
	)
}

func test2() {
	test(
		[]byte{'c', 'f', 'j'},
		'c',
		'f',
	)
}

func test3() {
	test(
		[]byte{'x', 'x', 'y', 'y'},
		'z',
		'x', // not found -> we return the first element in the array
	)
}

func main() {
	// 744. Find Smallest Letter Greater Than Target
	test1()
	test2()
	test3()
}
