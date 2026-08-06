package main

import (
	"container/list"
	"fmt"
)

func largestRectangleArea(heights []int) int {
	// todo: we can optimize and calculate both and right minimums in one O(N) run.
	// see https://algo.monster/liteproblems/84
	// I think this is a more complex code and still O(N), so not necessary.

	// For every element, we calculate the size of the rectangle where it is the limiting rectangle,
	// i.e., where it is the smallest or equal element in the rectangle.
	// The width limiting will be rectangles on both sides that are strictly smaller than the current element.
	// (since if the rectangle is smaller, it is already a limiting h[i], and we will calculate his maximum rectangle at his position).
	//
	// Then the width will be recalculated as right[i] - left[i] - 1.
	// Example: If height[1] is limited by height[0] and height[2], then the width will be 2 - 0 - 1 = 1.

	if len(heights) == 0 {
		return 0
	}

	leftLimits := GetPrevSmallerWithIndexes(heights, -1)
	rightLimits := GetNextSmallerWithIndexes(heights, -1) // todo: we should just pass noElementIndex = len(heights) here

	leftIndexes := OnlyIndexes(leftLimits)
	rightIndexes := OnlyIndexes(rightLimits)

	fmt.Printf("Left limiting indexes: %v \n", leftIndexes)
	fmt.Printf("Right limiting indexes: %v \n", rightIndexes)

	maxArea := -1

	for i, v := range heights {
		fmt.Println()
		fmt.Printf("height[%v] = %v. \n", i, v)

		leftIndex := leftIndexes[i]

		rightIndex := rightIndexes[i]
		if rightIndex == -1 { // if there is no right limit, virtual rectangle is after the end of the array
			rightIndex = len(heights)
		}

		fmt.Printf("leftIndex[%v] = %v. \n", i, leftIndex)
		fmt.Printf("rightIndex[%v] = %v. \n", i, rightIndex)

		width := rightIndex - leftIndex - 1
		area := width * v

		fmt.Printf("width: %v \n", width)
		fmt.Printf("area: %v \n", area)

		maxArea = max(maxArea, area)
	}

	return maxArea
}

type MatchingElement struct {
	value int
	index int // index of value in the array
}

func OnlyIndexes(a []MatchingElement) []int {
	result := make([]int, len(a))

	for i := range a {
		result[i] = a[i].index
	}

	return result
}

func GetNextSmallerWithIndexes(a []int, noElementValue int) []MatchingElement {
	// direction: right -> left
	// stack: decreasing from top to bottom
	// removal from stack: >= current value
	// select top as result: if < current value
	// push current value to stack: always
	stack := list.New()

	n := len(a)
	result := make([]MatchingElement, n)

	for i := n - 1; i >= 0; i-- {
		v := a[i]

		for (stack.Len() > 0) && (getStackTopWithIndex(stack).value >= v) {
			removeFromStackWithIndex(stack)
		}

		if (stack.Len() > 0) && (getStackTopWithIndex(stack).value < v) {
			result[i] = getStackTopWithIndex(stack)
		} else { // no next smaller element
			result[i] = MatchingElement{value: noElementValue, index: -1} // should default to -1
		}

		currentElement := MatchingElement{value: v, index: i}
		pushToStackWithIndex(stack, currentElement)
	}

	return result
}

func GetPrevSmallerWithIndexes(a []int, noElementValue int) []MatchingElement { // same as GetNextSmaller, but left-to-right
	// direction: left -> right
	// stack: decreasing from top to bottom
	// removal from stack: >= current value
	// select top as result: if < current value
	// push current value to stack: always
	stack := list.New()

	n := len(a)
	result := make([]MatchingElement, n)

	for i := 0; i < n; i++ {
		v := a[i]

		for (stack.Len() > 0) && (getStackTopWithIndex(stack).value >= v) {
			removeFromStackWithIndex(stack)
		}

		if (stack.Len() > 0) && (getStackTopWithIndex(stack).value < v) {
			result[i] = getStackTopWithIndex(stack)
		} else { // no next smaller element
			result[i] = MatchingElement{value: noElementValue, index: -1} // should default to -1
		}

		currentElement := MatchingElement{value: v, index: i}
		pushToStackWithIndex(stack, currentElement)
	}

	return result
}

func pushToStackWithIndex(stack *list.List, v MatchingElement) { // pushes to the end of the stack
	stack.PushFront(v)
}

func removeFromStackWithIndex(stack *list.List) MatchingElement { // removes from the top of the stack, only called when stack is not empty
	return stack.Remove(stack.Front()).(MatchingElement)
}

func getStackTopWithIndex(stack *list.List) MatchingElement { // only called when stack is not empty
	return stack.Front().Value.(MatchingElement)
}

func test(arr []int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Heights: %v \n", arr)

	result := largestRectangleArea(arr)

	fmt.Printf("Max rectangle in histogram area: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
		return
	}
}

func test1() {
	arr := []int{2, 1, 5, 6, 2, 3}
	expected := 10

	test(arr, expected)
}

func test2() {
	arr := []int{2, 4}
	expected := 4

	test(arr, expected)
}

func test3() {
	arr := []int{}
	expected := 0

	test(arr, expected)
}

func test4() {
	arr := []int{10}
	expected := 10

	test(arr, expected)
}

func main() {
	// 84. Largest Rectangle in Histogram
	test1()
	test2()
	test3()
	test4()
}
