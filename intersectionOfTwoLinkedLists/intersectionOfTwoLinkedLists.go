package main

import (
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// Solution with a hashtable of list1 and then going through list2
	// should be a pure O(M + N) time, i.e. faster, but O(list1) space
	// O(M + N) time, O(M) space
	// !!! Interestingly, it takes bigger time 50-65 ms, since hashMap is a heavier/slower structure than just pointers.

	// If we switch order of A and B iteration, to 45-55 ms :) But this is just "I am lucky with the test-cases" optimization.
	return getIntersectionNode_hashMap(headA, headB)

	// Basic two-pointer solution.
	// -> O(M + N) time to find lengths
	// -> O(M - N) to traverse in the longer list forward.
	// -> O(2 * N) to traverse in both lists to the end (until we find the intersection).
	// (M + N) + (M - N) + (N + N) = 2M + 2N
	// O(2M + 2N) time = O(M + N)

	// Passes in around 45-50 ms
	//return getIntersectionNode_twoPointers_clear(headA, headB)

	// iteration optimization that we can do everything in once cycle.
	// Shorter pointer goes (N + M - N) = M
	// Longer pointer goes M
	// Then they both go N
	// Total: 2M + 2N - same complexity as in the straightforward 2 pointers.

	// -> O(2M + 2N) time = O(M + N)
	// Passes in the same 45-50 ms
	// -> there is NO NEED in this optimization, it complexifies the code, but makes NO time win.
	//return getIntersectionNode_twoPointers_blackMagic(headA, headB)
}

func getIntersectionNode_hashMap(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)

	node := headA

	for node != nil {
		m[node] = true
		node = node.Next
	}

	node = headB

	for node != nil {
		if m[node] {
			return node
		}

		node = node.Next
	}

	return node // will be nil
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

	pA := headA
	pB := headB

	for pA != pB { // after 2 jumps to the lists start, we'll continue comparing from the same point (N nodes from the end of both lists)
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}

		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}

	return pA
}

func main() {
	// 160. Intersection of Two Linked Lists
	// todo: add tests (requires implementing 2 lists and then merge them)
}
