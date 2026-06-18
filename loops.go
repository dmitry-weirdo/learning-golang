package main

import "fmt"

func main() {
	infiniteLoop()
	loopTillCondition()
	counterBasedLoop()
}

func infiniteLoop() {
	fmt.Println("======= infiniteLoop ======= ")

	i := 1

	for { // infinite loop - no condition
		fmt.Println(i)
		i++

		if i > 666 {
			break
		}
	}
}

func loopTillCondition() {
	fmt.Println("======= loopTillCondition ======= ")

	i := 1

	for i < 3 { // parentheses not mandatory
		fmt.Println(i)
		i++
	}

	fmt.Println("Done!")
}

func counterBasedLoop() {
	fmt.Println("======= counterBasedLoop ======= ")

	for i := 1; i < 4; i++ { // parentheses not mandatory
		fmt.Println(i)
	}

	fmt.Println("Done!")
}
