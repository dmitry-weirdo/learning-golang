package main

import "fmt"

const (
	TOP_RIGHT = iota
	BOTTOM_LEFT
)

func findDiagonalOrder(mat [][]int) []int {
	rows := len(mat)
	columns := len(mat[0])

	i, j := 0, 0 // rowIndex, columnIndex
	index := 0   // within the result

	result := make([]int, rows*columns)

	dir := TOP_RIGHT

	var topReached, bottomReached, leftReached, rightReached bool

	for !(bottomReached && rightReached) {
		// add the current cell to result
		result[index] = mat[i][j]
		index++
		fmt.Printf("result updated: %v \n", result)
		fmt.Printf("Indexes: %v %v, dir: %v \n", i, j, dir)

		// update the current state
		topReached = isTopReached(i)
		bottomReached = isBottomReached(i, rows)
		leftReached = isLeftReached(j)
		rightReached = isRightReached(j, columns)

		fmt.Printf("topReached: %v, bottomReached %v, leftReached: %v, rightReached: %v \n", topReached, bottomReached, leftReached, rightReached)

		// update i and j according to the direction
		if dir == TOP_RIGHT { // top-right
			if topReached {
				if rightReached { // top-right corner, change dir and go down
					dir = BOTTOM_LEFT
					i += 1
				} else { // top reached, change dir and go right
					dir = BOTTOM_LEFT
					j += 1
				}
			} else if rightReached { // right reached, change dir and go down
				dir = BOTTOM_LEFT
				i += 1
			} else { // normal top-right, go diagonally, no direction change
				i -= 1
				j += 1
			}
		} else { // bottom-left
			if bottomReached {
				if leftReached { // left-bottom corner, change dir and go right
					dir = TOP_RIGHT
					j += 1
				} else { // bottom reached, change dir and go right
					dir = TOP_RIGHT
					j += 1
				}
			} else if leftReached { // left reached, change dir and go down
				dir = TOP_RIGHT
				i += 1
			} else { // normal bottom-left, go diagonally, no direction change
				i += 1
				j -= 1
			}
		}
	}

	return result
}

func isTopReached(row int) bool {
	return row == 0
}

func isBottomReached(row int, rows int) bool {
	return row == (rows - 1)
}

func isLeftReached(column int) bool {
	return column == 0
}

func isRightReached(column int, columns int) bool {
	return column == (columns - 1)
}

func printMatrix(mat [][]int) {
	rows := len(mat)
	columns := len(mat[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("%v ", mat[i][j])
		}

		fmt.Println()
	}
}

func handleMatrix(mat [][]int) {
	fmt.Println()
	fmt.Printf("======================== \n")
	printMatrix(mat)
	order := findDiagonalOrder(mat)
	fmt.Printf("Order: %v \n", order)
}

func test3x3() {
	// 123
	// 456
	// 789
	row0 := []int{1, 2, 3}
	row1 := []int{4, 5, 6}
	row2 := []int{7, 8, 9}

	mat := [][]int{
		row0,
		row1,
		row2,
	}

	handleMatrix(mat)
}

func test4x1() {
	// 1
	// 2
	// 3
	// 4
	row0 := []int{1}
	row1 := []int{2}
	row2 := []int{3}
	row3 := []int{4}

	mat := [][]int{
		row0,
		row1,
		row2,
		row3,
	}

	handleMatrix(mat)
}

func test1x4() {
	// 1, 2, 3, 4
	row0 := []int{1, 2, 3, 4}

	mat := [][]int{
		row0,
	}

	handleMatrix(mat)
}

func test2x2() {
	// 1, 2
	// 3 4
	row0 := []int{1, 2}
	row1 := []int{3, 4}

	mat := [][]int{
		row0,
		row1,
	}

	handleMatrix(mat)
}

func main() {
	// 498. Diagonal Traverse
	test3x3()
	test4x1()
	test1x4()
	test2x2()
}
