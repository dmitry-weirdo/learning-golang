package main

import "fmt"

var directions = [][]int{
	{-1, 0}, // top
	{0, 1},  // right
	{1, 0},  // bottom
	{0, -1}, // left
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]

	if originalColor == color {
		fmt.Printf("Image[%v][%v] is already in color %v. Nothing to do", sr, sc, color)
		return image
	}

	rows := len(image)
	columns := len(image[0])

	dfs(image, rows, columns, sr, sc, originalColor, color)

	// we're updating the colours in the original image matrix
	return image
}

func dfs(m [][]int, rows, columns, i, j int, originalColor int, newColor int) {
	// something not correct -> return
	if !cellExists(rows, columns, i, j) ||
		m[i][j] != originalColor {
		return
	}

	// fill the current point
	m[i][j] = newColor

	for _, v := range directions { // top, right, bottom, left
		newRow := i + v[0]
		newColumn := j + v[1]

		dfs(m, rows, columns, newRow, newColumn, originalColor, newColor)
	}
}

func cellExists(rows, columns, row, column int) bool {
	return (0 <= row) &&
		(row < rows) &&
		(0 <= column) &&
		(column < columns)
}

func test(image [][]int, sr int, sc int, color int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("image matrix:\n%v\n", image)
	fmt.Printf("Row: %v \n", sr)
	fmt.Printf("Column: %v \n", sc)
	fmt.Printf("New color: %v \n", color)

	result := floodFill(image, sr, sc, color)

	fmt.Printf("Image after flood fill: \n%v \n", result)
}

func test1() {
	m := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}

	row := 1
	col := 1
	newColor := 2

	test(m, row, col, newColor)
}

func main() {
	// 733. Flood Fill
	test1()
}
