package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1}
	n := dummyHead

	n1 := list1
	n2 := list2

	// while both lists are non-empty, we have to compare the values
	for (n1 != nil) && (n2 != nil) {
		if n1.Val < n2.Val { // append from list 1
			// append from list1 to result
			n.Next = n1

			// go ahead in list1
			n1 = n1.Next
		} else { // append from list 2
			// append from list2 to result
			n.Next = n2

			// go ahead in list2
			n2 = n2.Next
		}

		// go ahead in the result list
		n = n.Next
	}

	// append rest of list1 - no need to iterate anymore!
	if n1 != nil {
		n.Next = n1
	}

	// append rest of list1 - no need to iterate anymore!
	if n2 != nil {
		n.Next = n2
	}

	/*	// append rest of list1
		for n1 != nil {
			// append from list1 to result
			n.Next = n1

			// go ahead in the result list
			n = n.Next

			// go ahead in list1
			n1 = n1.Next
		}

		// append rest of list2
		for n2 != nil {
			// append from list2 to result
			n.Next = n2

			// go ahead in the result list
			n = n.Next

			// go ahead in list2
			n2 = n2.Next
		}
	*/

	// skip the dummy head
	return dummyHead.Next
}

func test(a1 []int, a2 []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("===========================")

	l1 := listsCommon.ArrayToList(a1)
	l2 := listsCommon.ArrayToList(a2)

	fmt.Println("Sorted list 1:")
	listsCommon.PrintList(l1)

	fmt.Println("Sorted list 2:")
	listsCommon.PrintList(l2)

	result := mergeTwoLists(l1, l2) // todo: replace with your function
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Println("Merged sorted list:")
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

func test1() {
	a1 := []int{1, 2, 4}
	a2 := []int{1, 3, 5}
	expected := []int{1, 1, 2, 3, 4, 5}

	test(a1, a2, expected)
}

func test2() {
	a1 := []int{}
	a2 := []int{}
	expected := []int{}

	test(a1, a2, expected)
}

func test3() {
	a1 := []int{1, 2, 4}
	a2 := []int{1, 3, 4}
	expected := []int{1, 1, 2, 3, 4, 4}

	test(a1, a2, expected)
}

func main() {
	// 21. Merge Two Sorted Lists
	test1()
	test2()
	test3()
}
