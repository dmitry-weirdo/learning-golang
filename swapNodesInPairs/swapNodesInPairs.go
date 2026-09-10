package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 0 or 1 nodes in the list -> nothing to handle
		return head
	}

	dummyHead := &ListNode{Val: -666, Next: head}

	before := dummyHead
	first := nextOrNil(before)
	second := nextOrNil(first)
	after := nextOrNil(second)

	//for (before != nil) && (first != nil) && (second != nil) { // while there are nodes to swap
	for second != nil { // checking just the last one is enough, it will be null if previous nodes are null
		before.Next = second
		second.Next = first
		first.Next = after

		// change the nodes
		before = first // jump 2 ahead
		first = nextOrNil(before)
		second = nextOrNil(first)
		after = nextOrNil(second)
	}

	return dummyHead.Next
}

func nextOrNil(node *ListNode) *ListNode {
	// avoids NPE failing on assigning
	// next = current.Next
	if node == nil {
		return nil
	}

	return node.Next
}

func test(arr []int, expectedResult []int) { // linked list to linked list
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n")
	listsCommon.PrintList(list)

	result := swapPairs(list)
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Printf("List with pair-swapped nodes: \n")
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
		[]int{1, 2, 3, 4},
		[]int{2, 1, 4, 3},
	)
}

func test2() {
	test(
		[]int{},
		[]int{},
	)
}

func test3() {
	test(
		[]int{1},
		[]int{1},
	)
}

func test4() {
	test(
		[]int{1, 2, 3},
		[]int{2, 1, 3},
	)
}

func test5() {
	test(
		[]int{1, 2},
		[]int{2, 1},
	)
}

func main() {
	// 24. Swap Nodes in Pairs
	test1()
	test2()
	test3()
	test4()
	test5()
}
