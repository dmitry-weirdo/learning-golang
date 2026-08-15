package main

import (
	"fmt"
)

func largestTriangleArea(points [][]int) float64 {
	maxArea := float64(0)

	n := len(points)

	for i := 0; i < n-2; i++ { // we need 2 points ahead
		for j := i + 1; j < n-1; j++ { // we need 1 point ahead
			for k := j + 1; k < n; k++ {
				triangleArea := getTriangleAreaShoelace(
					points[i][0], points[i][1],
					points[j][0], points[j][1],
					points[k][0], points[k][1],
				)

				//fmt.Printf("Area of triangle %v, %v, %v: %v \n", points[i], points[j], points[k], triangleArea)

				maxArea = max(maxArea, triangleArea)
			}
		}
	}

	return maxArea
}

func getTriangleAreaShoelace(x1, y1, x2, y2, x3, y3 int) float64 {
	// Shoelace formula to calculate the area of a polygon.
	// For the triangle (3 points), ordering is not important since we'll run through the vertices either clockwise or counterclockwise
	return float64(abs(
		(x2-x1)*(y3-y1)-(y2-y1)*(x3-x1),
	)) / float64(2)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func test(m [][]int, expectedResult float64) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Point coordinates: %v \n", m)
	result := largestTriangleArea(m) // todo: replace with your function

	fmt.Printf("Largest area of a triangle: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	points := [][]int{
		{0, 0},
		{0, 1},
		{1, 0},
		{0, 2},
		{2, 0},
	}

	expected := float64(2)

	test(points, expected)
}

func test2() {
	points := [][]int{
		{1, 0},
		{0, 0},
		{0, 1},
	}

	expected := float64(0.5)

	test(points, expected)
}

func main() {
	// 812. Largest Triangle Area
	test1()
	test2()
}
