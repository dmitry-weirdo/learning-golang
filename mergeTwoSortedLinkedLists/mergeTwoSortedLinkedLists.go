package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1}
	n := dummyHead

	n1 := list1
	n2 := list2

	// while both lists are non-empty, we have to compare the values
	for (n1 != nil) && (n2 != nil) {
		if n1.Val < n2.Val { // append from list 1
			// append from list1 to result
			n.Next = n1

			// go ahead in list1
			n1 = n1.Next
		} else { // append from list 2
			// append from list2 to result
			n.Next = n2

			// go ahead in list2
			n2 = n2.Next
		}

		// go ahead in the result list
		n = n.Next
	}

	// append rest of list1 - no need to iterate anymore!
	if n1 != nil {
		n.Next = n1
	}

	// append rest of list1 - no need to iterate anymore!
	if n2 != nil {
		n.Next = n2
	}

	/*	// append rest of list1
		for n1 != nil {
			// append from list1 to result
			n.Next = n1

			// go ahead in the result list
			n = n.Next

			// go ahead in list1
			n1 = n1.Next
		}

		// append rest of list2
		for n2 != nil {
			// append from list2 to result
			n.Next = n2

			// go ahead in the result list
			n = n.Next

			// go ahead in list2
			n2 = n2.Next
		}
	*/

	// skip the dummy head
	return dummyHead.Next
}

func arrayToList(arr []int) *ListNode {
	var node *ListNode = nil

	for i := len(arr) - 1; i >= 0; i-- {
		nextNode := &ListNode{arr[i], node}

		node = nextNode
	}

	return node
}

func printList(head *ListNode) {
	var node = head

	for node != nil {
		fmt.Printf("%d ", node.Val)

		node = node.Next
	}

	fmt.Println()
}

func test(a1, a2 []int) {
	fmt.Println()
	fmt.Println("=======================")

	l1 := arrayToList(a1)
	l2 := arrayToList(a2)

	fmt.Println("Sorted list 1:")
	printList(l1)

	fmt.Println("Sorted list 2:")
	printList(l2)

	mergedList := mergeTwoLists(l1, l2)

	fmt.Println("Merged list:")
	printList(mergedList)
}

func test1() {
	a1 := []int{1, 2, 4}
	a2 := []int{1, 3, 5}

	test(a1, a2)
}

func test2() {
	a1 := []int{}
	a2 := []int{}

	test(a1, a2)
}

func main() {
	test1()
	test2()
}
