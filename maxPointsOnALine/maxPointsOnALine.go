package main

import (
	"fmt"
	"math"
)

func maxPoints(points [][]int) int {
	// For every point, we calculate atan2 of this point and every other point.
	// We count the count of every atan2, the biggest one is probably updating the max.

	n := len(points)

	if n <= 1 {
		return n
	}

	maxPointsOnOneLine := 2 // every 2 points form a line

	for i, point1 := range points {
		m := make(map[float64]int) // atan2 -> count of points[j] that make this atan2 with points[i]

		//for j := i + 1; j < n; j++ { // no, we cannot continue for the next point, we need to check all vectors with all vectors
		for j, point2 := range points {
			if i == j { // do not calculate fot the point itself
				continue
			}

			//point2 := points[j]

			x1 := point1[0]
			y1 := point1[1]

			x2 := point2[0]
			y2 := point2[1]

			// we count atan2 for a vector point1 -> point2
			atan2 := math.Atan2(float64(y2-y1), float64(x2-x1))

			m[atan2]++
		}

		maxCountForPoint1 := 0

		for _, v := range m {
			maxCountForPoint1 = max(maxCountForPoint1, v)
		}

		maxCountForPoint1++ // we also count the points[i] itself!

		maxPointsOnOneLine = max(maxPointsOnOneLine, maxCountForPoint1)
	}

	return maxPointsOnOneLine
}

func testAtan2() {
	// atan2 is an angle between (0,0) and given (x, y)
	// arguments of atan2 are (y, x), NOT (x, y)

	// atan2 gives the quadrant values (unlike atan),
	// i.e. values range between [0; 180] and [0; -180] degrees in radian.

	x := float64(1)
	y := float64(0)
	atan2 := math.Atan2(y, x) // 0 = 0 grad
	fmt.Printf("x: %v, y: %v, atan2(y, x) = %v \n", x, y, atan2)

	x = float64(1)
	y = float64(1)
	atan2 = math.Atan2(y, x) // pi / 4 = 45 grad
	fmt.Printf("x: %v, y: %v, atan2(y, x) = %v \n", x, y, atan2)

	x = float64(-1)
	y = float64(1)
	atan2 = math.Atan2(y, x) // pi * 3 / 4 = 135 grad
	fmt.Printf("x: %v, y: %v, atan2(y, x) = %v \n", x, y, atan2)

	x = float64(-1)
	y = float64(0)
	atan2 = math.Atan2(y, x) // pi = 180 grad
	fmt.Printf("x: %v, y: %v, atan2(y, x) = %v \n", x, y, atan2)

	x = float64(0)
	y = float64(1)
	atan2 = math.Atan2(y, x) // pi / 2 = 90 grad
	fmt.Printf("x: %v, y: %v, atan2(y, x) = %v \n", x, y, atan2)

	x = float64(0)
	y = float64(-1)
	atan2 = math.Atan2(y, x) // -pi / 2 (same as pi * 3 / 2) = -90 grad
	fmt.Printf("x: %v, y: %v, atan2(y, x) = %v \n", x, y, atan2)
}

func test(m [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Points: %v \n", m)

	result := maxPoints(m)

	fmt.Printf("Max points in one line: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	points := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}

	expected := 3

	test(points, expected)
}

func test2() {
	points := [][]int{
		{1, 1},
		{3, 2},
		{5, 3},
		{4, 1},
		{2, 3},
		{1, 4},
	}

	expected := 4

	test(points, expected)
}

func main() {
	// 149. Max Points on a Line
	test1()
	test2()
}
