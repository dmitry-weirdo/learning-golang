package main

import (
	"fmt"
	"strings"
)

func reformatDate(date string) string {
	parts := strings.Split(date, " ")
	day := convertDay(parts[0])
	month := convertMonth(parts[1])
	year := parts[2]

	return fmt.Sprintf("%v-%v-%v", year, month, day)
}

func convertDay(s string) string {
	// remove 2 last characters "th", "st", "nd"
	s = s[:len(s)-2]

	// todo: we can use a format function to 2 digits
	if len(s) < 2 {
		return "0" + s
	}

	return s
}

func convertMonth(s string) string {
	switch s {
	case "Jan":
		return "01"
	case "Feb":
		return "02"
	case "Mar":
		return "03"
	case "Apr":
		return "04"
	case "May":
		return "05"
	case "Jun":
		return "06"
	case "Jul":
		return "07"
	case "Aug":
		return "08"
	case "Sep":
		return "09"
	case "Oct":
		return "10"
	case "Nov":
		return "11"
	case "Dec":
		return "12"
	default:
		panic(fmt.Sprintf("Unknown month value: \"%v\"", s))
	}
}

func test(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Date string: %v \n", s)

	result := reformatDate(s)

	fmt.Printf("Re-formatted data string: %v \n", result)
	fmt.Printf("Expected result:          %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("20th Oct 2052", "2052-10-20")
}

func test2() {
	test("6th Jun 1933", "1933-06-06")
}

func test3() {
	test("26th May 1960", "1960-05-26")
}

func main() {
	// 1507. Reformat Date
	test1()
	test2()
	test3()
}
