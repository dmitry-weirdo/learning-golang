package main

import (
	"fmt"
	"strconv"
)

func similarRGB(color string) string {
	v := []int{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}

	byte0 := hexStringToInt(color[1:3])
	byte1 := hexStringToInt(color[3:5])
	byte2 := hexStringToInt(color[5:7])

	//fmt.Printf("Bytes parsed: %v %v %v \n", byte0, byte1, byte2)

	r0 := getNearest(byte0, v)
	r1 := getNearest(byte1, v)
	r2 := getNearest(byte2, v)

	return "#" + intTo2HexDigits(r0) + intTo2HexDigits(r1) + intTo2HexDigits(r2)
}

func hexStringToInt(s string) int {
	v, _ := strconv.ParseInt(s, 16, 32) // use 64 if you need int64
	return int(v)
}

func intTo2HexDigits(v int) string {
	return fmt.Sprintf("%02x", v) // x is lower-case, X is upper-case
}

func getNearest(b int, a []int) int {
	if b == 0xFF {
		return 0xFF
	}

	for i, v := range a {
		if b == v {
			return v
		}

		next := a[i+1]
		if (b > v) && (b < next) {
			if (b - v) < (next - b) {
				return v
			}

			return next
		}
	}

	return a[len(a)-1]
}

func test(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Color string: %v \n", s)

	result := similarRGB(s)

	fmt.Printf("Nearest color that can be written as #xxx: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("#09f166", "#11ee66")
}

func test2() {
	test("#4e3fe1", "#5544dd")
}

func main() {
	// 800. Similar RGB Color
	test1()
	test2()
}
