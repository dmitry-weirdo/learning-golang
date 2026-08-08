package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil { // avoid case of 1-element list
		return head
	}

	// cache the head of the current list
	start := head // can be updated if we insert into the super-beginning of the list

	node := head.Next
	previousOfNode := head // will be NOT null

	for node != nil {
		fmt.Println()
		fmt.Println("============================")
		fmt.Printf("Handling next node %v \n", node.Val)
		fmt.Printf("Previous of node: %v \n", previousOfNode.Val)
		listsCommon.PrintList(start)

		// up to what node we go on insertion
		next := node.Next

		currentValue := node.Val

		var previousNode *ListNode = nil
		insertionNode := start

		sameNode := false

		for insertionNode != next {
			fmt.Println("")
			fmt.Println("============================")
			if previousNode == nil {
				fmt.Printf("Previous node is null \n")
			} else {
				fmt.Printf("Previous node: %v \n", previousNode.Val)
			}

			fmt.Printf("Insertion node: %v \n", insertionNode.Val)

			if insertionNode.Val >= currentValue {
				fmt.Printf("Inserting node %v before value %v \n", currentValue, insertionNode.Val)

				// todo: maybe a too ugly hack
				if node == insertionNode {
					fmt.Printf("Node %v and insertion node %v are the same node. Do nothing. \n", node.Val, insertionNode.Val)
					sameNode = true

					break
				}

				// inserting node between previousNode and insertionNode
				if previousNode == nil { // inserting before head
					fmt.Printf("Inserting node %v as the new head. \n", node.Val)

					start = node
					node.Next = insertionNode
				} else {
					fmt.Printf("Inserting node %v between values %v and %v. \n", currentValue, previousNode.Val, insertionNode.Val)

					previousNode.Next = node
					node.Next = insertionNode
				}

				break
			}

			previousNode = insertionNode
			insertionNode = insertionNode.Next
		}

		if !sameNode {
			// unextract node from its old position
			previousOfNode.Next = next
		} else {
			// if it's same node -> previous of the next node remains the current node
			previousOfNode = node
			fmt.Printf("Previous of node changed to %v", node.Val)
		}

		// jump to the next element that was after the node
		node = next
	}

	return start
}

func test(arr []int, expectedResult []int) { // linked list to linked list
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Unsorted list from array: \n")
	listsCommon.PrintList(list)

	result := insertionSortList(list)
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Printf("Sorted list: \n")
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
	arr := []int{4, 2, 1, 3}
	expected := []int{1, 2, 3, 4}

	test(arr, expected)
}

func test2() {
	arr := []int{-1, 5, 3, 4, 0}
	expected := []int{-1, 0, 3, 4, 5}

	test(arr, expected)
}

func test3() {
	//arr := []int{4, 3, 2, 1}
	arr := []int{4, 2, 2, 1}
	expected := []int{1, 2, 2, 4}

	test(arr, expected)
}

func main() {
	// 147. Insertion Sort List
	test1()
	test2()
	test3()
}
