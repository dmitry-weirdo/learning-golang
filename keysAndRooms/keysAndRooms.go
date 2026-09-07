package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)

	visitedCount := 0
	visited := make([]bool, n)

	queue := make([]int, 0)
	queue = append(queue, 0)

	for len(queue) > 0 {
		currentLevelElementsCount := len(queue)

		for range currentLevelElementsCount {
			// poll from queue
			room := queue[0]
			queue = queue[1:]

			if !visited[room] {
				visited[room] = true
				visitedCount++

				if visitedCount >= n {
					return true
				}
			}

			for _, v := range rooms[room] {
				if visited[v] {
					continue
				}

				queue = append(queue, v)
			}
		}
	}

	return false
}

func test(m [][]int, expectedResult bool) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Room keys: %v \n", m)

	result := canVisitAllRooms(m)

	fmt.Printf("Can visit all rooms: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m := [][]int{
		{1},
		{2},
		{3},
	}

	test(m, true)
}

func test2() {
	m := [][]int{
		{1, 3},
		{3, 0, 1},
		{2},
		{0},
	}

	test(m, false)
}

func main() {
	// 841. Keys and Rooms
	test1()
	test2()
}
