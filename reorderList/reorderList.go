package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use ListNode without prefix
	"fmt"
	"strconv"
)

func reorderList(head *ListNode) {
	if head == nil { // empty list -> nothing to do
		return
	}

	// find the middle of the list
	mid := getListMiddle(head)
	fmt.Printf("Middle of the list: %v \n", valToString(mid))

	// reverse the path after the middle
	reversedAfterMid := reverseList(mid.Next)
	mid.Next = nil // disconnect 2 lists, to not have the problematic (mid -> 2nd half) potentially cyclic link

	fmt.Println("List before mid (disconnected from the 2nd half):")
	listsCommon.PrintList(head)

	fmt.Println("Reversed list after mid:")
	listsCommon.PrintList(reversedAfterMid)

	// now we finally do 2 pointers insertion
	writePos := head                    // insert after this node
	writePosNext := nextOrNil(writePos) // insert before this node

	rightPos := reversedAfterMid        // what we copy between writePos and writePosNext
	rightPosNext := nextOrNil(rightPos) // to where we just after moving the rightPos node

	for rightPos != nil {
		// move rightPos after writePos
		writePos.Next = rightPos
		rightPos.Next = writePosNext

		// jump the write position
		writePos = writePosNext
		writePosNext = nextOrNil(writePos)

		// jump the right (copy from) position
		rightPos = rightPosNext
		rightPosNext = nextOrNil(rightPos)
	}
}

func getListMiddle(head *ListNode) *ListNode {
	// even elements list: 1 - 2 - 3 - 4 -> return 2
	// odd elements list:  1 - 2 - 3 - 4 - 5 -> return 3

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

func nextOrNil(node *ListNode) *ListNode {
	// avoids NPE failing on assigning
	// next = current.Next
	if node == nil {
		return nil
	}

	return node.Next
}

func valToString(node *ListNode) string {
	// avoids NPE failing on assigning
	// next = current.Next
	if node == nil {
		return "nil"
	}

	return strconv.Itoa(node.Val)
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

func test(arr []int, expectedResult []int) { // linked list to linked list
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n")
	listsCommon.PrintList(list)

	reorderList(list) // returns nothing, we must use the same head
	//	resultAsArray := listsCommon.ListToArray(list)

	fmt.Printf("Reordered list: \n")
	listsCommon.PrintList(list)

	resultAsArray := listsCommon.ListToArray(list) // todo: remove, uncomment above

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
	arr := []int{1, 2, 3, 4}
	expected := []int{1, 4, 2, 3}

	test(arr, expected)
}

func test2() {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{1, 5, 2, 4, 3}

	test(arr, expected)
}

func test3() {
	arr := []int{1}
	expected := []int{1}

	test(arr, expected)
}

func test4() {
	arr := []int{1, 2}
	expected := []int{1, 2}

	test(arr, expected)
}

func test5() {
	arr := []int{1, 2, 3}
	expected := []int{1, 3, 2}

	test(arr, expected)
}

func test6() {
	arr := []int{}
	expected := []int{}

	test(arr, expected)
}

func main() {
	// 143. Reorder List
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
