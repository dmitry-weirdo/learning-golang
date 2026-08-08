package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	d1 := 0
	d2 := 0
	addToNext := 0

	n1 := l1
	n2 := l2

	dummyHead := &ListNode{Val: -1}
	nextNode := dummyHead

	for (n1 != nil) || (n2 != nil) {
		if n1 == nil {
			d1 = 0
		} else {
			d1 = n1.Val
			n1 = n1.Next
		}

		if n2 == nil {
			d2 = 0
		} else {
			d2 = n2.Val
			n2 = n2.Next
		}

		currentSum := d1 + d2 + addToNext
		currentDigit := currentSum % 10

		nextNode.Next = &ListNode{Val: currentDigit}
		nextNode = nextNode.Next

		if currentSum > 9 { // 1 should be added to next digit
			addToNext = 1
		} else {
			addToNext = 0
		}

		fmt.Printf("currentDigit: %d \n", currentDigit)
	}

	// 1 could be added to the next digit that doesn't exist in any list
	if addToNext != 0 {
		fmt.Printf("+1 digit add to the next digit place: %d \n", addToNext)

		nextNode.Next = &ListNode{Val: addToNext}
		nextNode = nextNode.Next
	}

	// remove dummyHead from the result
	return dummyHead.Next
}

func test(a1 []int, a2 []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("===========================")

	l1 := listsCommon.ArrayToList(a1)
	l2 := listsCommon.ArrayToList(a2)

	fmt.Println("List 1:")
	listsCommon.PrintList(l1)

	fmt.Println("List 2:")
	listsCommon.PrintList(l2)

	result := addTwoNumbers(l1, l2)
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Println("Sum list:")
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
	a1 := []int{2, 4, 3}
	a2 := []int{5, 6, 4}
	expected := []int{7, 0, 8}

	test(a1, a2, expected)
}

func test2() {
	a1 := []int{9, 9, 9, 9, 9, 9, 9}
	a2 := []int{9, 9, 9, 9}
	expected := []int{8, 9, 9, 9, 0, 0, 0, 1}

	test(a1, a2, expected)
}

func test3() {
	a1 := []int{}
	a2 := []int{3, 2, 1}
	expected := []int{3, 2, 1}

	test(a1, a2, expected)
}

func main() {
	// 2. Add Two Numbers
	test1()
	test2()
	test3()
}
