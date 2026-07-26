package main

import (
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
)

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	if slow != fast { // no cycle
		return nil
	}

	slow2 := head

	for slow != slow2 {
		slow = slow.Next
		slow2 = slow2.Next
	}

	return slow
}

func main() {
	// 142. Linked List Cycle II
}
