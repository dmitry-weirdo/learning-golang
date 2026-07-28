package main

import (
	"fmt"
	"math"
)

const (
	MINUS      = '-'
	PLUS       = '+'
	ZERO       = '0'
	ONE        = '1'
	NINE       = '9'
	WHITESPACE = ' '
)

func myAtoi(s string) int {
	if len(s) < 1 {
		return 0
	}

	i := 0

	// skip leading whitespaces
	for (i < len(s)) && s[i] == WHITESPACE {
		i++
	}

	if i >= len(s) {
		fmt.Printf("String \"%v\" contains only whitespaces. Returning 0.\n", s)
		return 0
	}

	fmt.Printf("Skipped %v whitespace characters. First non-whitespace index: %v, First non-whitespace character: %c \n", i, i, s[i])

	// read minus or plus
	var sign int32

	if s[i] == MINUS { // skip explicit - sign
		sign = -1
		i++
	} else if s[i] == PLUS { // skip explicit + sign
		sign = 1
		i++
	} else { // don't skip the current character, number is positive by default
		sign = 1
	}

	if i >= len(s) {
		fmt.Printf("String \"%v\" contains only whitespaces and sign. Returning 0.\n", s)
		return 0
	}

	fmt.Printf("Sign detected: %v \n", sign)
	fmt.Printf("Next index after sign: %v. Char: %c. \n", i, s[i])

	// skip leading zeroes
	for (i < len(s)) && s[i] == ZERO {
		i++
	}

	if i >= len(s) {
		fmt.Printf("String \"%v\" contains only zeroes. Returning 0.\n", s)
		return 0
	}

	fmt.Printf("Skipped leading zeroes after sign: %v \n", sign)
	fmt.Printf("Next index after leading zeroes: %v. Char: %c. \n", i, s[i])

	// if now it's not digit -> return 0
	if !isDigitFrom1to9(s[i]) {
		fmt.Printf("String \"%v\" does not start with a non-0 digit after leading zeroes. Returning 0.\n", s)

		return 0
	}

	// make it int32 to frankly work with the overflow
	var v int32 = 0

	// -2147483648; 2147483647
	const overflowLimit = math.MaxInt32 / 10 // without the last digit

	for (i < len(s)) && isDigit(s[i]) { // proceed from current index until it's digit from 0 to 9. We already skipped the leading zeroes
		// -2147483648; 2147483647

		nextDigit := byteToInt(s[i])
		fmt.Printf("Next digit: %v \n", nextDigit)

		// Handle potential overflow.
		// Overflow limit is 214748364
		// 2147483647 will be no overflow, but 2147483648 is already an overflow (will fail if positive)

		if (v > overflowLimit) || ((v == overflowLimit) && (nextDigit > 7)) { // we reached the limit -> return the max value
			if sign < 0 {
				fmt.Printf("String \"%v\" contains overflow and is negative. Returning %v.\n", s, math.MinInt32)

				return math.MinInt32
			} else {
				fmt.Printf("String \"%v\" contains overflow and is positive. Returning %v.\n", s, math.MaxInt32)

				return math.MaxInt32
			}
		}

		v = v*10 + nextDigit
		fmt.Printf("Updated value to %v. Hit overflow limit: %v \n", v, v >= overflowLimit)

		// go next
		i++
	}

	result := sign * v

	fmt.Printf("End of digits in \"%v\" end reached while handling the digits. Returning %v.\n", s, result)
	return int(result)
}

func byteToInt(b byte) int32 {
	return int32(b - ZERO)
}

func isDigitFrom1to9(b byte) bool {
	return b >= ONE && b <= NINE
}

func isDigit(b byte) bool {
	return b >= ZERO && b <= NINE
}

func convert(s string) int {
	fmt.Println()
	fmt.Println("==========================")
	fmt.Printf("Handling string \"%v\"... \n", s)

	result := myAtoi(s)

	fmt.Printf("String \"%v\" converted to int32 = %v \n", s, result)

	return result
}

func main() {
	// 8. String to Integer (atoi)

	// int16 = [-32768; 32767]
	// int32 [-2147483648; 2147483647]
	// int is signed int64 [-9223372036854775808; 9223372036854775807]

	fmt.Printf("Min int: %v, max int : %v \n", math.MinInt, math.MaxInt)

	convert("     ")
	convert("")
	convert("   42")
	convert("42")
	convert("  -042")
	convert("-042")
	convert("   -042")
	convert("+042")
	convert("1337c0d3")
	convert("0-1")
	convert("-")
	convert("-00000")
	convert("   -   ")
	convert("   - 32")
	convert("words and 987")
	convert("9999999999999999999999999999999999")
	convert("-9999999999999999999999999999999999")
	convert("-91283472332")
	convert("2147483647")
	convert("2147483648")
	convert("-2147483647")
	convert("-2147483648")

	// 1089159116

	// 538412641
	// 912834723
	// 214748364
}
