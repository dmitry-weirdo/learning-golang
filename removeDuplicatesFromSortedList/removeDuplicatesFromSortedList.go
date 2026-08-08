package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	n := head

	for n != nil {
		for n.Next != nil && n.Next.Val == n.Val { // current node same as next -> remove next
			n.Next = n.Next.Next
		}

		// current node different from next -> set it to the next node
		n = n.Next
	}

	return head // head node remains in place
}

func test(arr []int, expectedResult []int) {
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
	arr := []int{1, 1, 2, 3, 3}
	expected := []int{1, 2, 3}

	test(arr, expected)
}

func test2() {
	arr := []int{1, 1, 2}
	expected := []int{1, 2}

	test(arr, expected)
}

func test3() {
	arr := []int{}
	expected := []int{}

	test(arr, expected)
}

func test4() {
	arr := []int{1}
	expected := []int{1}

	test(arr, expected)
}

func test5() {
	arr := []int{1, 1}
	expected := []int{1}

	test(arr, expected)
}

func test6() {
	arr := []int{1, 2}
	expected := []int{1, 2}

	test(arr, expected)
}

func main() {
	// 83. Remove Duplicates from Sorted List
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
