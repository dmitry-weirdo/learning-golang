package main

import "fmt"

func maximumBooks(books []int) int64 {
	// see https://www.youtube.com/watch?v=JDNraQ0QSfc

	// For the right index [i], we're searching for the leftmost index [j] where the "can take -1" progression breaks.
	// We can form a progression while a[j] >= a[i] - (i - j).
	// So the first element that breaks the progression is
	// a[j] < a[i] - (i - j).
	// a[j] < a[i] - i + j
	// a[j] - j < a[i] - i

	// So if we use the diff array: diff[i] = a[i] - i,
	// We can search for the "previous smaller index" - this is a typical monotonic stack problem.

	// It means that for the elements [prevSmaller + 1; i],
	// we can take the sum of the arithmetic progression from a[i] descending by i.
	// sum = (first + last) * n / 2

	// n is the count of element in the [prevSmaller + 1; i] interval.
	// n = i - (prevSmaller + 1) + 1 = i - prevSmaller - 1 + 1 = i - prevSmaller

	// last = a[i]

	// For first, we decrease a[i] by 1 * intervalLength:
	// first = a[i] - ( i - (prevSmaller + 1) ) = a[i] - (i - prevSmaller - 1) = a[i] - i + prevSmaller + 1

	// dp[i] is the maximum sum we can take from a subarray that ends on index [i].

	// If some element [i] does not have a prevSmaller, we can sum from it decreasing with -1,
	// until we either reach element value 1 or reach the beginning of the array.
	// Corner-case of this is dp[0] = a[0]

	// Example: [8 5 2]
	// From 5, we can sum (4 from 8) and (5 from 5), since we reached the start of the array.
	// From 2, we can sum (1 from 5) and (2 from 2). And we reached the 1 before reaching the start of the array.

	// I.e. in the interval, we will either have [i + 1] elements
	// or a[i] elements.
	//
	// So we take the minimum of it as the elements used:
	// intervalElements = min(i + 1, a[i])
	// first = a[i] - min(i + 1, a[i]) + 1

	// If [i] has a prevSmaller:
	// dp[i] = sum + dp[prevSmaller},
	// since we can continue with a smaller arithmetic progression ending on [prevSmaller]

	// calc diff array
	n := len(books)
	diff := make([]int, n)

	for i, v := range books {
		diff[i] = v - i
	}

	// calc prevSmaller on the diff array. Monotonic stack with indexes.
	prevSmallerWithIndexes := GetPrevSmallerWithIndexes(diff, -1)

	prevSmallerIndexes := OnlyIndexes(prevSmallerWithIndexes)

	//fmt.Printf("Diff array: %v \n", diff)
	//fmt.Printf("Previous smaller indexes on diff array: %v \n", prevSmallerIndexes)

	// calculate the DP array
	maxValue := books[0]

	dp := make([]int, n)
	dp[0] = books[0]

	for i, v := range books {
		prevSmaller := prevSmallerIndexes[i]

		if prevSmaller != -1 {
			elementsCount := i - prevSmaller            // [prevSmaller + 1, i]
			firstValue := v - elementsCount + 1         // we subtract 1 in every possible element. !!! It's NOT a[prevSmaller + 1], but a descending progression to this index
			sum := (firstValue + v) * elementsCount / 2 // arithmetic progression from firstValue to a[i]

			dp[i] = dp[prevSmaller] + sum // after we reached [prevSmaller + 1], we can continue the progression from a[prevSmaller], going -1 from it, etc.
		} else {
			elementsCount := min(i+1, v)                // [0; i] or [i - a[i] + 1; i]
			firstValue := v - elementsCount + 1         // we subtract 1 in every possible element. !!! It's NOT a[prevSmaller + 1], but a descending progression to this index
			sum := (firstValue + v) * elementsCount / 2 // arithmetic progression from firstValue to a[i]

			dp[i] = sum // no prevSmaller -> we only sum up from the current element to 1 or to the beginning of the array
		}

		maxValue = max(maxValue, dp[i])
	}

	return int64(maxValue)
}

func OnlyIndexes(a []MatchingElement) []int {
	result := make([]int, len(a))

	for i := range a {
		result[i] = a[i].index
	}

	return result
}

func GetPrevSmallerWithIndexes(a []int, noElementValue int) []MatchingElement { // same as GetNextSmaller, but left-to-right
	// direction: left -> right
	// stack: decreasing from top to bottom
	// removal from stack: >= current value
	// select top as result: if < current value
	// push current value to stack: always
	stack := createStackWithIndex()

	n := len(a)
	result := make([]MatchingElement, n)

	for i := 0; i < n; i++ {
		v := a[i]

		for stackIsNotEmptyWithIndex(stack) && (getStackTopWithIndex(stack).value >= v) {
			removeFromStackWithIndex(stack)
		}

		if stackIsNotEmptyWithIndex(stack) && (getStackTopWithIndex(stack).value < v) {
			result[i] = getStackTopWithIndex(stack)
		} else { // no next smaller element
			result[i] = MatchingElement{value: noElementValue, index: -1} // should default to -1
		}

		currentElement := MatchingElement{value: v, index: i}
		pushToStackWithIndex(stack, currentElement)
	}

	return result
}

type MatchingElement struct {
	value int
	index int // index of value in the array
}

// ======== stack of MatchingElement values ====== //
func createStackWithIndex() *[]MatchingElement {
	stack := make([]MatchingElement, 0)
	return &stack
}

func stackIsNotEmptyWithIndex(stack *[]MatchingElement) bool {
	return len(*stack) > 0
}

func stackIsEmptyWithIndex(stack *[]MatchingElement) bool {
	return len(*stack) == 0
}

func pushToStackWithIndex(stack *[]MatchingElement, v MatchingElement) { // pushes to the end of the stack
	// we push to the end of the slice = top of the stack
	*stack = append(*stack, v)
}

func removeFromStackWithIndex(stack *[]MatchingElement) MatchingElement { // removes from the top of the stack, only called when stack is not empty
	lastElement := (*stack)[len(*stack)-1]

	*stack = (*stack)[:len(*stack)-1] // remove the last element

	return lastElement
}

func getStackTopWithIndex(stack *[]MatchingElement) MatchingElement { // only called when stack is not empty
	return (*stack)[len(*stack)-1]
}

func test(arr []int, expectedResult int64) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array of books: %v \n", arr)

	result := maximumBooks(arr)

	fmt.Printf("Maximum consecutive books we can take as an increasing subarray: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]int{8, 5, 2, 7, 9},
		int64(19), // 1 + 2 + 7 + 9
	)
}

func test2() {
	test(
		[]int{7, 0, 3, 4, 5},
		int64(12), // 3 + 4 + 5
	)
}

func test3() {
	test(
		[]int{8, 2, 3, 7, 3, 4, 0, 1, 4, 3},
		int64(13), // 1 + 2 + 3 + 7
	)
}

func test4() {
	test(
		[]int{1, 1, 1, 2, 3},
		int64(6), // 1 + 2 + 3
	)
}

func test5() {
	test(
		[]int{1, 2, 3, 4}, // no prevSmaller test
		int64(10),         // 1 + 2 + 3 + 4
	)
}

func main() {
	// 2355. Maximum Number of Books You Can Take
	test1()
	test2()
	test3()
	test4()
	test5()
}
