package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 { // corner-case, but possible
		return true
	}

	plantedCount := 0
	prev := 0
	next := 0

	for i, v := range flowerbed {
		if i == 0 {
			// there is no flower at the left of the whole array
			prev = 0
		} else {
			prev = flowerbed[i-1]
		}

		if i == (len(flowerbed) - 1) {
			// there is no flower at the right of the whole array
			next = 0
		} else {
			next = flowerbed[i+1]
		}

		// we can plant if previous, current and next are all 0
		canPlant := (prev + next + v) == 0
		if canPlant {
			flowerbed[i] = 1 // plant the flower at [i]
			plantedCount++

			if plantedCount >= n {
				return true
			}
		}
	}

	fmt.Printf("Cannot plant %v flowers. Just %v flowers planted after reaching the end of the flowerbed. \n", n, plantedCount)

	return false
}

func test(flowerbed []int, n int, expectedResult bool) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Flowerbed: %v \n", flowerbed)

	canPlace := canPlaceFlowers(flowerbed, n)

	fmt.Printf("Can place %v flowers: %v \n", n, canPlace)
	fmt.Printf("Flowerbed after planting %v flowers: %v \n", n, flowerbed)

	if canPlace != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, canPlace)
	}
}

func test1() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	expected := true

	test(flowerbed, n, expected)
}

func test2() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 2
	expected := false

	test(flowerbed, n, expected)
}

func test3() {
	flowerbed := []int{1, 0, 1}
	n := 0 // corner-case -> even if we can't plant anything, with n == 0 the result should be true
	expected := true

	test(flowerbed, n, expected)
}

func main() {
	// 605. Can Place Flowers
	test1()
	test2()
	test3()
}
