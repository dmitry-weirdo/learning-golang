package main

import "fmt"

const HOLE_NODE = 0

const MOUSE_STARTING_NODE = 1
const CAT_STARTING_NODE = 2

const DRAW = 0
const MOUSE_WINS = 1
const CAT_WINS = 2

const DP_UNDEFINED = 0 // default state
const DP_DRAW = 1
const DP_MOUSE_WINS = 2
const DP_CAT_WINS = 3

func catMouseGame(graph [][]int) int { // graph is already adj!

	// This uses a magic limitation that after (2 * n) moves, we can guarantee it's a draw,
	// since (2 * N) should be enough to find a winning position.
	return catMouseGame_dp_topToBottom(graph)
}

func catMouseGame_dp_topToBottom(adj [][]int) int {
	// see https://www.youtube.com/watch?v=oGKnucI_ejw

	n := len(adj)

	// This uses a magic limitation that after (2 * N) moves, we can guarantee it's a draw,
	// since (2 * N) should be enough to find a winning position.
	movesLimit := 2 * n
	//fmt.Printf("N (number of nodes): %v, moves limit to draw: %v \n", n, movesLimit)

	// dp[mousePos][catPos][countOfMoves]
	// dp sizes: [n][n][2 * n + 1]
	dp := make([][][]int, n)
	for i := range n {
		dp[i] = make([][]int, n)

		for j := range n {
			dp[i][j] = make([]int, movesLimit+1)
		}
	}

	// i - mouse position
	// j - cat position
	// moves - total moves done. We can define whose move it is by (moves % 2) == 0 - > mouse move
	var dfs func(i, j, moves int) int

	dfs = func(i, j, moves int) int {
		//fmt.Printf("Mouse pos I: %v, cat pos J: %v, moves: %v \n", i, j, moves)

		// This uses a magic limitation that after (2 * N) moves, we can guarantee it's a draw,
		// since (2 * N) should be enough to find a winning position.
		if moves > movesLimit {
			//dp[i][j][moves] = DP_DRAW // this will be out of bounds
			return DRAW
		}

		// state already pre-calculated -> return it
		switch dp[i][j][moves] {
		case DP_DRAW:
			return DRAW

		case DP_MOUSE_WINS:
			return MOUSE_WINS

		case DP_CAT_WINS:
			return CAT_WINS

			// if dp is DP_UNDEFINED -> we return nothing
		}

		if i == j { // mouse and cat are in the same node -> cat wins
			dp[i][j][moves] = DP_CAT_WINS
			return CAT_WINS
		}

		if i == HOLE_NODE { // mouse is at the hole node 0 -> mouse wins
			dp[i][j][moves] = DP_MOUSE_WINS
			return MOUSE_WINS
		}

		mouseTurn := moves%2 == 0

		if mouseTurn { // mouse turn
			canDraw := false

			for _, neighbor := range adj[i] { // all possible edges from mouse node
				neighborResult := dfs(neighbor, j, moves+1) // mouse moves to neighbor

				if neighborResult == MOUSE_WINS { // we can move to a mouse winning state -> current state is also a mouse win
					dp[i][j][moves] = DP_MOUSE_WINS
					return MOUSE_WINS
				} else if neighborResult == DRAW {
					canDraw = true
				}
			}

			if canDraw { // mouse cannot win, but can draw
				dp[i][j][moves] = DP_DRAW
				return DRAW
			} else { // mouse loses, cat wins
				dp[i][j][moves] = DP_CAT_WINS
				return CAT_WINS
			}
		} else { // cat turn
			canDraw := false

			for _, neighbor := range adj[j] { // all possible edges from cat node
				if neighbor == HOLE_NODE { // cat cannot move to the hole node 0
					continue
				}

				neighborResult := dfs(i, neighbor, moves+1) // cat moves to neighbor

				if neighborResult == CAT_WINS { // we can move to a cat winning state -> current state is also a cat win
					dp[i][j][moves] = DP_CAT_WINS
					return CAT_WINS
				} else if neighborResult == DRAW {
					canDraw = true
				}
			}

			if canDraw { // cat cannot win, but can draw
				dp[i][j][moves] = DP_DRAW
				return DRAW
			} else { // cat loses, mouse wins
				dp[i][j][moves] = DP_MOUSE_WINS
				return MOUSE_WINS
			}
		}
	}

	return dfs(MOUSE_STARTING_NODE, CAT_STARTING_NODE, 0)
}

func test(adj [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Adjacency list: %v \n", adj)

	result := catMouseGame(adj)

	fmt.Printf("Winner: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[][]int{
			{2, 5},
			{3},
			{0, 4, 5},
			{1, 4, 5},
			{2, 3},
			{0, 2, 3},
		},
		0, // draw
	)
}

func test2() {
	test(
		[][]int{
			{1, 3},
			{0},
			{3},
			{0, 2},
		},
		1, // mouse wins
	)
}

func test3() {
	// failing test-case 66/92
	test(
		[][]int{
			{5, 7, 9},
			{3, 4, 5, 6},
			{3, 4, 5, 8},
			{1, 2, 6, 7},
			{1, 2, 5, 7, 9},
			{0, 1, 2, 4, 8},
			{1, 3, 7, 8},
			{0, 3, 4, 6, 8},
			{2, 5, 6, 7, 9},
			{0, 4, 8},
		},
		1, // mouse wins
	)
}

func main() {
	// 913. Cat and Mouse
	test1()
	test2()
	test3()
}
