package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
	"strconv"
)

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil { // just 1 element -> remove it, list will be empty
		return nil
	}

	middle, beforeMiddle, _, totalElements := getListMiddleAndBeforeMiddle(head)

	// in this problem, for even number of the elements in the list, we count middle AFTER the first half
	if totalElements%2 == 0 {
		beforeMiddle = beforeMiddle.Next
		middle = middle.Next
	}

	// remove the middle element
	beforeMiddle.Next = middle.Next

	return head
}

func getListMiddleAndBeforeMiddle(head *ListNode) (middle, beforeMiddle *ListNode, middleIndex, totalElements int) {
	// even elements list: 1 - 2 - 3 - 4 -> return 2, 1, 1, 4
	// odd elements list:  1 - 2 - 3 - 4 - 5 -> return 3, 2, 2, 5

	// middleIndex starts with 0

	// corner-cases:
	// 1 - 2 -> return 1
	// 1 -> return 1
	// nil -> return nil

	if head == nil {
		return nil, nil, -1, 0
	}

	dummyHead := &ListNode{Val: -666, Next: head}

	slowIndex := 0
	totalElements = 0

	beforeSlow := dummyHead
	slow := head
	fast := head.Next // to make it work for both even add odd nodes count

	for (fast != nil) && (fast.Next != nil) {
		slowIndex++
		totalElements += 2

		beforeSlow = beforeSlow.Next
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil {
		totalElements++
	} else if fast.Next == nil {
		totalElements += 2
	}

	//fmt.Printf("Slow index: %v \n", slowIndex)
	//fmt.Printf("Total elements: %v \n", totalElements)
	//fmt.Printf("Fast pointer: %v \n", valToString(fast))

	return slow, beforeSlow, slowIndex, totalElements // yes, we will return a dummyHead, so we can remove the head as well
}

func valToString(node *ListNode) string {
	// avoids NPE failing on printing node.Val in the log
	if node == nil {
		return "nil"
	}

	return strconv.Itoa(node.Val)
}

func test(arr []int, expectedResult []int) { // linked list to linked list
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n")
	listsCommon.PrintList(list)

	result := deleteMiddle(list)
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Printf("List with removed middle element: \n")
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
		[]int{1, 3, 4, 7, 1, 2, 6},
		[]int{1, 3, 4, 1, 2, 6},
	)
}

func test2() {
	test(
		[]int{1, 2, 3, 4},
		[]int{1, 2, 4},
	)
}

func test3() {
	test(
		[]int{2, 1},
		[]int{2},
	)
}

func test4() {
	test(
		[]int{1},
		[]int{},
	)
}

func main() {
	// 2095. Delete the Middle Node of a Linked List
	test1()
	test2()
	test3()
	test4()
}
