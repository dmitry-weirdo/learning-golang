package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Val: -666, Next: head}

	slow := dummyHead
	fast := dummyHead

	for i := 0; i < n; i++ {
		// move fast N steps ahead of slow
		fast = fast.Next
	}

	fmt.Printf("Moved fast %v nodes ahead of slow. \n", n)
	fmt.Printf("Slow: %v \n", slow.Val)
	fmt.Printf("Fast: %v \n", fast.Val)

	// now move fast to the end of the list
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	fmt.Printf("fast.Next reached the end of the list. \n")
	fmt.Printf("Slow: %v \n", slow.Val)
	fmt.Printf("Fast: %v \n", fast.Val)

	// remove the slow.Next element
	fmt.Printf("Removing the next element = %v from slow.Next... \n", slow.Next.Val)
	slow.Next = slow.Next.Next

	return dummyHead.Next
}

func test(arr []int, n int) {
	fmt.Println()
	fmt.Println("====================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n")
	listsCommon.PrintList(list)
	fmt.Printf("N: %v \n", n)

	updatedList := removeNthFromEnd(list, n)

	fmt.Printf("List with %v-th element from the end removed: \n", n)
	listsCommon.PrintList(updatedList)
}

func test1() {
	arr := []int{1, 2, 3, 4, 5}
	n := 2

	test(arr, n)
}

func test2() {
	arr := []int{1, 2}
	n := 1

	test(arr, n)
}

func test3() {
	arr := []int{1}
	n := 1

	test(arr, n)
}

func main() {
	// 19. Remove Nth Node From End of List
	test1()
	test2()
	test3()
}
