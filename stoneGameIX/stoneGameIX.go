package main

import "fmt"

func stoneGameIX(stones []int) bool {
	// we only care about mod 3 values count -> calculate them

	// Stones 0 does not make any change, they're just skipping a turn.

	// So let's start either with 1 or with 2.

	// If we start with 1, the sequence will be 1, (1, 2)*n
	// After it, we can take one remaining 1 stone (if any).
	// After it, we can take all 0 stones.
	// At this point, the next player is losing.
	// If stones already ended, Bob wins.

	// If we start with 2, the sequence will be 2, (2, 1)*n
	// After it, we can take one remaining 2 stone (if any).
	// After it, we can take all 0 stones.
	// At this point, the next player is losing.
	// If stones already ended, Bob wins.

	// Corner-case 1:
	// All stones are 0-stones -> Alice loses on the 1st move since 0 % 3 == 0.

	// Corner-case 2:
	// There are 1-stones, but no 2-stones.
	// Then the turns are 1, 1, (all 0 stones)

	// Corner-case 3:
	// There are 2-stones, but no 1-stones.
	// Then the turns are 2, 2, (all 0 stones)

	// counts of stones of every remainder 0, 1, 2
	c := make([]int, 3)

	// this will take O(n) time
	for _, v := range stones {
		c[v%3]++
	}

	fmt.Printf("Initial count of stones: %v \n", c)

	if (c[1] == 0) && (c[2] == 0) { // no way to make the first move
		return false
	}

	// there is at least one of 1-stone or 2-stone -> check both strategies
	scenarioStartWith1 := []int{c[0], c[1], c[2]}

	// these methods are symmetric, and we can just call them with swapped c[1] and c[2]
	scenarioStartWith2 := []int{c[0], c[2], c[1]}

	return firstPlayerCanWin(scenarioStartWith1) || firstPlayerCanWin(scenarioStartWith2)
}

func firstPlayerCanWin(c []int) bool {
	// we MUST start with c[1] in this scenario, else it would be the opposite scenario

	// corner-case -> no c[1] stones -> this scenario does not work
	if c[1] == 0 {
		return false
	}

	// corner-case -> 1 c[1] stone
	if c[1] == 1 {
		if c[2] == 0 { // stones exhaust without mod 3 lose -> we lose.
			return false
		}

		// make all possible valid turns
		turns := 1 + c[0]

		return firstPlayerWins(turns)
	}

	// normal-case, there will be 1, (1, 2) sequences
	turns := 1 // first move -> we take 1
	c[1]--

	// take all (1, 2) turns
	oneTwoCyclesCount := min(c[1], c[2])
	turns += 2 * oneTwoCyclesCount
	c[1] -= oneTwoCyclesCount
	c[2] -= oneTwoCyclesCount

	// if possible to take one more 1 turn
	if c[1] > 0 {
		turns++
		c[1]--
	}

	if (c[1] == 0) && (c[2] == 0) {
		// rest will just exhaust -> Alice loses
		return false
	}

	// take all possible (0) turns
	turns += c[0]
	c[0] = 0

	// there are some c1 or c2 left -> next player will lose
	return firstPlayerWins(turns)
}

func firstPlayerWins(validTurnsCount int) bool {
	if validTurnsCount%2 == 1 { // last valid turn was Alice -> Alice wins
		return true
	}

	// last valid turn was Bob -> Bob wins
	return false
}

func test(arr []int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Stones array: %v \n", arr)

	result := stoneGameIX(arr)

	fmt.Printf("Alice can win: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{2, 1}, true)
}

func test2() {
	test([]int{2}, false)
}

func test3() {
	test([]int{3, 6, 9}, false)
}

func test4() {
	test([]int{5, 1, 2, 4, 3}, false)
}

func main() {
	// 2029. Stone Game IX
	test1()
	test2()
	test3()
	test4()
}
