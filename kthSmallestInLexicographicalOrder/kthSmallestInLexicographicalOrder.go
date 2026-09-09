package main

import "fmt"

func findKthNumber(n int, k int) int {

	current := 1 // start counting from 1
	k--          // convert to 0-indexed, we're starting from number 1

	for k > 0 {
		currentSubtreeCount := getCount(n, current)

		//fmt.Println()
		//fmt.Printf("Remaining K: %v. Current node: %v. Subtree of %v count: %v. \n", k, current, current, currentSubtreeCount)

		if k >= currentSubtreeCount {
			//fmt.Printf("Skipping the subtree of %v, skipping total %v nodes. \n", current, currentSubtreeCount)

			// we still need more numbers than the current subtree
			//-> skip this subtree, decrease the remaining count by count of elements of this subtree
			k -= currentSubtreeCount

			// go to the next sibling node
			current++
		} else {
			//fmt.Printf("Going to the next level of %v = %v, skipping 1 node %v. \n", current, current*10, current)

			// k-th number within the current subtree -> search on the level below of this subtree
			k-- // count the current node

			current *= 10
		}
	}

	return current
}

func getCount(n int, current int) int {
	count := 0

	// n = 1 -> nextSibling = 2

	// If we start a level below:
	// n = 10, nextSibling = 11
	// Then we will descend to 100 / 110
	nextSibling := current + 1

	for current <= n {
		// (n - current + 1) are current level up to N.
		// e.g. n = 13, current = 10 -> (n - current + 1) = 4 (it will be 10, 11, 12, 13)
		currentLevelUpToN := n - current + 1

		// (nextSibling - current) are all nodes at the current level.
		// e.g. n = 30, current = 10, nextSibling = 20 -> (nextSibling - current) = 10 (10, 11, ... , 19)
		completeCurrentLevelUpToNextSibling := nextSibling - current

		// we can either go within the current level up to N or take all nodes of the current level
		count += min(currentLevelUpToN, completeCurrentLevelUpToNextSibling)

		// move to the next level from current, i.e. multiply by 1 (1 -> 10 - > 100 -> ...)
		// on level 10, current parent siblings will be [10; 19], next sibling will be 20.
		current *= 10     // 1 -> 10 -> 100 -> ...
		nextSibling *= 10 // 2 -> 20 -> 200 -> ...
	}

	return count
}

func test(n, k int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number N: %v \n", n)
	fmt.Printf("K - to find K-th in lexicographical order up to N: %v \n", k)

	result := findKthNumber(n, k)

	fmt.Printf("%v-th lexicographical number up to %v: %v \n", k, n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(13, 2, 10) // 1, 10
}

func test2() {
	test(1, 1, 1) // 1
}

func test3() {
	test(100, 10, 17) // 1, 10, 100, 11, 12, 13, 14, 15, 16, 17
}

func main() {
	// 720. Longest Word in Dictionary
	test1()
	test2()
	test3()
}
