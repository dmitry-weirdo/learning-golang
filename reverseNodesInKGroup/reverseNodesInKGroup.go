package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	// use dummy node to have prev even for the old head
	dummyHead := &ListNode{-1, head}

	previous := dummyHead // before the current group
	var next *ListNode    // after the current group

	node := dummyHead

	endedBeforeKElements := false

	for !endedBeforeKElements {
		fmt.Println()
		fmt.Printf("Previous = %v, node = %v. Searching for the next %v elements... \n", previous.Val, node.Val, k)

		// count K elements
		for range k {
			if node.Next == nil {
				fmt.Printf("Node %v. No next node and not enough for %v nodes in the next group after the node %v. Do nothing for this remaining part of the list. \n", node.Val, k, previous.Val)
				endedBeforeKElements = true
				break
			} else {
				node = node.Next
			}
		}

		if !endedBeforeKElements {
			fmt.Printf("Found a group of %v elements. \n", k)
			fmt.Printf("First element before the group: %v. \n", previous.Val)
			fmt.Printf("Last element of the group: %v. \n", node.Val)

			next = node.Next

			if next != nil {
				fmt.Printf("Next element after the group: %v. \n", next.Val)
			} else {
				fmt.Printf("Next element after the group: %v. \n", nil)
			}

			// disconnect the last element of the group before reversing the group
			node.Next = nil

			firstElementInTheGroup := previous.Next
			reversedHead := reverse(firstElementInTheGroup)

			// connect previous to the head of the reversed list
			previous.Next = reversedHead

			// first element in the k-group is now the tail of the reversed list
			// connect the tail of the reversed list to the first element after the group
			firstElementInTheGroup.Next = next

			fmt.Printf("Reversed a group of %v elements: \n", k)
			listsCommon.PrintList(dummyHead)

			// jump to the next group
			previous = firstElementInTheGroup
			node = firstElementInTheGroup
		}
	}

	// skip the dummy head
	return dummyHead.Next
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head

	for current != nil {
		temp := current.Next
		current.Next = prev

		prev = current
		current = temp
	}

	return prev
}

func test(arr []int, k int, expectedResult []int) {
	fmt.Println("")
	fmt.Println("=============================")

	list := listsCommon.ArrayToList(arr)

	fmt.Println("Original list:")
	listsCommon.PrintList(list)

	fmt.Printf("K (k-group size): %v \n", k)

	result := reverseKGroup(list, k)
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Printf("Reversed by %v-groups list: \n", k)
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
	k := 2
	expected := []int{2, 1, 4, 3, 5}

	test(arr, k, expected)
}

func test2() {
	arr := []int{1, 2, 3, 4, 5}
	k := 3
	expected := []int{3, 2, 1, 4, 5}

	test(arr, k, expected)
}

func test3() {
	arr := []int{1, 2, 3, 4}
	k := 2
	expected := []int{2, 1, 4, 3}

	test(arr, k, expected)
}

func test4() {
	arr := []int{1}
	k := 2
	expected := []int{1}

	test(arr, k, expected)
}

func test5() {
	arr := []int{1}
	k := 1
	expected := []int{1}

	test(arr, k, expected)
}

func test6() {
	arr := []int{1, 2, 3}
	k := 3
	expected := []int{3, 2, 1}

	test(arr, k, expected)
}

func main() {
	// 25. Reverse Nodes in k-Group
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
