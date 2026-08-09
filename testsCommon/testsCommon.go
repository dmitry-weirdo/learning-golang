package testsCommon

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

// common test formats to copy-paste just from this file instead of searching by many files

// ==================== Linked List functions ==================== //
func ListToListFunction(head *ListNode) *ListNode {
	return head
}

func testListToList(arr []int, expectedResult []int) { // linked list to linked list
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n")
	listsCommon.PrintList(list)

	result := ListToListFunction(list) // todo: replace with your function
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Printf("List with removed duplicates: \n") // todo: replace with your text
	listsCommon.PrintList(result)

	fmt.Printf("Result as array: %v \n", resultAsArray)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(resultAsArray) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(resultAsArray))
		return
	}

	for i, v := range resultAsArray {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func TwoListsToListFunction(l1, l2 *ListNode) *ListNode {
	return l1
}

func testTwoListsToList(a1 []int, a2 []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("===========================")

	l1 := listsCommon.ArrayToList(a1)
	l2 := listsCommon.ArrayToList(a2)

	fmt.Println("List 1:")
	listsCommon.PrintList(l1)

	fmt.Println("List 2:")
	listsCommon.PrintList(l2)

	result := TwoListsToListFunction(l1, l2) // todo: replace with your function
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Println("Sum list:") // todo: replace with your text
	listsCommon.PrintList(result)

	fmt.Printf("Result as array: %v \n", resultAsArray)
	fmt.Printf("Expected result: %v \n", expectedResult)

	for i, v := range resultAsArray {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func ListToBooleanFunction(head *ListNode) bool {
	return false
}

func testListToBoolean(values []int, expectedResult bool) {
	fmt.Println()
	fmt.Println("===========================")

	list := listsCommon.ArrayToList(values)

	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("List: ")
	listsCommon.PrintList(list)

	result := ListToBooleanFunction(list) // todo: replace with your function

	fmt.Printf("List of %v is a palindrome: %v \n", values, result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func ListToIntArrayFunction(head *ListNode) []int {
	return nil
}

func testListToIntArray(arr []int, expectedResult []int) { // linked list to linked list
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n") // todo: replace with your text if required
	listsCommon.PrintList(list)

	result := ListToIntArrayFunction(list) // todo: replace with your function

	fmt.Printf("Result: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

// ==================== []int int arrays functions ==================== //
func IntArrayToIntArrayFunction(arr []int) []int {
	return arr
}

func testIntArrayToIntArray(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr) // todo: replace with your text

	result := IntArrayToIntArrayFunction(arr) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func IntArrayToIntFunction(arr []int) int {
	return len(arr)
}

func testIntArrayToInt(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr) // todo: replace with your text if required

	result := IntArrayToIntFunction(arr)

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func IntArrayToPairsArrayFunction(arr []int) [][]int {
	return nil
}

func testIntArrayToPairsArrayFunction(arr []int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array: %v \n", arr) // todo: replace with your text if required

	result := IntArrayToPairsArrayFunction(arr)

	fmt.Printf("Pairs array:   %v \n", result) // todo: replace with your text if required
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i := 0; i < len(expectedResult); i++ {
		r := result[i]
		er := expectedResult[i]

		if r[0] != er[0] || r[1] != er[1] { // we only compare pairs, i.e. [0] and [1] elements
			fmt.Printf("FAILURE: expected result[%v] = [%v; %v], actual result[%v] = [%v; %v] \n", i, er[0], er[1], i, r[0], r[1])
		}
	}
}
