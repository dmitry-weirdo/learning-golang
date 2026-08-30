package main

import "fmt"

type MyCircularQueue struct {
	capacity int // we can use len(a), this cached value is for nicer code
	a        []int
	size     int
	head     int
	tail     int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		capacity: k,
		a:        make([]int, k),
		size:     0,
		head:     0,
		tail:     -1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool { // appends to tail if there is space
	if this.IsFull() { // capacity full -> cannot add
		return false
	}

	this.size++

	this.tail++
	this.tail %= this.capacity // if we went over the last element, start from the beginning of the array

	this.a[this.tail] = value

	return true
}

func (this *MyCircularQueue) DeQueue() bool { // removes from head if there are any elements in the queue
	if this.IsEmpty() { // nothing to delete
		return false
	}

	this.size--

	if this.IsEmpty() {
		// if queue is empty -> reset to the initial state to append from the start of the array again
		this.head = 0
		this.tail = -1
	} else {
		this.head++
		this.head %= this.capacity // if we went over the last element, start from the beginning of the array
	}

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.a[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.a[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size <= 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size >= this.capacity
}

func (this *MyCircularQueue) Size() int {
	return this.size
}

// ==================== Test methods ==================== //
func testEnQueue(q *MyCircularQueue, v int, expectedResult bool, expectedSize int) {
	fmt.Println()
	fmt.Println("========================")

	result := q.EnQueue(v)

	fmt.Printf("Tried to enqueue value %v. \n", v)
	fmt.Printf("Result: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}

	size := q.Size()

	fmt.Printf("Queue size: %v \n", size)
	fmt.Printf("Expected queue size: %v \n", expectedSize)

	if size != expectedSize {
		fmt.Printf("FAILURE: expected size = %v, actual size = %v \n", expectedSize, size)
	}
}

func testDeQueue(q *MyCircularQueue, expectedResult bool, expectedSize int) {
	fmt.Println()
	fmt.Println("========================")

	result := q.DeQueue()

	fmt.Printf("Tried to dequeue a value. \n")
	fmt.Printf("Result: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}

	size := q.Size()

	fmt.Printf("Queue size: %v \n", size)
	fmt.Printf("Expected queue size: %v \n", expectedSize)

	if size != expectedSize {
		fmt.Printf("FAILURE: expected size = %v, actual size = %v \n", expectedSize, size)
	}
}

func testRear(q *MyCircularQueue, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	sizeBefore := q.Size()

	result := q.Rear()

	sizeAfter := q.Size()

	fmt.Printf("Rear of the queue: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}

	if sizeAfter != sizeBefore { // size must NOT change on .Rear() operation
		fmt.Printf("FAILURE: expected size = %v, actual size = %v \n", sizeBefore, sizeAfter)
	}
}

func testIsFull(q *MyCircularQueue, expectedResult bool) {
	fmt.Println()
	fmt.Println("========================")

	result := q.IsFull()

	fmt.Printf("Queue is full: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	// Input
	//["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
	// [[3], [1], [2], [3], [4], [], [], [], [4], []]

	// Output
	// [null, true, true, true, false, 3, true, true, true, 4]

	myCircularQueue := Constructor(3)
	q := &myCircularQueue

	testEnQueue(q, 1, true, 1)  // 1
	testEnQueue(q, 2, true, 2)  // 1, 2
	testEnQueue(q, 3, true, 3)  // 1, 2, 3
	testEnQueue(q, 4, false, 3) // over the capacity -> remains 1, 2, 3

	testRear(q, 3)
	testIsFull(q, true)

	testDeQueue(q, true, 2) // 3 removed -> 1, 2

	testEnQueue(q, 4, true, 3) // enqueue is now allowed
	testRear(q, 4)
}

func main() {
	// 622. Design Circular Queue
	test1()
}
