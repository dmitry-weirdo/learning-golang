package main

import (
	"container/list"
	"fmt"
)

// stack of int values
type Stack struct {
	values *list.List
}

func NewStack() Stack {
	return Stack{
		values: list.New(),
	}
}

func (this *Stack) PushToTop(x int) {
	this.values.PushFront(x)
}

func (this *Stack) PopFromTop() int {
	if this.values.Len() <= 0 {
		panic("Stack is empty, nothing to pop")
	}

	front := this.values.Front()

	this.values.Remove(front)

	return front.Value.(int)
}

func (this *Stack) Peek() int {
	if this.values.Len() <= 0 {
		panic("Stack is empty, nothing at peek")
	}

	return this.values.Front().Value.(int)
}

func (this *Stack) Size() int {
	return this.values.Len()
}

func (this *Stack) IsEmpty() bool {
	return this.Size() <= 0
}

type MyQueue struct {
	front int // cash of queue front (first inserted), just for s1! I.e. the bottom element of s1
	s1    Stack
	s2    Stack // stack with queue order
}

func Constructor() MyQueue {
	return MyQueue{
		s1: NewStack(),
		s2: NewStack(),
	}
}

func (this *MyQueue) Push(x int) {
	if this.s1.IsEmpty() {
		this.front = x
	}

	// we push always to s1
	this.s1.PushToTop(x)
}

func (this *MyQueue) Pop() int {
	if this.s2.IsEmpty() {
		// Move elements from s1 to s2 - O(n)
		for !this.s1.IsEmpty() {
			this.s2.PushToTop(this.s1.PopFromTop())
		}

	}

	// we always pop from the top of s2
	return this.s2.PopFromTop()
}

func (this *MyQueue) Peek() int {
	if this.s2.IsEmpty() {
		return this.front // return the bottom of s1
	}

	return this.s2.Peek() // return top of s2
}

func (this *MyQueue) Empty() bool {
	return this.s1.IsEmpty() && this.s2.IsEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func testStack() {
	stack := NewStack()

	var v int

	v = 1
	stack.PushToTop(v)
	fmt.Printf("Pushed %v to stack. \n", v)

	v = 2
	stack.PushToTop(v)
	fmt.Printf("Pushed %v to stack. \n", v)

	v = 3
	stack.PushToTop(v)
	fmt.Printf("Pushed %v to stack. \n", v)

	fmt.Printf("Stack size: %v \n", stack.Size())
	fmt.Printf("Peek: %v \n", stack.Peek())

	var popped int

	popped = stack.PopFromTop()
	fmt.Printf("Popped: %v \n", popped)

	popped = stack.PopFromTop()
	fmt.Printf("Popped: %v \n", popped)

	popped = stack.PopFromTop()
	fmt.Printf("Popped: %v \n", popped)

	fmt.Printf("Stack size: %v \n", stack.Size())
	fmt.Printf("Stack is empty: %v \n", stack.IsEmpty())
}

func testQueue() {
	queue := Constructor()

	var v int

	v = 1
	queue.Push(v)
	fmt.Printf("Pushed %v to queue. \n", v)

	v = 2
	queue.Push(v)
	fmt.Printf("Pushed %v to queue. \n", v)

	v = 3
	queue.Push(v)
	fmt.Printf("Pushed %v to queue. \n", v)

	fmt.Printf("Queue is empty: %v \n", queue.Empty())
	fmt.Printf("Peek: %v \n", queue.Peek())

	var popped int

	popped = queue.Pop()
	fmt.Printf("Popped: %v \n", popped)

	popped = queue.Pop()
	fmt.Printf("Popped: %v \n", popped)

	popped = queue.Pop()
	fmt.Printf("Popped: %v \n", popped)

	fmt.Printf("Queue is empty: %v \n", queue.Empty())

	v = 4
	queue.Push(v)
	fmt.Printf("Pushed %v to queue. \n", v)

	v = 5
	queue.Push(v)
	fmt.Printf("Pushed %v to queue. \n", v)

	fmt.Printf("Peek: %v \n", queue.Peek())

	popped = queue.Pop() // move to s2
	fmt.Printf("Popped: %v \n", popped)

	v = 6
	queue.Push(v)
	fmt.Printf("Pushed %v to queue. \n", v)

	v = 7
	queue.Push(v)
	fmt.Printf("Pushed %v to queue. \n", v)

	popped = queue.Pop() // top of s2
	fmt.Printf("Popped: %v \n", popped)

	popped = queue.Pop() // top of s2
	fmt.Printf("Popped: %v \n", popped)

	popped = queue.Peek() // front of s1
	fmt.Printf("Peek: %v \n", popped)
}

func main() {
	// 232. Implement Queue using Stacks
	//testStack()
	testQueue()
}
