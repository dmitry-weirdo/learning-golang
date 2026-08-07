package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	student0 := 0
	student1 := 0

	for _, v := range students {
		if v == 0 {
			student0++
		} else {
			student1++
		}
	}

	//fmt.Printf("Students 0 total: %v \n", student0)
	//fmt.Printf("Students 1 total: %v \n", student1)
	for _, v := range sandwiches {
		if v == 0 {
			if student0 > 0 {
				student0--
			} else { // blocked on sandwich 0, only students 1 left
				return student1
			}
		} else {
			if student1 > 0 {
				student1--
			} else { // blocked on sandwich 1, only students 0 left
				return student0
			}
		}
	}

	// end reached -> no student left
	return 0
}

func test(students, sandwiches []int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Students: %v \n", students)
	fmt.Printf("Sandwiches: %v \n", sandwiches)

	result := countStudents(students, sandwiches)

	fmt.Printf("Number of students unable to eat %v: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}
	expected := 0

	test(students, sandwiches, expected)
}

func test2() {
	students := []int{1, 1, 1, 0, 0, 1}
	sandwiches := []int{1, 0, 0, 0, 1, 1}
	expected := 3

	test(students, sandwiches, expected)
}

func main() {
	// 1700. Number of Students Unable to Eat Lunch
	test1()
	test2()
}
