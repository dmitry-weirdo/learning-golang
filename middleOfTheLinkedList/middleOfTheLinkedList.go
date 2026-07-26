package main

import (
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
)

func middleNode(head *ListNode) *ListNode {
	//dummyHead := &ListNode{
	//	Val:  -666,
	//	Next: head,
	//}

	// in case of even list, it will return the 2nd of the middle
	// 1 - 2 - 3 - 4 -> slow will be at 4
	slow := head
	fast := head

	// if we want the first of the middle
	// 1 - 2 - 3 - 4 -> slow will be at 2
	//fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func main() {
	// 876. Middle of the Linked List
}
