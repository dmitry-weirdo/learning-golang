package main

import (
	"fmt"
)

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	xDist := distToRange(x1, x2, xCenter)
	yDist := distToRange(y1, y2, yCenter)

	return xDist*xDist+yDist*yDist <= radius*radius
}

func distToRange(rangeStart, rangeEnd, point int) int {
	if rangeStart <= point && point <= rangeEnd {
		return 0
	}

	if point < rangeStart {
		return rangeStart - point
	} else {
		return point - rangeEnd
	}
}

func test(radius, x, y int, x1, y1, x2, y2 int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Circle radius: %v \n", radius)
	fmt.Printf("Circle center: [%v; %v] \n", x, y)
	fmt.Printf("Rectangle bottom-left: [%v; %v] \n", x1, y1)
	fmt.Printf("Rectangle top-right: [%v; %v] \n", x2, y2)

	result := checkOverlap(radius, x, y, x1, y1, x2, y2)

	fmt.Printf("Circle and rectangle overlap: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		1, 0, 0,
		1, -1, 3, 1,
		true,
	)
}

func test2() {
	test(
		1, 1, 1,
		1, -3, 2, -1,
		false,
	)
}

func test3() {
	test(
		1, 0, 0,
		-1, 0, 0, 1,
		true,
	)
}

func main() {
	// 1401. Circle and Rectangle Overlapping
	test1()
	test2()
	test3()
}
