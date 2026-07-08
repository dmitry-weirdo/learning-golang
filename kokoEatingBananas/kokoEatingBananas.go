package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	// left is 1 per hour (problem conditions)
	// right is max(piles), since h >= len(piles)

	// n - total piles
	// find max time - O(n)
	// binary search for every of n piles - O(n * log maxPile) >> O(n)
	// Total time = O(n * log maxPile)

	left := 1
	right := piles[0]

	for _, v := range piles {
		right = max(right, v)
	}

	fmt.Printf("Performing binary search of speed k in range [%v; %v]. \n", left, right)

	minResult := right

	for left <= right {
		mid := (left + right) / 2

		hoursRequired := calculate(piles, mid, h)

		if hoursRequired > h { // not reached the target hours -> we need to increase the speed K
			left = mid + 1
		} else { // reached the target hours (hoursRequired <= targetHours) -> let's try to decrease K even more
			minResult = min(minResult, mid)
			right = mid - 1
		}
	}

	return minResult
}

func calculate(piles []int, k int, breakOnLimit int) int {
	totalHours := 0

	for _, v := range piles {
		// todo: we can also break earlier if sum overrides the limit
		totalHours += ceil(v, k)

		// we are already at overflow -> no need to calculate further
		if totalHours > breakOnLimit {
			return totalHours
		}
	}

	return totalHours
}

func ceil(n int, divisor int) int {
	// this is a heuristic version of ceil

	// example validation:
	// 6 / 3 -> (6 + 3 - 1) / 3 = 8 / 3 = 2
	// 7 / 3 -> (7 + 3 - 1) / 3 = 9 / 3 = 3
	// 8 / 3 -> (8 + 3 - 1) / 3 = 10 / 3 = 3
	// 9 / 3 -> (9 + 3 - 1) / 3 = 11 / 3 = 3
	// 10 / 3 -> (10 + 3 - 1) / 3 = 12 / 3 = 4

	return (n + divisor - 1) / divisor
}

func ceilNaive(n int, divisor int) int {
	if n%divisor == 0 {
		return n / divisor
	}

	return (n / divisor) + 1
}

func test(piles []int, h int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Piles: %v \n", piles)
	fmt.Printf("Hours: %v \n", h)

	minimumSpeed := minEatingSpeed(piles, h)

	fmt.Printf("Minimum required speed: %v \n", minimumSpeed)

	if minimumSpeed != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, minimumSpeed)
	}
}

func test1() {
	piles := []int{3, 6, 7, 11}
	h := 8
	expected := 4

	test(piles, h, expected)
}

func test2() {
	piles := []int{30, 11, 23, 4, 20}
	h := 5
	expected := 30

	test(piles, h, expected)
}

func test3() {
	piles := []int{30, 11, 23, 4, 20}
	h := 6
	expected := 23

	test(piles, h, expected)
}

func main() {
	// 875. Koko Eating Bananas
	test1()
	test2()
	test3()
}
