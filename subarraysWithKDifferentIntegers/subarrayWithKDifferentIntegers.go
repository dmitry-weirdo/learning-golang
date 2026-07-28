package main

import "fmt"

func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithKDistinct_threePointerslidingWindow(nums, k)
	//return subarraysWithKDistinct_slidingWindow(nums, k)
	//return subarraysWithKDistinct_kAndKMinusOne(nums, k)
}

func subarraysWithKDistinct_threePointerslidingWindow(nums []int, k int) int {
	// see https://www.youtube.com/watch?v=etI6HqWVa8U

	// basically the same as (k-1) leftmost, k leftmost.
	// leftFar is kLeftmost
	// leftNear is kMinusOneLeftMost - 1 (i.e. the next element is already (k - 1) distinct elements from right

	leftNear := 0 // moves to the last element before "k distinct from the current right" will become "k - 1 distinct from the current right)
	leftFar := 0  // stays at the left where "k distinct for the current right" started

	freq := make(map[int]int)

	count := 0

	for right := range nums {
		numRight := nums[right]

		freq[numRight]++

		for len(freq) > k {
			// too many chars ->  shrink from left, move leftFar also to the same pointer
			numLeft := nums[leftNear]

			freq[numLeft]--

			if freq[numLeft] == 0 {
				delete(freq, numLeft)
			}

			leftNear++
			leftFar = leftNear
		}

		// most tricky thing -> move leftNear to right until it reaches a point
		// where the frequency of that char is 1, i.e. the next element is "leftmost of k - 1"
		numLeft := nums[leftNear]

		for freq[numLeft] > 1 {
			freq[numLeft]--

			if freq[numLeft] == 0 { // this should not happen since we're only moving if freq[numLeft] > 1
				delete(freq, numLeft)
			}

			leftNear++
			numLeft = nums[leftNear]
		}

		if len(freq) == k {
			// we're basically adding ("k - 1 leftmost" - "k leftmost")
			count += leftNear - leftFar + 1
		}
	}

	return count
}

func subarraysWithKDistinct_slidingWindow(nums []int, k int) int {
	// no, this won't work - we cannot shrink from left since right can still expand without increasing distinct characters
	left := 0

	freq := make(map[int]int)

	count := 0

	for right := 0; right < len(nums); right++ {
		numRight := nums[right]

		freq[numRight]++

		if len(freq) > k { // shrink from left
			numLeft := nums[left]
			freq[numLeft]--

			if freq[numLeft] == 0 {
				delete(freq, numLeft)
			}
		}

		if len(freq) == k {
			for len(freq) == k {
				count++
				fmt.Printf("Adding matching array [%v; %v] = %v. New count = %v. \n", left, right, nums[left:right+1], count)

				numLeft := nums[left]
				freq[numLeft]--

				if freq[numLeft] == 0 {
					delete(freq, numLeft)
				}

				left++
			}

			count++
		}

	}

	return count
}

func subarraysWithKDistinct_kAndKMinusOne(nums []int, k int) int {
	l := len(nums)

	// todo: we can just calculate for every pos and not store the complete arrays, but we need to store all variable from the function
	kMinusOnePositions := getLeftPos(nums, k-1)
	fmt.Printf("At most (k - 1) = %v distinct values - leftmost positions: %v \n", k-1, kMinusOnePositions)

	kPositions := getLeftPos(nums, k)
	fmt.Printf("At most (k) = %v distinct values - leftmost positions: %v \n", k, kPositions)

	result := 0

	for i := range l {
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
			fmt.Printf("Left index: %v \n", left)

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

func test(nums []int, k int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("nums: %v \n", nums)
	fmt.Printf("k: %v \n", k)

	result := subarraysWithKDistinct(nums, k)

	fmt.Printf("Expected count of subarrays: %v \n", expectedResult)
	fmt.Printf("Result count of subarrays: %v \n", result)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
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

func test3() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 3

	expected := 3

	test(nums, k, expected)
}

func main() {
	// 992. Subarrays with K Different Integers
	test1()
	test2()
	test3()
}
