package main

import (
	"container/list"
	"fmt"
)

type IndexAndValue struct {
	Index int
	Value int
}

func maxSlidingWindow(nums []int, k int) []int {
	// todo: we don't handle k < len(nums) case
	result := make([]int, 0)

	queue := list.New()

	// put the first K elements
	for i := 0; i < k; i++ {
		// remove all the elements <= than the current element from the end of the queue
		v := nums[i]

		for queue.Len() > 0 && queue.Back().Value.(*IndexAndValue).Value < v {
			removeLastFromQueue(queue)
		}

		value := &IndexAndValue{i, nums[i]}
		appendToQueue(queue, value)

		fmt.Printf("Index[%v]. Queue: \n", i)
		printQueue(queue)
	}

	for i := k; i < len(nums); i++ {
		fmt.Println("==========================")

		result = append(result, queue.Front().Value.(*IndexAndValue).Value)

		left := i - k + 1
		right := i
		fmt.Printf("left: %v, right: %v \n", left, right)

		// remove all max elements that are out of the current range
		for queue.Len() > 0 && queue.Front().Value.(*IndexAndValue).Index < left {
			removeFirstFromQueue(queue)
		}

		fmt.Printf("Index[%v]. Removed out-of-the-window max elements (index < %d) from the queue: \n", i, left)
		printQueue(queue)

		// current element
		current := nums[i]
		fmt.Printf("Current element [%v]: %v \n", i, current)

		for queue.Len() > 0 && queue.Back().Value.(*IndexAndValue).Value < current {
			removeLastFromQueue(queue)
		}

		value := &IndexAndValue{i, current}
		appendToQueue(queue, value)

		fmt.Printf("Appended the current element index = %v, value = %v to the queue. Queue: \n", i, current)
		printQueue(queue)
	}

	// todo: last result should still be returned
	result = append(result, queue.Front().Value.(*IndexAndValue).Value)

	return result
}

func removeFirstFromQueue(queue *list.List) *IndexAndValue {
	return queue.Remove(queue.Front()).(*IndexAndValue)
}

func removeLastFromQueue(queue *list.List) *IndexAndValue {
	return queue.Remove(queue.Back()).(*IndexAndValue)
}

func printQueue(queue *list.List) {
	v := queue.Front()

	for v != nil {
		fmt.Printf("%v -> %v ", v.Value.(*IndexAndValue).Index, v.Value.(*IndexAndValue).Value)

		v = v.Next()
	}

	fmt.Println()
}

func getFirstFromQueue(queue *list.List) *IndexAndValue {
	return queue.Front().Value.(*IndexAndValue)
}

func appendToQueue(queue *list.List, s *IndexAndValue) {
	queue.PushBack(s)
}

func test(nums []int, k int, expected []int) {
	fmt.Println()
	fmt.Println("=======================")
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("k (sliding window size): %v\n", k)

	result := maxSlidingWindow(nums, k)

	fmt.Printf("Expected sliding window maximums: %v\n", expected)
	fmt.Printf("Actual sliding window maximums: %v\n", result)
}

func test1() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	expected := []int{3, 3, 5, 5, 6, 7}

	test(nums, k, expected)
}

func main() {
	test1()

}
