package main

import "fmt"

func minDays(bloomDay []int, requiredBouquetsCount int, flowersInBouquet int) int {
	n := len(bloomDay)

	if requiredBouquetsCount*flowersInBouquet > n { // required flowers > flowers in the garden -> no solution
		return -1
	}

	maxDays := bloomDay[0]

	for _, v := range bloomDay {
		maxDays = max(maxDays, v)
	}

	//fmt.Printf("Max days to bloom: %v \n", maxDays)

	// binary search by count of days
	// logic is similar to "1760. Minimum Limit of Balls in a Bag"
	left := 1
	right := maxDays // we should find the solution since (m * k <= n)

	// we search for minimum day where getBouquetsCount(day) >= k
	for left < right {
		midDay := (left + right) / 2

		// small optimization -> in the last day, the complete array is blooming, there is no need to check it.
		if (midDay >= maxDays) || getBouquetsCount(bloomDay, midDay, flowersInBouquet, requiredBouquetsCount) >= requiredBouquetsCount { // target condition
			right = midDay // in this template it is always mid, NOT mid - 1
		} else {
			left = midDay + 1
		}
	}

	return left
}

func getBouquetsCount(a []int, day int, flowerInBouquet int, requiredBouquetsCount int) int {
	totalBouquets := 0

	count := 0 // current consecutive flowers after cutting the previous bouquet

	for _, v := range a {
		if day >= v { // current flower blooms in given day
			count++
		} else { // current flower does not bloom in given day
			count = 0
		}

		if count >= flowerInBouquet { // cut a bouquet
			totalBouquets++
			count = 0

			if totalBouquets >= requiredBouquetsCount { // enough bouquets already -> stop the iteration
				return totalBouquets
			}
		}
	}

	return totalBouquets
}

func test(arr []int, bouquetsCount int, flowerInBouquet int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Flower bloom days: %v \n", arr)
	fmt.Printf("M - Required bouquets count: %v \n", bouquetsCount)
	fmt.Printf("K - Adjacent flowers in bouquet: %v \n", flowerInBouquet)

	result := minDays(arr, bouquetsCount, flowerInBouquet)

	fmt.Printf("Minimum number of days required: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{1, 10, 3, 10, 2}
	m := 3
	k := 1
	expected := 3

	test(arr, m, k, expected)
}

func test2() {
	arr := []int{1, 10, 3, 10, 2}
	m := 3
	k := 2
	expected := -1 // m * k > n

	test(arr, m, k, expected)
}

func test3() {
	arr := []int{7, 7, 7, 7, 12, 7, 7}
	m := 2
	k := 3
	expected := 12 // can collect after the max days

	test(arr, m, k, expected)
}

func main() {
	// 1482. Minimum Number of Days to Make m Bouquets
	// Similar binary search of result + check feasibility as "1760. Minimum Limit of Balls in a Bag"
	test1()
	test2()
	test3()
}
