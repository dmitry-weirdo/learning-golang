package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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
	previousOfNode := head

	for node != nil {
		fmt.Println()
		fmt.Println("============================")
		fmt.Printf("Handling next node %v \n ", node.Val)
		printList(start)

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
					fmt.Printf("Inserting node %v as the new head.", node.Val)

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
		}

		// jump to the next element that was after the node
		node = next
	}

	return start
}

func printList(head *ListNode) {
	var node = head

	for node != nil {
		fmt.Printf("%d ", node.Val)

		node = node.Next
	}

	fmt.Println()
}

func test(arr []int) {
	var node *ListNode = nil

	for i := len(arr) - 1; i >= 0; i-- {
		nextNode := &ListNode{arr[i], node}

		node = nextNode
	}

	// node will be the head
	fmt.Println("Unsorted list:")
	printList(node)

	sortedHead := insertionSortList(node)

	fmt.Println("Sorted list:")
	printList(sortedHead)
}

func test1() {
	arr := []int{4, 2, 1, 3}

	test(arr)
}

func test2() {
	arr := []int{-1, 5, 3, 4, 0}

	test(arr)
}

func main() {
	//test1()
	test2()
}
