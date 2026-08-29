package main

import "fmt"

func judgeCircle(moves string) bool {
	x := 0
	y := 0

	for _, c := range moves { // first value is index, second value is char
		switch c {
		case 'U':
			y++
		case 'D':
			y--

		case 'L':
			x--
		case 'R':
			x++
		}
	}

	return (x == 0) && (y == 0)
}

func test(s string, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Moves: %v \n", s)

	result := judgeCircle(s)

	fmt.Printf("Robot ends at [0; 0]: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("UD", true)
}

func test2() {
	test("LL", false)
}

func main() {
	// 657. Robot Return to Origin
	test1()
	test2()
}
