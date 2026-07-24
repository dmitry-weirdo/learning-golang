package listsCommon

import "fmt"

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
