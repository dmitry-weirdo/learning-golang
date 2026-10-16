package main

import "fmt"

func minTransfers(transactions [][]int) int {
	return minTransfers_backtracking(transactions)
}

func minTransfers_backtracking(transactions [][]int) int {
	// For every account, calculate its balance
	m := make(map[int]int)

	for _, v := range transactions {
		from := v[0]
		to := v[1]
		amount := v[2]

		m[from] += amount
		m[to] -= amount
	}

	//fmt.Printf("Amounts map: %v \n", m)

	// Filter out accounts with non-0 balances
	balances := make([]int, 0)

	for _, v := range m {
		if v == 0 {
			continue
		}

		balances = append(balances, v)
	}

	//fmt.Printf("Non-zero balances: %v \n", balances)

	n := len(balances)
	//fmt.Printf("Total non-zero balances: %v \n", n)

	const TOO_BIG_BALANCE = 1_000_000

	// backtracking - trying to move the balance of the current account to every other account with the opposite sign balance
	var dfs func(i int) int

	dfs = func(i int) int {
		// skip all currently 0 balances
		for (i < n) && (balances[i] == 0) {
			i++
		}

		if i == n { // all balances are 0 -> no transactions required anymore
			return 0
		}

		minTransactionsCount := TOO_BIG_BALANCE

		for j := i + 1; j < n; j++ { // try to move to all further positions with the opposite balance sign
			if balances[i]*balances[j] < 0 {
				// add i balance to j (current to next)
				balances[j] += balances[i]

				// 1 transaction for moving i -> j
				nextTransactionsCount := 1 + dfs(i+1)

				minTransactionsCount = min(minTransactionsCount, nextTransactionsCount)

				// backtrack - remove balance i from j
				balances[j] -= balances[i]
			}
		}

		return minTransactionsCount
	}

	return dfs(0)
}

func test(m [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Transactions: %v \n", m)

	result := minTransfers(m)

	fmt.Printf("Min transactions required to clear all the debts: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[][]int{
			{0, 1, 10},
			{2, 0, 5},
		},
		2,
	)
}

func test2() {
	test(
		[][]int{
			{0, 1, 10},
			{1, 0, 1},
			{1, 2, 5},
			{2, 0, 5},
		},
		1,
	)
}

func main() {
	// 465. Optimal Account Balancing
	test1()
	test2()
}
