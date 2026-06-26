package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func printList(head *ListNode) {
	var node = head

	for node != nil {
		fmt.Printf("%d ", node.Val)

		node = node.Next
	}

	fmt.Println()
}

func arrayToList(arr []int) *ListNode {
	var node *ListNode = nil

	for i := len(arr) - 1; i >= 0; i-- {
		nextNode := &ListNode{arr[i], node}

		node = nextNode
	}
	return node
}

func test(a1 []int, a2 []int) {
	fmt.Println()
	fmt.Println("===========================")

	l1 := arrayToList(a1)
	l2 := arrayToList(a2)

	fmt.Println("List 1:")
	printList(l1)

	fmt.Println("List 2:")
	printList(l2)

	sumList := addTwoNumbers(l1, l2)

	fmt.Println("Sum list:")
	printList(sumList)
}

func test1() {
	a1 := []int{2, 4, 3}
	a2 := []int{5, 6, 4}

	test(a1, a2)
}

func test2() {
	a1 := []int{9, 9, 9, 9, 9, 9, 9}
	a2 := []int{9, 9, 9, 9}

	test(a1, a2)
}

func test3() {
	a1 := []int{}
	a2 := []int{3, 2, 1}

	test(a1, a2)
}

func main() {
	test1()
	test2()
	test3()
}
