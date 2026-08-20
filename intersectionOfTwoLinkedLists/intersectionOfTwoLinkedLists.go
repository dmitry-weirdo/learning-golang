package main

import (
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// todo: solution with a hashtable of list1 and then going through list2 will be a pure O(M + N) time, i.e. faster, but O(list1) space

	// Basic two-pointer solution.
	// -> O(M + N) time to find lengths
	// -> O(M - N) to traverse in the longer list forward.
	// -> O(2 * N) to traverse in both lists to the end (until we find the intersection).
	// (M + N) + (M - N) + (N + N) = 2M + 2N
	// O(2M + 2N) time = O(M + N)

	// Passes in around 45-50 ms
	return getIntersectionNode_twoPointers_clear(headA, headB)

	// iteration optimization that we can do everything in once cycle.
	// Shorter pointer goes (N + M - N) = M
	// Longer pointer goes M
	// Then they both go N
	// Total: 2M + 2N - same complexity as in the straightforward 2 pointers.

	// -> O(2M + 2N) time = O(M + N)
	//return getIntersectionNode_twoPointers_blackMagic(headA, headB)
}

func getIntersectionNode_twoPointers_clear(headA, headB *ListNode) *ListNode {
	lengthA := getListLength(headA)
	lengthB := getListLength(headB)

	shortHead := headA
	longHead := headB

	shorterLength, longerLength := getMinAndMax(lengthA, lengthB)
	lengthDiff := longerLength - shorterLength

	if lengthA > lengthB {
		shortHead = headB
		longHead = headA
	}

	// in the longer list, skip (M - N) nodes
	for range lengthDiff {
		longHead = longHead.Next
	}

	// now heads are the same N nodes from the end
	// iterate them both until they are same

	for shortHead != longHead {
		shortHead = shortHead.Next
		longHead = longHead.Next
	}

	return shortHead
}

func getListLength(head *ListNode) int {
	length := 0
	node := head

	for node != nil {
		length++
		node = node.Next
	}

	return length
}

func getMinAndMax(a, b int) (smaller, greater int) {
	if a <= b {
		return a, b
	}

	return b, a
}

func getIntersectionNode_twoPointers_blackMagic(headA, headB *ListNode) *ListNode {
	// optimization

	// M - length of longer list
	// N - length of shorter list

	// - when the shorter list ends, it moves to the head of the longer list, and continues the iteration.
	// - at this time, in longer list, we passed N nodes, and in the longer list, there are (M - N) values remaining.

	// - when the longer list ends, it moves to the head of the shorter list.
	// - It happens after in the longer list, we moved (M - N) values.
	// - At this time, another pointer in the longer list will be in position (M - N).

	// Therefore, we're in the same distance from the end in both lists:
	// - In the longer list, the poiner is at (M - N)
	// - In the shorter list, the pointer is at the beginning

	// todo: implement method
	return nil
}

func main() {
	// 160. Intersection of Two Linked Lists
	// todo: add tests (requires implementing 2 lists and then merge them)
}
