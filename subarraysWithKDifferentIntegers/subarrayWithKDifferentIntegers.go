package main

import "fmt"

func subarraysWithKDistinct(nums []int, k int) int {
	l := len(nums)

	// todo: we can just calculate for every pos and not store the complete arrays, but we need to store all variable from the function
	kMinusOnePositions := getLeftPos(nums, k-1)
	fmt.Printf("(k - 1) = %v positions: %v \n", k-1, kMinusOnePositions)

	kPositions := getLeftPos(nums, k)
	fmt.Printf("(k) = %v positions: %v \n", k, kPositions)

	result := 0

	for i := 0; i < l; i++ {
		result += kMinusOnePositions[i] - kPositions[i]
	}

	return result
}

func getLeftPos(nums []int, k int) []int {
	l := len(nums)

	// number to frequency
	frequencies := make(map[int]int)

	leftPositions := make([]int, l)

	// left pointer moves only right for every right positions, sliding window
	left := 0

	distinctCount := 0 // we can also just use len(frequencies)

	// calculate for every right position
	for right := 0; right < l; right++ {
		fmt.Println("=========================")
		fmt.Printf("K = %v, Right index: %v \n", k, right)

		rightValue := nums[right]

		// add the right to the frequencies
		_, ok := frequencies[rightValue]
		if !ok {
			frequencies[rightValue] = 1

			distinctCount++

			fmt.Printf("Added rightValue = %v as a new frequency.\n", rightValue)
		} else {
			frequencies[rightValue] += 1

			fmt.Printf("RightValue = %v increased frequency to %v.\n", rightValue, frequencies[rightValue])
		}

		// move the left pointers until we reach the distinctCount <= k
		for distinctCount > k {
			leftValue := nums[left]
			fmt.Printf("Left index: %v \n", right)

			_, ok = frequencies[leftValue]
			if !ok {
				// this should never happen
				fmt.Printf("!!! Left value %v is not found in the frequencies map. \n", leftValue)
			} else {
				frequencies[leftValue] -= 1

				fmt.Printf("LeftValue = %v decreased frequency to %v.\n", leftValue, frequencies[leftValue])

				if frequencies[leftValue] <= 0 {
					distinctCount--
					delete(frequencies, leftValue)

					fmt.Printf("Removed leftValue = %v with frequency 0 from the frequencies. \n", leftValue)
				}
			}

			left++
		}

		// put the leftmost index to the result
		leftPositions[right] = left
	}

	return leftPositions
}

func test(nums []int, k int, expected int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("nums: %v \n", nums)
	fmt.Printf("k: %v \n", k)

	result := subarraysWithKDistinct(nums, k)

	fmt.Printf("Expected count of subarrays: %v \n", expected)
	fmt.Printf("Result count of subarrays: %v \n", result)
}

func test1() {
	nums := []int{1, 2, 1, 2, 3}
	k := 2

	expected := 7

	test(nums, k, expected)
}

func test2() {
	nums := []int{1, 2, 1, 3, 4}
	k := 3

	expected := 3

	test(nums, k, expected)
}

func main() {
	test1()
	test2()
}
