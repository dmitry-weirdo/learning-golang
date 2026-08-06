package main

import "fmt"

func validMountainArray(arr []int) bool {
	return validMountainArray_optimized(arr) // logic can be simpler, less if-s -> smaller time
	//return validMountainArray_trivial(arr) // my straightforward solution
}

func validMountainArray_optimized(arr []int) bool {
	n := len(arr)

	i := 0

	// skip increasing part
	for (i < n-1) && (arr[i] < arr[i+1]) {
		i++
	}

	if i == 0 { // a[0] >= a[1] -> fail
		return false
	}

	if i >= n-1 { // increasing part ended on last element -> fail
		return false
	}

	// skip decreasing part after direction change
	for (i < n-1) && (arr[i] > arr[i+1]) {
		i++
	}

	// success if we reached the last element by decreasing sequence
	return i == n-1
}

func validMountainArray_trivial(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	if arr[0] >= arr[1] {
		return false
	}

	nextMustBeGreater := true

	i := 0

	for i < len(arr)-1 {
		if arr[i] == arr[i+1] { // equal elements -> fail
			return false
		}

		if nextMustBeGreater { // current state must be increasing
			if arr[i] < arr[i+1] { // increasing next element -> continue to the right
				i++
				continue
			}

			if arr[i] > arr[i+1] { // peak found -> switch from increasing to decreasing
				nextMustBeGreater = false
			}
		} else { // current state must be decreasing
			if arr[i] < arr[i+1] { // increasing while must be decreasing -> fail
				return false
			}

			if arr[i] > arr[i+1] { // decreasing next element -> continue to the right
				i++
				continue
			}
		}
	}

	if nextMustBeGreater { // reached pre-last still on increasing state -> fail
		return false
	}

	return true
}

func test(arr []int, expectedResult bool) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Heights array: %v \n", arr)

	result := validMountainArray(arr)

	fmt.Printf("Array is a valid mountain array: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
		return
	}
}

func test1() {
	arr := []int{1, 2}
	expected := false // less than 3 elements

	test(arr, expected)
}

func test2() {
	arr := []int{1, 3, 2}
	expected := true // minimum length of a valid array

	test(arr, expected)
}

func test3() {
	arr := []int{3, 5, 5, 2} // equal elements -> fail
	expected := false

	test(arr, expected)
}

func test4() {
	arr := []int{0, 3, 2, 1}
	expected := true

	test(arr, expected)
}

func main() {
	// 941. Valid Mountain Array
	test1()
	test2()
	test3()
	test4()
}
