package main

import "fmt"

func waysToMakeFair(nums []int) int {
	return waysToMakeFair_optimized(nums) // optimized, O(1) space, O(2*N) time. Saves time on arrays allocation.
	//return waysToMakeFair_trivial(nums)   // my trivial solution, 2 arrays of prefix sums, O(2*N) space, O(2*N) time
}

func waysToMakeFair_optimized(nums []int) int {
	// we don't need to store the prefixSum arrays
	// we just need the total sums, and then we can get the result with just running sums for even and odd elements

	// calculate total sums (this is prefixSumsEven[n] and prefixSumsOdd[n])
	totalEven := 0
	totalOdd := 0

	for i, v := range nums {
		if i%2 == 0 {
			totalEven += v
		} else {
			totalOdd += v
		}
	}

	//fmt.Printf("Total even: %v \n", totalEven)
	//fmt.Printf("Total odd: %v \n", totalOdd)

	sumEven := 0
	sumOdd := 0

	count := 0

	for i, v := range nums {
		evenBefore := sumEven // - 0 that is prefixSumsEven[0]
		oddBefore := sumOdd   //  - 0 that is prefixSumsOdd[0]

		// add the current number to get prefixSumsEven[i + 1] and prefixSumsOdd[i + 1}
		if i%2 == 0 {
			sumEven += v
		} else {
			sumOdd += v
		}

		evenAfter := totalEven - sumEven
		oddAfter := totalOdd - sumOdd

		if evenBefore+oddAfter == oddBefore+evenAfter {
			count++
		}
	}

	return count
}

func waysToMakeFair_trivial(nums []int) int {
	n := len(nums)

	//prefixSumsEven := getPrefixSumsEven(nums)
	//prefixSumsOdd := getPrefixSumsOdd(nums)
	prefixSumsEven, prefixSumsOdd := getPrefixSumsEvenAndOdd(nums)

	//fmt.Printf("Prefix sums even: %v \n", prefixSumsEven)
	//fmt.Printf("Prefix sums odd: %v \n", prefixSumsOdd)

	count := 0

	for i := range nums { // try to remove every index
		evenBefore := prefixSumsEven[i] - prefixSumsEven[0] // prefixSumsEven[0] is actually 0
		oddBefore := prefixSumsOdd[i] - prefixSumsOdd[0]    // prefixSumsOdd[0] is actually 0

		evenAfter := prefixSumsEven[n] - prefixSumsEven[i+1]
		oddAfter := prefixSumsOdd[n] - prefixSumsOdd[i+1]

		// After removing the element,
		// evenAfter becomes oddAfter,
		// oddAfter becomes evenAfter.

		if evenBefore+oddAfter == oddBefore+evenAfter {
			count++
		}
	}

	return count
}

func getPrefixSumsEvenAndOdd(a []int) (even, odd []int) { // calculate both sums in O(n) run instead of 2 separate runs
	// to not mix up with indexes conversion, let's count usual prefixSums just with 0 on even positions

	// with prefixSums[0] = 0, sum up to a[i] will be in prefixSums[i+1]
	prefixSumsEven := make([]int, len(a)+1)
	prefixSumsOdd := make([]int, len(a)+1)

	prefixSumsEven[0] = 0
	prefixSumsOdd[0] = 0

	for i, v := range a {
		prefixSumIndex := i + 1

		if i%2 == 0 {
			// even position -> add current value
			prefixSumsEven[prefixSumIndex] = prefixSumsEven[i] + v

			// even position -> skip current value
			prefixSumsOdd[prefixSumIndex] = prefixSumsOdd[i]
		} else {
			// odd position -> skip current value
			prefixSumsEven[prefixSumIndex] = prefixSumsEven[i]

			// odd position -> add current value
			prefixSumsOdd[prefixSumIndex] = prefixSumsOdd[i] + v
		}
	}

	return prefixSumsEven, prefixSumsOdd
}

func getPrefixSumsEven(a []int) []int {
	// to not mix up with indexes conversion, let's count usual prefixSums just with 0 on even positions

	// with prefixSums[0] = 0, sum up to a[i] will be in prefixSums[i+1]
	prefixSums := make([]int, len(a)+1)

	prefixSums[0] = 0

	for i, v := range a {
		prefixSumIndex := i + 1

		if i%2 == 0 { // even position -> add current value
			prefixSums[prefixSumIndex] = prefixSums[i] + v
		} else { // odd position -> skip current value
			prefixSums[prefixSumIndex] = prefixSums[i]
		}
	}

	return prefixSums
}

func getPrefixSumsOdd(a []int) []int {
	// to not mix up with indexes conversion, let's count usual prefixSums just with 0 on odd positions

	// with prefixSums[0] = 0, sum up to a[i] will be in prefixSums[i+1]
	prefixSums := make([]int, len(a)+1)

	prefixSums[0] = 0

	for i, v := range a {
		prefixSumIndex := i + 1

		if i%2 != 0 { // odd position -> add current value
			prefixSums[prefixSumIndex] = prefixSums[i] + v
		} else { // even position -> skip current value
			prefixSums[prefixSumIndex] = prefixSums[i]
		}
	}

	return prefixSums
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr) // todo: replace with your text

	result := waysToMakeFair(arr)

	fmt.Printf("Count of fair arrays by removing 1 element: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{2, 1, 6, 4}
	expected := 1

	test(arr, expected)
}

func test2() {
	arr := []int{1, 1, 1}
	expected := 3

	test(arr, expected)
}

func test3() {
	arr := []int{1, 2, 3}
	expected := 0

	test(arr, expected)
}

func main() {
	// 1664. Ways to Make a Fair Array
	test1()
	test2()
	test3()
}
