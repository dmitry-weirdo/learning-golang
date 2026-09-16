package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{-666, head}

	before := dummyHead

	start := head
	end := start

	count := 0 // count of nodes with the same value

	for start != nil {
		count = 0
		end = start

		for (end != nil) && (end.Val == start.Val) {
			count++
			end = end.Next
		}

		if count > 1 {
			// remove nodes from start to end
			before.Next = end // before remains the same
			start = end       // start jumps to the next value
		} else { // just 1 count
			before = start
			start = end
		}
	}

	return dummyHead.Next
}

func test(arr []int, expectedResult []int) { // linked list to linked list
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n")
	listsCommon.PrintList(list)

	result := deleteDuplicates(list)
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Printf("List with removed duplicates: \n")
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

func test1() {
	test(
		[]int{1, 2, 3, 3, 4, 4, 5},
		[]int{1, 2, 5},
	)
}

func test2() {
	test(
		[]int{1, 1, 1, 2, 3},
		[]int{2, 3},
	)
}

func test3() {
	test(
		[]int{1, 1},
		[]int{},
	)
}

func test4() {
	test(
		[]int{1},
		[]int{1}, // no duplicates
	)
}

func main() {
	// 82. Remove Duplicates from Sorted List II
	// More complex variation of "83. Remove Duplicates from Sorted List". Here we need to delete all copies of the duplicates, NOT leave a single copy in place.
	test1()
	test2()
	test3()
	test4()
}
