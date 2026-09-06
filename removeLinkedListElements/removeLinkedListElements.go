package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{-666, head}

	current := dummyHead
	next := current.Next

	for next != nil {
		if next.Val == val { // remove next
			current.Next = next.Next
			next = current.Next
			continue
		}

		// proceed to next node
		current = next
		next = current.Next
	}

	return dummyHead.Next
}

func test(arr []int, val int, expectedResult []int) { // linked list to linked list
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n")
	listsCommon.PrintList(list)

	result := removeElements(list, val)
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Printf("List with removed %v values: \n", val)
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
		[]int{1, 2, 6, 3, 4, 5, 6},
		6,
		[]int{1, 2, 3, 4, 5},
	)
}

func test2() {
	test(
		[]int{},
		1,
		[]int{},
	)
}

func test3() {
	test(
		[]int{7, 7, 7, 7},
		7,
		[]int{},
	)
}

func main() {
	// 203. Remove Linked List Elements
	test1()
	test2()
	test3()
}
