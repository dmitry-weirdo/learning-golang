package main

import (
	"fmt"
)

var directions = [][]int{
	{-1, 0}, // top
	{0, 1},  // right
	{1, 0},  // bottom
	{0, -1}, // left
}

const (
	EMPTY    = iota // "."
	START           // "S"
	LITTER          // "L"
	OBSTACLE        // "X"
	RESTORE         // "R"
)

type CellInfo struct {
	i      int // row
	j      int // column
	mask   int // bit mask of what litter cells has been collected
	energy int // how many energy left in this cell
	steps  int // how many steps to this cell
}

func minMoves(classroom []string, energy int) int {
	m := stringArrayToIntMatrix(classroom)

	//fmt.Println("String converted to int matrix:")
	//PrintIntMatrix(m)

	rows, columns := getRowsAndColumns(m)

	// find start and litter cells

	litterMasks := createIntMatrix(rows, columns)

	litterCellsCount := 0
	startRow := -1
	startColumn := -1

	// Since there are just up to 10 litter cells, we can count the handled litter cells with a bitmask.
	// For every litter cell, we save the bit just for this cell  in its litterMasks[i][j].
	// 0th mask is 0001, 1st mask is 0010, etc.

	for i := range rows {
		for j := range columns {
			if m[i][j] == START {
				startRow = i
				startColumn = j
				continue
			}

			if m[i][j] == LITTER {
				litterMasks[i][j] = 1 << litterCellsCount // for count == 0, it will be just 0001
				litterCellsCount++
			}
		}
	}

	//fmt.Printf("Start cell: [%v][%v] \n", startRow, startColumn)
	//fmt.Printf("Total litter cells: %v \n", litterCellsCount)

	//fmt.Println("Litter cells matrix:")
	PrintIntMatrix(litterMasks)

	// For every [i][j], we have an array of all masks ( 2 ^ litterCellsCount )
	// To save the bestEnergy for this mask.
	// This bestEnergy is used to prevent infinite cycle over the same cells
	allLitterHandledMask := (1 << litterCellsCount) - 1 // mask with all 1, target value to stop the iteration
	//fmt.Printf("All litter handled mask: %v = %b \n", allLitterHandledMask, allLitterHandledMask)

	// all energies set to -1 to be fewer than any possible minimum
	//bestEnergy := int[rows][columns][allLitterHandledMask]
	bestEnergy := make([][][]int, rows)
	for i := range rows {
		bestEnergy[i] = make([][]int, columns)

		for j := range columns {
			/// size = allMasks + 1 since we also store the 0-th index
			// examples:
			// - count = 1, we need to store 0 and 1, i.e. 2 values
			// - count = 3, we need to store values from 0 to 7 (111), i.e. 8 values
			bestEnergy[i][j] = createIntArrayWithDefaultValues(allLitterHandledMask+1, -1)
		}
	}

	// start cell is set to full energy for 0th mask (no litter collected)
	bestEnergy[startRow][startColumn][0] = energy

	// BFS until we reach allLitterHandledMask
	startCellInfo := CellInfo{
		i:      startRow,
		j:      startColumn,
		mask:   0,      // no litter handled
		energy: energy, // full energy at the start
		steps:  0,      // no steps at the start yet
	}

	queue := make([]CellInfo, 0)
	queue = append(queue, startCellInfo)

	for len(queue) > 0 {
		// level-by-level is not necessary, since we can return back to the same cells

		// poll from queue
		ci := queue[0]
		queue = queue[1:]

		if ci.mask == allLitterHandledMask { // reached the state with all litter collected
			return ci.steps
		}

		if ci.energy == 0 { // no energy left to make the next move -> stop processing this cell
			continue
		}

		// try to move in all 4 directions from this cell
		for _, d := range directions {
			i := ci.i + d[0]
			j := ci.j + d[1]

			// went over the border or into obstacle -> skip this direction
			if !cellExists(rows, columns, i, j) || (m[i][j] == OBSTACLE) {
				continue
			}

			newEnergy := ci.energy - 1 // moving to the neighbor cell will decrease our energy by 1
			if m[i][j] == RESTORE {    // but the R (restore) cell will reset us to full energy
				newEnergy = energy
			}

			// if neighbor cell is litter cell, add the bit of this cell to the collected litter bit mask
			newMask := ci.mask | litterMasks[i][j]

			// !!! Trick to reduce cycles over the same cells
			//-> only move to the next cell if its energy state is better than energy state for this cell's litter big mask
			if newEnergy <= bestEnergy[i][j][newMask] {
				continue
			}

			// all checks passed -> add the neighbor cell to the queue
			bestEnergy[i][j][newMask] = newEnergy

			neighborCellInfo := CellInfo{
				i:      i,
				j:      j,
				mask:   newMask,
				energy: newEnergy,
				steps:  ci.steps + 1, // moving to the next cell will increase 1 step from the current cell
			}

			queue = append(queue, neighborCellInfo)
		}
	}

	// not reached the target litter mask -> cannot collect all litter cells
	return -1
}

func stringArrayToIntMatrix(s []string) [][]int {
	rows := len(s)
	columns := len(s[0])

	m := createIntMatrix(rows, columns)

	for i := range rows {
		for j := range columns {
			m[i][j] = stringFieldToInt(s[i][j])
		}
	}

	return m
}

func stringFieldToInt(ch byte) int {
	switch ch {
	case '.':
		return EMPTY
	case 'S':
		return START
	case 'L':
		return LITTER
	case 'X':
		return OBSTACLE
	case 'R':
		return RESTORE
	default:
		panic(fmt.Sprintf("Incorrect charater: %c", ch))
	}
}

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
}

func PrintIntMatrix(mat [][]int) {
	rows, columns := getRowsAndColumns(mat)

	for i := range rows {
		for j := range columns {
			fmt.Printf("%v ", mat[i][j])
		}

		fmt.Println()
	}
}

func getRowsAndColumns(mat [][]int) (rows, columns int) {
	if len(mat) <= 0 {
		return 0, 0
	}

	return len(mat), len(mat[0]) // !!! we assume that all rows have the same length
}

func cellExists(rows, columns, row, column int) bool {
	return (0 <= row) &&
		(row < rows) &&
		(0 <= column) &&
		(column < columns)
}

func createIntArrayWithDefaultValues(n int, defaultValue int) []int {
	a := make([]int, n)

	for i := range n {
		a[i] = defaultValue
	}

	return a
}

func test(arr []string, energy int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Cells: %v \n", arr)
	fmt.Printf("Energy: %v \n", energy)

	result := minMoves(arr, energy)

	fmt.Printf("Minimum moves to collect all litter: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m := []string{
		"S.",
		"XL",
	}

	energy := 2

	expected := 2

	test(m, energy, expected)
}

func test2() {
	m := []string{
		"LS",
		"RL",
	}

	energy := 4

	expected := 3

	test(m, energy, expected)
}

func test3() {
	m := []string{
		"L.S",
		"RXL",
	}

	energy := 3

	expected := -1

	test(m, energy, expected)
}

func test4() {
	m := []string{
		"S..L",
		".XXX",
		"LXXX",
	}

	energy := 10

	expected := 7 // go down -> up -> right (better than right -> left -> right that will be 9 steps)

	test(m, energy, expected)
}

func main() {
	// 3568. Minimum Moves to Clean the Classroom
	test1()
	test2()
	test3()
	test4()
}
