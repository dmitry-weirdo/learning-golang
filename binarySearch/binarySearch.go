package main

import "fmt"

func leftMostBinarySearch(a []int, v int) int { // returns index
	left := 0
	right := len(a) // todo: what if len - 1?

	for left < right {
		mid := (left + right) / 2

		if a[mid] < v {
			left = mid + 1
		} else { // if a[i] == mid, we will move left -> to the leftmost value
			right = mid
		}
	}

	return left
}

func rightMostBinarySearch(a []int, v int) int { // returns index
	left := 0
	right := len(a) // todo: what if len - 1?

	for left < right {
		mid := (left + right) / 2

		//fmt.Printf("left: %v, right: %v, mid: %v, a[mid] = %v \n", left, right, mid, a[mid])

		if a[mid] <= v {
			left = mid + 1 // this will jump to the next value after v
		} else {
			right = mid
		}
	}

	// we're now at the left-most of the NEXT value
	// if the previous element is the target -> return it
	previousIndex := left - 1

	if (previousIndex > 0) && (a[previousIndex] == v) {
		return previousIndex
	}

	return left
}

func rightMostBinarySearch2(a []int, v int) int { // returns index
	left := 0
	right := len(a) // todo: what if len - 1?

	result := -1

	for left < right {
		mid := (left + right) / 2

		//fmt.Printf("left: %v, right: %v, mid: %v, a[mid] = %v \n", left, right, mid, a[mid])
		if a[mid] == v {
			result = mid // remember the latest position, it will be the right-most

			left = mid + 1 // this will jump to the next value after v
		} else if a[mid] < v {
			left = mid + 1 // this will jump to the next value after v
		} else {
			right = mid
		}
	}

	return result
}

func test(a []int, v int) {
	fmt.Println()
	fmt.Println("==========================")

	fmt.Printf("Array: \n%v \n", a)

	leftOne := leftMostBinarySearch(a, v)
	rightOne := rightMostBinarySearch(a, v)
	rightOne2 := rightMostBinarySearch2(a, v)

	fmt.Printf("Left-most position of %v = [%v] \n", v, leftOne)
	fmt.Printf("(Alg. 1) Right-most position of %v = [%v] \n", v, rightOne)
	fmt.Printf("(Alg. 2) Right-most position of %v = [%v] \n", v, rightOne2)
}

func test1() {
	arr := []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 4}

	test(arr, 1)
	test(arr, 2)
	test(arr, 3)
	test(arr, 4)
	test(arr, 5)
}

func main() {
	test1()
}
