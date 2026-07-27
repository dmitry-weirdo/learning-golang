package main

import "fmt"

type MinStack struct {
	data []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	var newMin int

	if len(this.data) > 0 {
		newMin = min(val, this.mins[len(this.mins)-1])
	} else {
		newMin = val
	}

	this.data = append(this.data, val)
	this.mins = append(this.mins, newMin)
}

func (this *MinStack) Pop() {
	lastIndex := len(this.data) - 1
	//lastValue := this.data[lastIndex]

	this.data = this.data[0:lastIndex]
	this.mins = this.mins[0:lastIndex]

	//return lastValue
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

func getExpectedTop(st *MinStack, expectedTop int) {
	actualTop := st.Top()

	fmt.Printf("Expected top: %v \n", expectedTop)
	fmt.Printf("Actual top: %v \n", actualTop)

	if actualTop != expectedTop {
		fmt.Printf("FAILURE: expected top = %v, actual top = %v \n", expectedTop, actualTop)
	}
}

func getExpectedMin(st *MinStack, expectedMin int) {
	actualMin := st.GetMin()

	fmt.Printf("Expected min: %v \n", expectedMin)
	fmt.Printf("Actual min: %v \n", actualMin)

	if actualMin != expectedMin {
		fmt.Printf("FAILURE: expected min = %v, actual min = %v \n", expectedMin, actualMin)
	}
}

func test() {
	minStack := Constructor()
	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(0)

	getExpectedMin(&minStack, 0)
	//minStack.GetMin() // return 0

	minStack.Pop()

	getExpectedTop(&minStack, 2)
	//minStack.Top() // return 2

	getExpectedMin(&minStack, 1)
	// minStack.GetMin() // return 1
}

func main() {
	// 155. Min Stack
	test()
}
