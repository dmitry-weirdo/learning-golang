package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ipToCIDR(ip string, n int) []string {
	ipInt := ipToInt(ip)
	fmt.Printf("IP %v converted to int: %v \n", ip, ipInt)

	result := make([]string, 0)

	for n > 0 {
		// Determine the maximum size of the block with proper alignment.
		// We take the last significant bit
		lastBit := lsb(ipInt)

		// 0 is aligned to the entire IPv4 address space. To not fail on "0.0.0.0" case
		if lastBit == 0 {
			lastBit = 1 << 32
		}

		// If lastBit is > current value of n, we decrease it
		for lastBit > n {
			lastBit >>= 1
		}

		prefix := 32
		if lastBit > 0 {
			prefix = 32 - log2(lastBit) // mask of /32, /29 etc. We subtract the current lastBit power of 2 from 32
		}

		ipString := ipToString(ipInt) // all 4 octets of the current ip value as string

		result = append(result, fmt.Sprintf("%v/%v", ipString, prefix))

		ipInt += lastBit // move up in the ip list
		n -= lastBit     // decrease the remaining values
	}

	return result
}

func ipToInt(ip string) int { // todo: maybe int64
	split := strings.Split(ip, ".")
	//fmt.Printf("Ip split: %v \n", split)

	result := 0

	for _, v := range split {
		// each octet is a power of 256:
		// 256^3 * a[0] + 256^2 * a[1] + 256^1 * a[2] + 256^0 * a[3]
		octetAsInt, _ := strconv.Atoi(v)
		result = 256*result + octetAsInt
	}

	return result
}

func ipToString(ip int) string {
	return fmt.Sprintf(
		"%v.%v.%v.%v",
		ip>>24&255,
		ip>>16&255,
		ip>>8&255,
		ip>>0&255,
	)
}

func lsb(x int) int { // least significant bit
	return x & -x
}

func log2(n int) int {
	// todo: log2(0) should be handled separately, it's undefined
	return int(math.Log2(float64(n)))

	// for positive integers, counting bits can be used:
	// bits.Len(uint(n)) - 1
}

func test(ip string, n int, expectedResult []string) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("IP string: %v \n", ip)
	fmt.Printf("N addresses: %v \n", n)

	result := ipToCIDR(ip, n)

	fmt.Printf("Minimum required IP masks: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	test(
		"255.0.0.7",
		10,
		[]string{"255.0.0.7/32", "255.0.0.8/29", "255.0.0.16/32"},
	)
}

func test2() {
	test(
		"117.145.102.62",
		8,
		[]string{"117.145.102.62/31", "117.145.102.64/30", "117.145.102.68/31"},
	)
}

func test3() {
	// test-case 101 / 106
	test(
		"0.0.0.0",
		1,
		[]string{"0.0.0.0/32"},
	)
}

func test4() {
	// test-case 102 / 106
	test(
		"0.0.0.0",
		2,
		[]string{"0.0.0.0/31"},
	)
}

func main() {
	// 751. IP to CIDR
	test1()
	test2()
	test3()
	test4()
}
