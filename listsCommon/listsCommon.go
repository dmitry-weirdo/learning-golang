package listsCommon

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(head *ListNode) {
	n := head

	for n != nil {
		fmt.Printf("%v ", n.Val)

		n = n.Next
	}

	fmt.Println()
}

func ArrayToList(arr []int) *ListNode {
	var node *ListNode = nil

	for i := len(arr) - 1; i >= 0; i-- {
		nextNode := &ListNode{arr[i], node}

		node = nextNode
	}

	return node
}

func ArraysToLists(arrays ...[]int) []*ListNode {
	lists := make([]*ListNode, len(arrays))

	for i, arr := range arrays {
		list := ArrayToList(arr)
		lists[i] = list
	}

	return lists
}

func ListToArray(head *ListNode) []int {
	n := head

	result := make([]int, 0)

	for n != nil {
		result = append(result, n.Val)

		n = n.Next
	}

	return result
}

// next functions are to copy to the solutions.
// they are made non-public to avoid name clashes when you import this complete package
func reverseList(head *ListNode) *ListNode { // returns the new head (was tail)
	//fmt.Println("Original list:")
	//listsCommon.PrintList(head)

	var previous *ListNode = nil
	current := head

	for current != nil {
		temp := current.Next
		current.Next = previous

		previous = current
		current = temp
	}

	head = previous

	//fmt.Println("Reversed list:")
	//listsCommon.PrintList(head)

	return head
}

func reverseListRecursive(node *ListNode) *ListNode {
	if (node == nil) || (node.Next == nil) {
		return node
	}

	// we're just returning the last node
	nextHead := reverseListRecursive(node.Next)

	node.Next.Next = node
	node.Next = nil

	return nextHead
}

func nextOrNil(node *ListNode) *ListNode {
	// avoids NPE failing on assigning
	// next = current.Next
	if node == nil {
		return nil
	}

	return node.Next
}

func valToString(node *ListNode) string {
	// avoids NPE failing on printing node.Val in the log
	if node == nil {
		return "nil"
	}

	return strconv.Itoa(node.Val)
}

func getListMiddle(head *ListNode) *ListNode {
	// even elements list: 1 - 2 - 3 - 4 -> return 2
	// odd elements list:  1 - 2 - 3 - 4 - 5 -> return 3

	// corner-cases:
	// 1 - 2 -> return 1
	// 1 -> return 1
	// nil -> return nil

	if head == nil {
		return nil
	}

	slow := head
	fast := head.Next // to make it work for both even add odd nodes count

	for (fast != nil) && (fast.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
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
