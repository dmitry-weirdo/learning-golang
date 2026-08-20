package main

import (
	"container/list"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func sliceFunctions() {
	fmt.Println()
	fmt.Println("==================== slices/arrays functions ====================")

	// remove a central element from the slice
	a := []int{1, 2, 3, 4}

	fmt.Printf("Slice before removal: %v \n", a)

	indexToRemove := 1
	removedElement := a[indexToRemove]
	a = append(a[:indexToRemove], a[indexToRemove+1:]...)

	fmt.Printf("Element removed from a[%v] = %v \n", indexToRemove, removedElement)
	fmt.Printf("Slice after removing a[%v]: %v \n", indexToRemove, a)

	// remove last element
	a = a[:len(a)-1]
	fmt.Printf("Slice after removing last element: %v \n", a)

	// call the pointer operations
	b := &a
	slicePointerFunction(b)

	fmt.Printf("Slice after performing operations by pointer: %v \n", b)

	// reverse works in place
	slices.Reverse(a)
	fmt.Printf("Slice reversed in place: %v \n", a)

	// we can reverse just a part of the array
	slices.Reverse(a[1:3]) // end is non-inclusive
	fmt.Printf("Slice part [1:2] reversed in place: %v \n", a)

	startIndexInclusive := 1
	endIndexNonInclusive := 3
	sliceWithDeleteFromMiddle := slices.Delete(a, startIndexInclusive, endIndexNonInclusive) // in-places is also removed, but length remains, end elements will be set to 0

	fmt.Printf("Slice after slices.Delete[%v; %v) is modified, but the length remains: %v \n", startIndexInclusive, endIndexNonInclusive, a)
	fmt.Printf("Separate modified slice after slices.Delete[%v; %v): %v \n", startIndexInclusive, endIndexNonInclusive, sliceWithDeleteFromMiddle)

	// check if slice contains a value
	contains666 := slices.Contains(sliceWithDeleteFromMiddle, 666)
	contains100 := slices.Contains(sliceWithDeleteFromMiddle, 100)
	fmt.Printf("Slice %v contains 666: %v \n", sliceWithDeleteFromMiddle, contains666)
	fmt.Printf("Slice %v contains 100: %v \n", sliceWithDeleteFromMiddle, contains100)
}

func slicePointerFunction(a *[]int) { // passing a pointer will modify the argument
	// replace element by pointer
	(*a)[0] = 666

	// append to pointer to slice
	*a = append(*a, 100, 200, 300)

	// remove last from pointer to slice
	*a = (*a)[:len(*a)-1]
}

func copyArray(arr []int) []int {
	arrayCopy := make([]int, len(arr))
	copy(arrayCopy, arr)
	return arrayCopy
}

func maxInArray(arr []int) int {
	// we assume the array is non-empty
	m := arr[0]

	for _, v := range arr {
		m = max(m, v)
	}

	return m
}

func listFunctions() {
	fmt.Println()
	fmt.Println("==================== list.List functions ====================")

	var l = list.New() // list is NOT type, we can push any values

	// push to back
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	// push to front
	l.PushFront(0)
	l.PushFront(-1)

	fmt.Printf("list.List: %v \n", l)

	// remove from front
	fromFront := l.Remove(l.Front()).(int) // cast to your type
	fromBack := l.Remove(l.Back()).(int)   // cast to your type

	fmt.Printf("Element removed from list front: %v \n", fromFront)
	fmt.Printf("Element removed from list back: %v \n", fromBack)

	// iterate list forwards
	fmt.Println("Iterating list forwards: ")

	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(int)
		fmt.Printf("%v ", v)
	}

	fmt.Println()

	// iterate list backwards
	fmt.Println("Iterating list backwards: ")

	for e := l.Back(); e != nil; e = e.Prev() {
		v := e.Value.(int)
		fmt.Printf("%v ", v)
	}

	fmt.Println()
}

func stringBuilder() {
	fmt.Println()
	fmt.Println("==================== StringBuilder functions ====================")

	var r rune = 'a'
	var ch byte = 'b' // 1-byte character is of type byte

	var intValue = 666

	var sb = strings.Builder{}

	sb.WriteRune(r)
	sb.WriteByte(ch)                       // one-byte character
	sb.WriteString(strconv.Itoa(intValue)) // there is no WriteInt

	s := sb.String() // toString

	fmt.Printf("s from StringBuilder: %v \n", s)

	// flush/clean the StringBuilder
	sb.Reset()
	fmt.Printf("StringBuilder after reset: %v \n", sb)
	fmt.Printf("sb.String() after reset: %v \n", sb.String())
}

func stringToIntArray(s string) []int { // every byte converted to int
	a := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		a[i] = int(s[i])
	}

	return a
}

func reverseString(s string) string {
	stringAsSlice := []byte(s)
	slices.Reverse(stringAsSlice) // reverses in place

	return string(stringAsSlice)
}

func getLongAndShortStrings(a, b string) (long, short string) {
	if len(a) > len(b) {
		return a, b
	}

	return b, a
}

func getMinAndMax(a, b int) (smaller, greater int) {
	if a <= b {
		return a, b
	}

	return b, a
}

func uniquePrimeFactors(n int) []int { // gets unique prime factors
	// todo: handle n == 1

	// todo: we can use precalculated prime numbers to only iterate only via them
	factors := make([]int, 0)

	// todo: fix the potential overflow on (d * d)
	for d := 2; d*d <= n; d++ { // iterate up to ceil(sqrt(n))
		if n%d == 0 {
			factors = append(factors, d) // only append this prime factor once

			// Remove all occurrences of d, so that all non-prime will drop as well (dividing by 2 will remove 4, 6 etc).
			for n%d == 0 {
				n /= d
			}
		}
	}

	// if the remaining value is prime (not 1 after all the divisions) -> add it as well
	// e.g. 7 -> 7 will remain
	// 26 -> 13 will remain
	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

func uniquePrimeFactorsWithPrecalculatedPrimeNumbers(n int, primeNumbers []int) []int { // gets unique prime factors
	// todo: handle n == 1

	factors := make([]int, 0)

	for _, d := range primeNumbers { // we only iterate over the prime numbers
		// todo: fix the potential overflow on (d * d)
		if d*d > n {
			break
		}

		if n%d == 0 {
			factors = append(factors, d) // only append this prime factor once

			// Remove all occurrences of d, so that all non-prime will drop as well (dividing by 2 will remove 4, 6 etc).
			for n%d == 0 {
				n /= d
			}
		}
	}

	// if the remaining value is prime (not 1 after all the divisions) -> add it as well
	// e.g. 7 -> 7 will remain
	// 26 -> 13 will remain
	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

func primesUpTo(n int) []int {
	// Sieve of Eratosthenes
	if n < 2 {
		return []int{}
	}

	// to not init all to true in a separate O(n) cycle, we use the negation
	isNotPrime := make([]bool, n+1)

	// 1 is NOT prime
	isNotPrime[1] = true

	// Remove multiples of each prime.
	for p := 2; p*p <= n; p++ {
		if !isNotPrime[p] {
			for multiple := p * p; multiple <= n; multiple += p {
				isNotPrime[multiple] = true
			}
		}
	}

	// Collect the primes.
	primes := []int{}
	for i := 2; i <= n; i++ {
		if !isNotPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func lcm(a, b int) int { // least common multiple
	gcdOfAAndB := gcd(a, b)

	return (a / gcdOfAAndB) * b
}

func gcd(a, b int) int { // greatest common divisor
	// Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func main() {
	stringBuilder()
	listFunctions()
	sliceFunctions()
}
