package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	// we check that any point compared to p[0] has the same tg as p[0] and p[1]
	// dx01 / dy01 = dx0i / dy0i
	// dx01 * dy0i = dx0i * dy01

	if len(coordinates) == 2 { // there is a line through any 2 points, and there are at least 2 points in the array
		return true
	}

	dx01, dy01 := getDeltas(coordinates[0], coordinates[1])

	for i := 2; i < len(coordinates); i++ {
		dx0i, dy0i := getDeltas(coordinates[0], coordinates[i])

		if dx01*dy0i != dx0i*dy01 {
			return false
		}
	}

	return true
}

func getDeltas(p1, p2 []int) (deltaX, deltaY int) {
	deltaX = p2[0] - p1[0]
	deltaY = p2[1] - p1[1]

	return deltaX, deltaY
}

func test(m [][]int, expectedResult bool) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Points coordinates: %v \n", m)

	result := checkStraightLine(m)

	fmt.Printf("All points are in one straight line: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[][]int{
			{1, 2},
			{2, 3},
			{3, 4},
			{4, 5},
			{5, 6},
			{6, 7},
		},
		true,
	)
}

func test2() {
	test(
		[][]int{
			{1, 1},
			{2, 2},
			{3, 4},
			{4, 5},
			{5, 6},
			{7, 7},
		},
		false,
	)
}

func main() {
	// 1232. Check If It Is a Straight Line
	test1()
	test2()
}
