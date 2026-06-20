package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	var minusSign string
	if sig(numerator)*sig(denominator) < 0 {
		minusSign = "-"
	} else {
		minusSign = ""
	}

	//fmt.Printf("Minus sign: %v \n", minusSign)

	remainder := numerator % denominator
	if remainder == 0 { // integer division!
		return strconv.Itoa(numerator / denominator)
	}

	var n = abs(numerator)
	var d = abs(denominator)

	integerPart := n / d
	remainder = n % d

	// remainder to position
	var m = make(map[int]int)
	remainderFoundInMap := false

	var sb strings.Builder
	sb.WriteString(minusSign)
	sb.WriteString(strconv.Itoa(integerPart))
	sb.WriteString(".")

	var pos = sb.Len()

	for !remainderFoundInMap {
		// if newRemainder == 0 -> no brackets, it's a non-periodic division
		if remainder == 0 {
			fmt.Printf("Divided with 0 remainder! Returning non-periodic result %v. \n", sb.String())

			return sb.String()
		}

		remainderPos, ok := m[remainder]
		if ok {
			remainderFoundInMap = true

			fmt.Printf("Remainder %v already found in position %v \n", remainder, remainderPos)

			// remainder cycle found -> append the closing bracket
			sb.WriteString(")")

			withoutOpeningBracket := sb.String()

			return withoutOpeningBracket[:remainderPos] + "(" + withoutOpeningBracket[remainderPos:]
		} else { // handle the new reminder
			m[remainder] = pos
			pos++

			fmt.Printf("map updated with new remainder %v. Map: %v \n", remainder, m)
		}

		// handle the new cycle
		newNumerator := remainder * 10
		newDivision := newNumerator / d
		newRemainder := newNumerator % d

		fmt.Printf("Next division: %v / %v = (%v, %v) \n", newNumerator, d, newDivision, newRemainder)

		// in any case, add the division to the result
		sb.WriteString(strconv.Itoa(newDivision))
		fmt.Printf("Added division %v to the result: %v \n", newDivision, sb.String())

		// update the values for the new cycle
		remainder = newRemainder
	}

	// this must never happen
	return sb.String()
}

func abs(i int) int {
	if i >= 0 {
		return i
	}

	return -i
}

func sig(i int) int {
	if i < 0 {
		return -1
	} else if i > 0 {
		return 1
	}

	return 0
}

func main() {
	var n, d int

	n, d = 1, 2
	fmt.Printf("%v / %v = %v \n", n, d, fractionToDecimal(n, d))

	n, d = 2, 1
	fmt.Printf("%v / %v = %v \n", n, d, fractionToDecimal(n, d))

	n, d = 4, 333
	fmt.Printf("%v / %v = %v \n", n, d, fractionToDecimal(n, d))

	n, d = 1, 6
	fmt.Printf("%v / %v = %v \n", n, d, fractionToDecimal(n, d))
}
