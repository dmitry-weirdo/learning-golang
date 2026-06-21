package main

import (
	"fmt"
	"slices"
)

func rankTeams(votes []string) string {
	totalTeams := len(votes[0])

	// collect team names
	teams := []byte(votes[0])
	slices.Sort(teams)

	teamToRow := make(map[byte]int)
	for i, v := range teams {
		teamToRow[v] = i
	}

	// todo: remove this, debug-only
	var teamsAsStrings = make([]string, totalTeams)

	for i, v := range teams {
		teamsAsStrings[i] = string(v)
	}

	fmt.Printf("Different teams (total %v teams): %v | %v \n", totalTeams, teams, teamsAsStrings)

	var m = make([][]int, totalTeams)

	for i := 0; i < totalTeams; i++ {
		row := make([]int, totalTeams+1)
		row[0] = int(teams[i]) // 0th in a row is a team name

		m[i] = row

		// others should be 0 by default
	}

	fmt.Println("Added team names to the matrix:")
	printMatrix(m)

	for _, vote := range votes {
		for j := 0; j < totalTeams; j++ {
			team := vote[j]

			rowIndex := teamToRow[team]
			columnIndex := 1 + j // skip the 0th with team name

			//fmt.Printf("Increasing team %v, vote %v \n", string(team), j)

			inc(m, rowIndex, columnIndex)
		}
	}

	fmt.Println("Added vote counts to matrix:")
	printMatrix(m)

	// sort the matrix
	slices.SortFunc(m, func(a, b []int) int {

		for i := 1; i < totalTeams+1; i++ { // starting from 1, i.e. skip the team name
			if a[i] != b[i] {
				return b[i] - a[i] // we sort by votes desc!
			}
		}

		// sort alphabetically by team name if all positions are equal
		return a[0] - b[0]
	})

	fmt.Println("Sorted the matrix:")
	printMatrix(m)

	// collect the result - 0th column from every row
	resultBytes := make([]byte, totalTeams)

	for i := 0; i < totalTeams; i++ {
		resultBytes[i] = byte(m[i][0]) // convert from int
	}

	return string(resultBytes)
}

func inc(mat [][]int, row int, column int) {
	curr := mat[row][column]
	mat[row][column] = curr + 1
}

func printMatrix(mat [][]int) {
	rows := len(mat)
	columns := len(mat[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if j == 0 { // print as string instead of numeric byte
				fmt.Printf("%v ", string(mat[i][j]))
			} else {
				fmt.Printf("%v ", mat[i][j])
			}

		}

		fmt.Println()
	}
}

func handle(votes []string) {
	fmt.Println()
	fmt.Printf("======================= \n")

	fmt.Printf("Teams: %v \n", votes)

	ranked := rankTeams(votes)

	fmt.Printf("Teams ranked: %v \n", ranked)
}

func main() {
	handle([]string{"ABC", "ACB", "ABC", "ACB", "ACB"})
	//handle([]string{"CAB", "ACB", "ABC", "ACB", "ACB"})
	handle([]string{"WXYZ", "XYZW"})
	handle([]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"})
	handle([]string{"AB", "BA"})
	handle([]string{"AXYB", "AYXB", "AXYB", "AYXB"}) // letters can be non-consecutive!
}
