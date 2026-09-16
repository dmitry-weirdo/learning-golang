package main

import "fmt"

const HOLE_NODE = 0

const MOUSE_STARTING_NODE = 1
const CAT_STARTING_NODE = 2

const MOUSE_TURN = 0
const CAT_TURN = 1

const DRAW = 0
const MOUSE_WINS = 1
const CAT_WINS = 2

const DP_UNDEFINED = 0 // default state
const DP_DRAW = 1
const DP_MOUSE_WINS = 2
const DP_CAT_WINS = 3

func catMouseGame(graph [][]int) int { // graph is already adj!
	// passes in 95-120 ms, is much faster!
	return catMouseGame_dp_bottomUp(graph)

	// This uses a magic limitation that after (2 * n) moves, we can guarantee it's a draw,
	// since (2 * N) should be enough to find a winning position.

	// (2 * N) cut-off actually does not pass on test-case, 66/92
	// Changing the heuristics to (5 * N) is working, but this is an ugly hack!
	// With (5 * N), passes in 560-780 ms, i.e. very slow.
	//return catMouseGame_dp_topToBottom(graph)
}

type GamePosition struct {
	mousePos  int
	catPos    int
	whoseTurn int
}

func catMouseGame_dp_bottomUp(adj [][]int) int {
	n := len(adj)

	// dp[mousePos][catPos][whoseTurn]
	// dp sizes: [n][n][2]
	dp := make([][][]int, n)
	for i := range n {
		dp[i] = make([][]int, n)

		for j := range n {
			dp[i][j] = make([]int, 2)
		}
	}

	// initial -> populate all known winning states
	queue := make([]*GamePosition, 0)

	for i := 1; i < n; i++ { // skip node 0
		for turn := range 2 { // handle both MOUSE_TURN and CAT_TURN states
			// mouse win for mousePos = 0 and any catPos and whoseTurn
			dp[0][i][turn] = DP_MOUSE_WINS
			queue = append(queue, &GamePosition{mousePos: 0, catPos: i, whoseTurn: turn})

			// cat wins if (i > 0) and mousePos == catPost
			dp[i][i][turn] = DP_CAT_WINS
			queue = append(queue, &GamePosition{mousePos: i, catPos: i, whoseTurn: turn})
		}
	}

	for len(queue) > 0 {
		// pop from queue
		pos := queue[0]
		queue = queue[1:]

		currentResult := dp[pos.mousePos][pos.catPos][pos.whoseTurn]

		previousStates := getPreviousStates(adj, pos)

		for _, prevState := range previousStates {
			prevStateResult := dp[prevState.mousePos][prevState.catPos][prevState.whoseTurn]
			if prevStateResult != DP_UNDEFINED { // previous state result already pre-calculated -> don't handle this state again
				continue
			}

			// handle the previous state

			// we know that currentResult is NOT undefined
			canAlwaysWin := false
			//if pos.whoseTurn == MOUSE_TURN {
			if prevState.whoseTurn == CAT_TURN { // clearer: current turn mouse -> prev turn cat -> if current is "cat wins", then prev is also "cat wins"
				canAlwaysWin = currentResult == DP_CAT_WINS
				//} else if pos.whoseTurn == CAT_TURN {
			} else if prevState.whoseTurn == MOUSE_TURN { // clearer: current turn cat -> prev turn mouse -> if current is "mouse wins", then prev is also "mouse wins"
				canAlwaysWin = currentResult == DP_MOUSE_WINS
			}

			if canAlwaysWin { // prev state also runs to "winner of the prev state turn maker"
				dp[prevState.mousePos][prevState.catPos][prevState.whoseTurn] = currentResult
				queue = append(queue, prevState)
			} else if allNextMovesFail(dp, adj, prevState) { // todo: this check can be optimized, to not iterate all adj[prevState] every time, we might cache/update the allNextMoveFail[node][whoseTurn] and return it immediately instead of iterating all adjacency edges
				// prev state leads to lose of "prev state turn maker"
				// "prev state turn maker"

				prevStateLoseResult := -666

				switch prevState.whoseTurn {
				case MOUSE_TURN:
					prevStateLoseResult = DP_CAT_WINS
				case CAT_TURN:
					prevStateLoseResult = DP_MOUSE_WINS
				default:
					panic(fmt.Sprintf("Unknown prevState.whoseTurn value: %v.", prevState.whoseTurn))
				}

				dp[prevState.mousePos][prevState.catPos][prevState.whoseTurn] = prevStateLoseResult
				queue = append(queue, prevState)
			}
		}

	}

	// return result for mousePos = 1, catPos = 2, whoseTurn = MOUSE_TURN
	return converStartingDpResultToReturnValue(dp)
}

