package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if (head == nil) || (head.Next == nil) || (head.Next.Next == nil) {
		return false
	}

	// this is working
	//slow := head
	//fast := head.Next.Next

	// this is also working
	//slow := head
	//fast := head.Next

	// this is also working and is the fastest based on the test-set on Leetcode
	slow := head.Next
	fast := head.Next.Next

	for (slow != fast) && (fast.Next != nil) && (fast.Next.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if slow == fast { // two pointers met
		return true
	}

	return false
}

func test(head *ListNode, expected bool) {
	listHasCycle := hasCycle(head)

	fmt.Println()
	fmt.Println("============================")
	fmt.Printf("Expected has cycle: %v \n", expected)
	fmt.Printf("Detected has cycle: %v \n", listHasCycle)
}

func test1() {
	node3 := &ListNode{3, nil}
	node2 := &ListNode{2, nil}
	node0 := &ListNode{0, nil}
	nodeMinus4 := &ListNode{-4, nil}

	node3.Next = node2
	node2.Next = node0
	node0.Next = nodeMinus4
	nodeMinus4.Next = node2 // cycle!

	head := node3

	expected := true

	test(head, expected)
}

func test2() {
	node3 := &ListNode{3, nil}
	node2 := &ListNode{2, nil}
	node0 := &ListNode{0, nil}

	node3.Next = node2
	node2.Next = node0
	node0.Next = node2

	head := node3

	expected := true

	test(head, expected)
}

func test3() {
	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}

	node1.Next = node2
	node2.Next = node1

	head := node1

	expected := true

	test(head, expected)
}

func test4() {
	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	head := node1

	expected := false

	test(head, expected)
}

func main() {
	//test1()
	//test2()
	//test3()
	test4()
}
