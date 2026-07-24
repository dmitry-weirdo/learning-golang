package main

import "fmt"

func isReflected(points [][]int) bool {
	// for fast lookup by x and y coordinates
	m := make(map[[2]int]bool)

	minX := points[0][0]
	maxX := points[0][0]

	for _, v := range points {
		minX = min(minX, v[0])
		maxX = max(maxX, v[0])

		// put to map for quick access
		key := [2]int{v[0], v[1]}
		m[key] = true
	}

	fmt.Printf("minX: %v, maxX: %v \n", minX, maxX)

	// reflection line is (minX + maxX) / 2 (can be 0.5)
	// we don't want to compare floats
	// let mid = (minX + maxX) / 2

	// If x is at left of mid,
	// Reflection is: (mid - x) + mid = 2 * mid - x = minX + maxX - x

	// If x is at right of mid,
	// Reflection is: mid - (x - mid) = 2 * mid - x = minX + maxX - x

	sum := minX + maxX

	for _, v := range points {
		reflectionX := sum - v[0]

		// reflectionX, same Y
		reflectionKey := [2]int{reflectionX, v[1]}

		if _, ok := m[reflectionKey]; !ok {
			fmt.Printf("Point [%v; %v] does not have a reflected point [%v; %v]. Returning false. \n", v[0], v[1], reflectionX, v[1])

			return false
		}
	}

	return true
}

func test(points [][]int, expectedResult bool) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Points: %v \n", points)

	result := isReflected(points)

	fmt.Printf("Points can be reflected: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	p := [][]int{
		{1, 1},
		{-1, 1},
	}

	expected := true

	test(p, expected)
}

func test2() {
	p := [][]int{
		{1, 1},
		{-1, -1},
	}

	expected := false

	test(p, expected)
}

func test3() {
	p := [][]int{
		{1, 0},
		{1, 0},
		{-1, 0}, // one match is enough for duplicate points
	}

	expected := true

	test(p, expected)
}

func main() {
	// 356. Line Reflection
	test1()
	test2()
	test3()
}
