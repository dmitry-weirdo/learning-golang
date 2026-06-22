package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// find the middle of the list
	var slow = head
	var fast = head.Next // ! we start fast from the 2nd node

	for (fast != nil) && (fast.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fmt.Printf("Middle node value: %v \n", slow.Val)

	middle := slow
	var secondHalfHead = middle.Next

	// todo: disconnect 2nd path from the first path, to avoid cycles
	slow.Next = nil

	fmt.Println("Disconnected the 2nd half, 1st half is: ")
	printList(head)

	// reverse the 2nd path
	// middle itself must not be reversed? At least in case of odd count
	current := secondHalfHead
	var previous *ListNode = nil

	for current != nil {
		fmt.Printf("Current: %v \n", current.Val)

		temp := current.Next
		current.Next = previous

		previous = current
		current = temp
	}

	secondHalfHead = previous

	fmt.Println("2nd half of the list, reversed:")
	printList(secondHalfHead)

	// compare 1st and 2nd half (reversed), up to minimal length (1st half may be +1 length for odd number of initial nodes)
	n1 := head
	n2 := secondHalfHead

	for n1 != nil && n2 != nil {
		if n1.Val != n2.Val {
			return false
		}

		n1 = n1.Next
		n2 = n2.Next
	}

	return true
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

func runTest(values []int) {
	list := generateLinkedList(values)

	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("List: ")
	printList(list)

	palindrome := isPalindrome(list)
	fmt.Printf("List of %v is a palindrome: %v \n", values, palindrome)
}

func test1() {
	runTest([]int{1, 2, 3, 2, 1})
}

func test2() {
	runTest([]int{1, 2, 2, 1})
}

func test3() {
	runTest([]int{1, 2, 300, 200, 100})
}

func test4() {
	runTest([]int{1})
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
