package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head           // write position for odd nodes
	firstEven := odd.Next // cache the first even node, we will append the odd nodes before it
	lastEven := odd.Next  // last even, from this we jump 2 steps to find the next odd
	nextOdd := odd        // it moves 2 steps ahead (if possible) to find the next odd node

	for (nextOdd != nil) && (nextOdd.Next) != nil && (nextOdd.Next.Next != nil) { // next odd node exists
		nextOdd = nextOdd.Next.Next

		// fmt.Printf("Next odd node: %v \n", nextOdd.Val)

		// save the link to the next even node
		nextEven := nextOdd.Next

		// append next odd node to the current odd write position
		odd.Next = nextOdd
		nextOdd.Next = firstEven

		// move odd writing position one node ahead
		odd = nextOdd

		// connect next even element to the last even
		lastEven.Next = nextEven

		// since we're jumping 2 nodes ahead from the nextOdd, set nextOdd to previous lastEven, this was a tail of evens before the moved nextOdd
		nextOdd = lastEven

		// lastEven is now the next even
		lastEven = nextEven
	}

	return head // 1st node will stay the same, since it's odd
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n")
	listsCommon.PrintList(list)

	result := oddEvenList(list)
	resultAsArray := listsCommon.ListToArray(result)

	fmt.Printf("List with odd-even reordering: \n")
	listsCommon.PrintList(result)

	fmt.Printf("Result as array: %v \n", resultAsArray)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(resultAsArray) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(resultAsArray))
		return
	}

	for i, v := range resultAsArray {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{1, 3, 5, 2, 4}

	test(arr, expected)
}

func test2() {
	arr := []int{2, 1, 3, 5, 6, 4, 7}
	expected := []int{2, 3, 6, 7, 1, 5, 4}

	test(arr, expected)
}

func test3() {
	arr := []int{}
	expected := []int{}

	test(arr, expected)
}

func test4() {
	arr := []int{1}
	expected := []int{1}

	test(arr, expected)
}

func test5() {
	arr := []int{1, 2}
	expected := []int{1, 2}

	test(arr, expected)
}

func test6() {
	arr := []int{1, 2, 3}
	expected := []int{1, 3, 2}

	test(arr, expected)
}

func main() {
	// 328. Odd Even Linked List
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
