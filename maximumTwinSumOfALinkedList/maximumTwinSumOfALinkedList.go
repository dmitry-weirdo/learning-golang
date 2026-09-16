package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func pairSum(head *ListNode) int {

	// we're guaranteed to have at least 2 nodes in the list
	middle := getListMiddle(head)

	secondHalfHead := middle.Next

	//fmt.Printf("Middle: %v \n", middle.Val)
	//fmt.Printf("Second half start: %v \n", secondHalfHead.Val)

	// disconnect the 2nd half
	middle.Next = nil

	// reverse the 2nd half
	secondHalfHead = reverseList(secondHalfHead)

	//fmt.Printf("Reversed 2nd half: \n")
	//listsCommon.PrintList

	// count twins
	maxTwinSum := head.Val + secondHalfHead.Val

	n1 := head
	n2 := secondHalfHead

	for n1 != nil && n2 != nil {
		maxTwinSum = max(maxTwinSum, n1.Val+n2.Val)

		n1 = n1.Next
		n2 = n2.Next
	}

	return maxTwinSum
}

func getListMiddle(head *ListNode) *ListNode {
	// even elements list: 1 - 2 - 3 - 4 -> return 2
	// odd elements list:  1 - 2 - 3 - 4 - 5 -> return 3

	// corner-cases:
	// 1 - 2 -> return 1
	// 1 -> return 1
	// nil -> return nil

	if head == nil {
		return nil
	}

	slow := head
	fast := head.Next // to make it work for both even add odd nodes count

	for (fast != nil) && (fast.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode { // returns the new head (was tail)
	//fmt.Println("Original list:")
	//listsCommon.PrintList(head)

	var previous *ListNode = nil
	current := head

	for current != nil {
		temp := current.Next
		current.Next = previous

		previous = current
		current = temp
	}

	head = previous

	//fmt.Println("Reversed list:")
	//listsCommon.PrintList(head)

	return head
}

func test(values []int, expectedResult int) {
	fmt.Println()
	fmt.Println("===========================")

	list := listsCommon.ArrayToList(values)

	fmt.Printf("List: ")
	listsCommon.PrintList(list)

	result := pairSum(list)

	fmt.Printf("Max twin sum in a list: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]int{5, 4, 2, 1},
		6, // 5 + 1 = 4 + 2
	)
}

func test2() {
	test(
		[]int{4, 2, 2, 3},
		7, // 4 + 3 = 7
	)
}

func test3() {
	test(
		[]int{1, 100_000},
		100_001, // 100_000 + 1 = 100_001
	)
}

func main() {
	// 2130. Maximum Twin Sum of a Linked List
	test1()
	test2()
	test3()
}
