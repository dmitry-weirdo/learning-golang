package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	writePos := m + n - 1
	left := m - 1  // index in left array
	right := n - 1 // index in right array

	// if we go from 0, we will need insertions that take a lot of time
	// so we go from right, filling in a1 from right to left

	// todo: we can actually just check right >= 0
	// if we reached the left of a2, all values from a1 are already in place -> nothing to do
	for writePos >= 0 && right >= 0 {
		fmt.Println()
		fmt.Printf("left: %v, right: %v \n", left, right)
		fmt.Printf("a1: %v \n", nums1)
		fmt.Printf("a2: %v \n", nums2)

		// reached the end of the left -> just copy from right
		if left < 0 {
			nums1[writePos] = nums2[right]

			right--
			writePos--

			continue
		}

		if nums1[left] >= nums2[right] { // copy from the left array
			nums1[writePos] = nums1[left]
			left--
			writePos--
			continue
		}

		// copy from the right array
		nums1[writePos] = nums2[right]
		right--
		writePos--
	}
}

func test(a1, a2 []int, m, n int) {
	fmt.Println()
	fmt.Println("=======================")

	fmt.Printf("a1: %v \n", a1)
	fmt.Printf("m: %v \n", m)

	fmt.Printf("a2: %v \n", a2)
	fmt.Printf("n: %v \n", n)

	merge(a1, m, a2, n)

	fmt.Printf("a1 after merge: %v \n", a1)
	fmt.Printf("a2 after merge: %v \n", a2)
}

func test1() {
	a1 := []int{1, 2, 3, 0, 0, 0}
	m := 3

	a2 := []int{2, 5, 6}
	n := 3

	test(a1, a2, m, n)
}

func test2() {
	a1 := []int{1}
	m := 1

	a2 := []int{}
	n := 0

	test(a1, a2, m, n)
}

func test3() {
	a1 := []int{0}
	m := 0

	a2 := []int{1}
	n := 1

	test(a1, a2, m, n)
}

func test4() {
	a1 := []int{4, 5, 6, 0, 0, 0}
	m := 3

	a2 := []int{1, 2, 3}
	n := 3

	test(a1, a2, m, n)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
