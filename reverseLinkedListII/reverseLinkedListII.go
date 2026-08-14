package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
	"strconv"
)

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	if left >= right { // nothing to reverse
		return head
	}

	dummyHead := &ListNode{-666, head}

	beforeFirstToReverse := dummyHead

	pos := 1 // we're counting 1-based

	for pos != left { // we should be BEFORE left, so in case of 1 we will be before head
		beforeFirstToReverse = beforeFirstToReverse.Next
		pos++
	}

	// we're 1 node before position left, so proceed 1 step further to start on the node number left
	lastToReverse := beforeFirstToReverse.Next

	for pos != right {
		lastToReverse = lastToReverse.Next
		pos++
	}

	firstToReverse := beforeFirstToReverse.Next
	firstAfterReverse := lastToReverse.Next

	fmt.Printf("Before first to reverse: %v \n", valToString(beforeFirstToReverse))
	fmt.Printf("First to reverse: %v \n", valToString(firstToReverse))

	fmt.Printf("Last to reverse: %v \n", valToString(lastToReverse))
	fmt.Printf("First after reversed path: %v \n", valToString(firstAfterReverse))

	// disconnect the reversed part
	beforeFirstToReverse.Next = nil
	lastToReverse.Next = nil

	// reverse
	reversedHead := reverseList(firstToReverse)

	// connect the reversed path back
	beforeFirstToReverse.Next = reversedHead
	firstToReverse.Next = firstAfterReverse // first in the reversed path is now last in the reversed path

	// skip the dummy head
	return dummyHead.Next
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

func valToString(node *ListNode) string {
	// avoids NPE failing on printing node.Val in the log
	if node == nil {
		return "nil"
	}

	return strconv.Itoa(node.Val)
}

func test(arr []int, left, right int, expectedResult []int) { // linked list to linked list
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n")
	listsCommon.PrintList(list)

	result := reverseBetween(list, left, right)
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Printf("List with reversed elements [%v; %v]: \n", left, right)
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
	arr := []int{1, 2, 3, 4, 5}
	left := 2
	right := 4

	expected := []int{1, 4, 3, 2, 5}

	test(arr, left, right, expected)
}

func test2() {
	arr := []int{5}
	left := 1
	right := 1

	expected := []int{5}

	test(arr, left, right, expected)
}

func test3() {
	arr := []int{}
	left := 0
	right := 0

	expected := []int{}

	test(arr, left, right, expected)
}

func test4() {
	arr := []int{1, 2, 3, 4, 5} // reverse from the start
	left := 1
	right := 3

	expected := []int{3, 2, 1, 4, 5}

	test(arr, left, right, expected)
}

func test5() {
	arr := []int{1, 2, 3, 4, 5} // reverse from the start
	left := 1
	right := 2

	expected := []int{2, 1, 3, 4, 5}

	test(arr, left, right, expected)
}

func test6() {
	arr := []int{1, 2, 3, 4} // reverse at the end
	left := 3
	right := 4

	expected := []int{1, 2, 4, 3}

	test(arr, left, right, expected)
}

func test7() {
	arr := []int{1, 2, 3, 4} // reverse at the end
	left := 2
	right := 4

	expected := []int{1, 4, 3, 2}

	test(arr, left, right, expected)
}

func test8() {
	arr := []int{1, 2, 3, 4} // reverse the whole array
	left := 1
	right := 4

	expected := []int{4, 3, 2, 1}

	test(arr, left, right, expected)
}

func test9() {
	arr := []int{1, 2} // reverse the whole array
	left := 1
	right := 2

	expected := []int{2, 1}

	test(arr, left, right, expected)
}

func main() {
	// 92. Reverse Linked List II
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
	test9()
}
