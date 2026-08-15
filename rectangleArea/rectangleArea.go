package main

import "fmt"

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	r1 := createRectangle(Point{ax1, ay1}, Point{ax2, ay2})
	r2 := createRectangle(Point{bx1, by1}, Point{bx2, by2})

	intersectionArea := getIntersectionArea(r1, r2)

	//fmt.Printf("Intersection area of 2 rectangles: %v \n", intersectionArea)

	return r1.square() + r2.square() - intersectionArea
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

func (r Rectangle) square() int {
	return (r.right - r.left) * (r.top - r.bottom)
}

func getIntersectionArea(r1 Rectangle, r2 Rectangle) int {
	// make r1 be left or equal to r2
	r1, r2 = getLeftRightRectangles(r1, r2)

	fmt.Printf("Rectangle 1 (more left): %v \n", r1)
	fmt.Printf("Rectangle 2 (more right): %v \n", r2)

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

func test(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Rectangle 1: bottom-left: (%v; %v), top-right: (%v, %v). \n", ax1, ay1, ax2, ay2)
	fmt.Printf("Rectangle 2: top-right: (%v; %v), top-right: (%v, %v). \n", bx1, by1, bx2, by2)

	result := computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)

	fmt.Printf("Area covered by both Rectangle 1 and Rectangle 2: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	ax1, ay1 := -3, 0
	ax2, ay2 := 3, 4
	bx1, by1 := 0, -1
	bx2, by2 := 9, 2

	expected := 45

	test(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2, expected)
}

func test2() {
	// same 4x4 squares
	ax1, ay1 := -2, -2
	ax2, ay2 := 2, 2
	bx1, by1 := -2, -2
	bx2, by2 := 2, 2

	expected := 16

	test(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2, expected)
}

func main() {
	// 223. Rectangle Area
	test1()
	test2()
}
