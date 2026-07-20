package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

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
		for i := 0; i < k; i++ {
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
			printList(dummyHead)

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

func printList(head *ListNode) {
	n := head

	for n != nil {
		fmt.Printf("%v ", n.Val)

		n = n.Next
	}

	fmt.Println()
}

func generateLinkedList(values []int) *ListNode {
	head := &ListNode{Val: values[0]}

	n := head

	for i := 1; i < len(values); i++ {
		newNode := &ListNode{Val: values[i]}

		n.Next = newNode
		n = n.Next
	}

	return head
}

func test(arr []int, k int) {
	fmt.Println("")
	fmt.Println("=============================")

	list := generateLinkedList(arr)

	fmt.Println("Original list:")
	printList(list)

	reversedList := reverseKGroup(list, k)

	fmt.Printf("Reversed by %v-groups list: \n", k)
	printList(reversedList)
}

func test1() {
	arr := []int{1, 2, 3, 4, 5}
	k := 2

	test(arr, k)
}

func test2() {
	arr := []int{1, 2, 3, 4, 5}
	k := 3

	test(arr, k)
}

func test3() {
	arr := []int{1, 2, 3, 4}
	k := 2

	test(arr, k)
}

func test4() {
	arr := []int{1}
	k := 2

	test(arr, k)
}

func test5() {
	arr := []int{1}
	k := 1

	test(arr, k)
}

func test6() {
	arr := []int{1, 2, 3}
	k := 3

	test(arr, k)
}

func main() {
	// 25. Reverse Nodes in k-Group
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	test6()
}
