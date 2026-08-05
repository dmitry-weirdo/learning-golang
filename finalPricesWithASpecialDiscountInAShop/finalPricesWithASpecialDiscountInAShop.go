package main

import (
	"container/list"
	"fmt"
)

func finalPrices(prices []int) []int {
	// For discounts, we're calculating the "next <= element" via monotonic stack.
	// Then we just subtract prices - discounts.

	n := len(prices)

	discounts := make([]int, n)

	// monotonic stack - top is > next etc.
	stack := list.New()

	for i := n - 1; i >= 0; i-- {
		v := prices[i]

		fmt.Println()
		fmt.Printf("prices[%v] = %v \n", i, v)

		fmt.Println("Stack: ")
		printStack(stack)

		// remove all values > current element
		// These values are worse than the current for all the elements to the left of the current since:
		// - current element is <=
		// - current element is more at the left

		// Values smaller than the current value remain in the stack since they can be smaller element
		// for the values at the left that are smaller than current.
		for (stack.Len() > 0) && (getStackTop(stack) > v) {
			removeFromStack(stack)
		}

		fmt.Printf("Removed all values > %v from the stack: \n", v)
		printStack(stack)

		// top of the stack is the biggest nearest candidate AFTER the current element
		// that is <= than the current element (elements bigger than the current we removed above)
		// (since we're going from right to left)
		if (stack.Len() > 0) && (getStackTop(stack) <= v) {
			discounts[i] = getStackTop(stack)
		} else { // no next <= element -> no discount
			discounts[i] = 0
		}

		// Push the current element to the stack if it is != top.
		// All the remaining values (if present) are smaller than the current element,
		// So the stack remains strictly decreasing from the top.
		// The current element still can be a valid "next <=" for the values at the left that are bigger or equal.
		if (stack.Len() == 0) || (getStackTop(stack) != v) {
			pushToStack(stack, v)

			fmt.Printf("Pushed current value %v to the stack: \n", v)
			printStack(stack)
		} else {
			fmt.Printf("Current stack top is equal to the current value %v. Do not push anything to the stack. \n", v)
		}
	}

	fmt.Printf("Discounts: %v \n", discounts)

	// for the result, we're just subtracting prices[i] - discounts[i]
	// todo: we can use the original array if required to save memory
	result := make([]int, n)

	for i := range n {
		result[i] = prices[i] - discounts[i]
	}

	return result
}

func pushToStack(stack *list.List, v int) { // pushes to the end of the stack
	stack.PushFront(v)
}

func removeFromStack(stack *list.List) int { // removes from the top of the stack, only called when stack is not empty
	return stack.Remove(stack.Front()).(int)
}

func getStackTop(stack *list.List) int { // only called when stack is not empty
	return stack.Front().Value.(int)
}

func printStack(l *list.List) {
	fmt.Printf("[ ")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}

	fmt.Printf("]")

	fmt.Println()
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array of prices: %v \n", arr)

	result := finalPrices(arr)

	fmt.Printf("Prices with discounts (substracting next smaller or equal element): %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	arr := []int{8, 4, 6, 2, 3}
	expectedResult := []int{4, 2, 4, 2, 3}

	test(arr, expectedResult)
}

func test2() {
	arr := []int{1, 2, 3, 4, 5}
	expectedResult := []int{1, 2, 3, 4, 5}

	test(arr, expectedResult)
}

func test3() {
	arr := []int{10, 1, 1, 6}
	expectedResult := []int{9, 0, 1, 6}

	test(arr, expectedResult)
}

func main() {
	// 1475. Final Prices With a Special Discount in a Shop
	// Similar to "496. Next Greater Element I", but we're searching for the opposite -> next smaller or equal element
	test1()
	test2()
	test3()
}
