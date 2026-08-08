package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func isPalindrome(head *ListNode) bool {
	// find the middle of the list
	var slow = head
	var fast = head.Next // ! we start fast from the 2nd node

	for (fast != nil) && (fast.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fmt.Printf("Middle node value: %v \n", slow.Val)

	middle := slow
	var secondHalfHead = middle.Next

	// disconnect 2nd path from the first path, to avoid cycles
	slow.Next = nil

	fmt.Println("Disconnected the 2nd half, 1st half is: ")
	listsCommon.PrintList(head)

	// reverse the 2nd path
	// middle itself must not be reversed? At least in case of odd count
	current := secondHalfHead
	var previous *ListNode = nil

	for current != nil {
		temp := current.Next
		current.Next = previous

		previous = current
		current = temp
	}

	secondHalfHead = previous

	fmt.Println("2nd half of the list, reversed:")
	listsCommon.PrintList(secondHalfHead)

	// compare 1st and 2nd half (reversed), up to minimal length (1st half may be +1 length for odd number of initial nodes)
	n1 := head
	n2 := secondHalfHead

	for n1 != nil && n2 != nil {
		if n1.Val != n2.Val {
			return false
		}

		n1 = n1.Next
		n2 = n2.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode { // returns the new head (was tail)
	fmt.Println("Original list:")
	listsCommon.PrintList(head)

	var previous *ListNode = nil
	current := head

	for current != nil {
		temp := current.Next
		current.Next = previous

		previous = current
		current = temp
	}

	head = previous

	fmt.Println("Reversed list:")
	listsCommon.PrintList(head)

	return head
}

func reverseListRecursive(node *ListNode) *ListNode {
	if (node == nil) || (node.Next == nil) {
		return node
	}

	// we're just returning the last node
	nextHead := reverseListRecursive(node.Next)

	node.Next.Next = node
	node.Next = nil

	return nextHead
}

func test(values []int, expectedResult bool) {
	fmt.Println()
	fmt.Println("===========================")

	list := listsCommon.ArrayToList(values)

	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("List: ")
	listsCommon.PrintList(list)

	result := isPalindrome(list)
	fmt.Printf("List of %v is a palindrome: %v \n", values, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{1, 2, 3, 2, 1}, true)
}

func test2() {
	test([]int{1, 2, 2, 1}, true)
}

func test3() {
	test([]int{1, 2, 300, 200, 100}, false)
}

func test4() {
	test([]int{1}, true)
}

func main() {
	// 234. Palindrome Linked List
	test1()
	test2()
	test3()
	test4()

	// test the reverseList separate function
	fmt.Println()
	fmt.Println("========================")

	values := []int{10, 20, 30, 40, 50}
	list := listsCommon.ArrayToList(values)
	//reversedList := reverseList(list)
	reversedList := reverseListRecursive(list)
	reverseList(reversedList)
}
