package main

import "fmt"

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	r1 := createRectangle(Point{rec1[0], rec1[1]}, Point{rec1[2], rec1[3]})
	r2 := createRectangle(Point{rec2[0], rec2[1]}, Point{rec2[2], rec2[3]})

	intersectionArea := getIntersectionArea(r1, r2)

	return intersectionArea > 0
}

type Point struct {
	x, y int
}

type Rectangle struct {
	topLeft, topRight, bottomLeft, bottomRight Point
	left, right, top, bottom                   int
}

func createRectangle(bottomLeft, topRight Point) Rectangle {
	left := bottomLeft.x
	right := topRight.x

	top := topRight.y
	bottom := bottomLeft.y

	return Rectangle{
		topLeft:     Point{left, top},
		topRight:    Point{right, top},
		bottomLeft:  Point{left, bottom},
		bottomRight: Point{right, bottom},
		left:        left,
		right:       right,
		top:         top,
		bottom:      bottom,
	}
}

func getIntersectionArea(r1 Rectangle, r2 Rectangle) int {
	// make r1 be left or equal to r2
	r1, r2 = getLeftRightRectangles(r1, r2)

	//fmt.Printf("Rectangle 1 (more left): %v \n", r1)
	//fmt.Printf("Rectangle 2 (more right): %v \n", r2)

	// get horizontal intersection
	if r2.left >= r1.right { // no horizontal intersection -> square of intersection is 0
		return 0
	}

	horizontalIntersection := 0

	if r2.right <= r1.right { // r2 horizontally within r1 -> take full width of r2
		horizontalIntersection = r2.right - r2.left
	} else { // r2.left within r1, r2.right out of r1 -> intersection is from r2.left to r1.right
		horizontalIntersection = r1.right - r2.left
	}

	// make r1 be top or equal to r2
	r1, r2 = getTopBottomRectangles(r1, r2)

	// get vertical intersection
	if r1.bottom >= r2.top { // no horizontal intersection -> square of intersection is 0
		return 0
	}

	verticalIntersection := 0

	if r2.bottom >= r1.bottom { // r2 vertically within r1 -> take full height of f2
		verticalIntersection = r2.top - r2.bottom
	} else { // r2.top within r1, r2.bottom out of r1 -> intersection from r2.top to r1.bottom
		verticalIntersection = r2.top - r1.bottom
	}

	return horizontalIntersection * verticalIntersection
}

func getLeftRightRectangles(r1, r2 Rectangle) (left, right Rectangle) {
	if r1.left < r2.left {
		return r1, r2
	}

	if r1.left > r2.left {
		return r2, r1
	}

	// r1.left == r2.left -> let's return the higher first
	if r1.top >= r2.top {
		return r1, r2
	}

	return r2, r1
}

func getTopBottomRectangles(r1, r2 Rectangle) (left, right Rectangle) {
	if r1.top > r2.top {
		return r1, r2
	}

	if r1.top < r2.top {
		return r2, r1
	}

	// r1.top == r2.top -> let's return the lefter first
	if r1.left <= r2.left {
		return r1, r2
	}

	return r2, r1
}

func test(r1, r2 []int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Rectangle 1: %v \n", r1)
	fmt.Printf("Rectangle 2: %v \n", r2)

	result := isRectangleOverlap(r1, r2)

	fmt.Printf("Rectangles overlap: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]int{0, 0, 2, 2},
		[]int{1, 1, 3, 3},
		true,
	)
}

func test2() {
	test(
		[]int{0, 0, 1, 1},
		[]int{1, 0, 2, 1},
		false,
	)
}

func test3() {
	test(
		[]int{0, 0, 1, 1},
		[]int{2, 2, 3, 3},
		false,
	)
}

func main() {
	// 836. Rectangle Overlap
	// This is a simplified version of "223. Rectangle Area".
	test1()
	test2()
	test3()
}