func getPreviousStates(adj [][]int, pos *GamePosition) []*GamePosition {
	result := make([]*GamePosition, 0)

	if pos.whoseTurn == MOUSE_TURN {
		// previous turn was cat turn from catPos
		// -> prev turn was catPosNeighbor -> catPos
		for _, catPosNeighbor := range adj[pos.catPos] { // graph is undirected, so by adj we can go both directions
			if catPosNeighbor == HOLE_NODE { // cat cannot move to hole node 0
				continue
			}

			previousPos := &GamePosition{
				mousePos:  pos.mousePos,
				catPos:    catPosNeighbor,
				whoseTurn: CAT_TURN,
			}

			result = append(result, previousPos)
		}
	} else {
		// previous turn was mouse turn from mousePos
		// -> prev turn was mousePosNeighbor -> mousePos
		for _, mousePosNeighbor := range adj[pos.mousePos] { // graph is undirected, so by adj we can go both directions
			previousPos := &GamePosition{
				mousePos:  mousePosNeighbor,
				catPos:    pos.catPos,
				whoseTurn: MOUSE_TURN,
			}

			result = append(result, previousPos)
		}
	}

	return result
}

func allNextMovesFail(dp [][][]int, adj [][]int, pos *GamePosition) bool {
	if pos.whoseTurn == MOUSE_TURN { // current turn is mouse turn
		for _, mousePosNeighbor := range adj[pos.mousePos] { // graph is undirected, so by adj we can go both directions
			if dp[mousePosNeighbor][pos.catPos][CAT_TURN] != DP_CAT_WINS { // in any next position cat is not yet winning -> return false
				return false
			}
		}
	} else { // current turn is cat turn
		for _, catPosNeighbor := range adj[pos.catPos] { // graph is undirected, so by adj we can go both directions
			if catPosNeighbor == HOLE_NODE { // cat cannot move to hole node 0 // todo: this check may be overkill since this neighbor will be a mouse win
				continue
			}

			if dp[pos.mousePos][catPosNeighbor][MOUSE_TURN] != DP_MOUSE_WINS { // in any next position mouse is not yet winning -> return false
				return false
			}
		}
	}

	return true
}

func converStartingDpResultToReturnValue(dp [][][]int) int {
	resultValue := dp[MOUSE_STARTING_NODE][CAT_STARTING_NODE][MOUSE_TURN]

	switch resultValue {
	case DP_UNDEFINED:
		return DRAW
	case DP_MOUSE_WINS:
		return MOUSE_WINS
	case DP_CAT_WINS:
		return CAT_WINS
	default:
		panic(fmt.Sprintf("Unknown dp[%v][%v][%v] value: %v.", MOUSE_STARTING_NODE, CAT_STARTING_NODE, MOUSE_TURN, resultValue))
	}
}

func catMouseGame_dp_topToBottom(adj [][]int) int {
	// see https://www.youtube.com/watch?v=oGKnucI_ejw

	n := len(adj)

	// This uses a magic limitation that after (2 * N) moves, we can guarantee it's a draw,
	// since (2 * N) should be enough to find a winning position.

	// todo: this is actually a random non-reliable heuristic, does not work with (2 * N), passes tests with (5 * N)
	movesLimit := 5 * n
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
	// test-case 66/92 is failing on (2 * N) cut off
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
