package main

import "fmt"

func canWinNim(n int) bool {
	// We can take 1, 2, 3 stones, i.e. on every other player turn, other player can take the compliment to 4,
	// thus remaining the mod to 4.
	//
	// This is the optimal strategy of both players.
	// So the (mod 4) will remain the same on every turn of 2 players.
	// On remaining 1, 2, 3 - 1st player wins by taking all these stones.
	// On remaining 4 - on any 1st player move, 2nd player takes all other stones.
	//
	// I.e. on any (mod 4 != 0), the first player takes 1 stone and then just responds:
	// 5 - we take 1 and win since player 2 has 4
	// 6 - we take 2 and win since player 2 has 4
	// 7 - we take 3 and win since player 2 has 4

	return n%4 != 0
}

func test(x int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number of stones: %v \n", x)

	result := canWinNim(x)

	fmt.Printf("Player 1 can win: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(4, false)
}

func test2() {
	test(1, true)
}

func test3() {
	test(3, true)
}

func test4() {
	test(2, true)
}

func test5() {
	test(5, true)
}

func test6() {
	test(6, true)
}

func test7() {
	test(7, true)
}

func test8() {
	test(8, false)
}

func main() {
	// 292. Nim Game
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
}
