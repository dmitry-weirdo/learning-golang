package main

import (
	"fmt"
	"slices"
)

type UnionFind struct {
	parents []int // if parent[i] = i, it is the root, else it's the index of the parent
	sizes   []int // sizes of the tree for every element
}

func newUnionFind(n int) UnionFind {
	parents := make([]int, n)
	sizes := make([]int, n)

	for i := range n {
		// every group is just a root
		parents[i] = i

		// every group has a size of 1
		sizes[i] = 1
	}

	return UnionFind{
		parents: parents,
		sizes:   sizes,
	}
}

func (uf UnionFind) find(x int) int { // recursive version
	if uf.parents[x] == x { // parent points to itself -> reached the root
		return x
	}

	// path compression -> set the root to every parents[i] in the chain
	uf.parents[x] = uf.find(uf.parents[x])

	return uf.parents[x]
}

func (uf UnionFind) print() {
	fmt.Printf("Parents: %v \n", uf.parents)
	fmt.Printf("Sizes: %v \n", uf.sizes)
}

func (uf UnionFind) union(x, y int) bool { // returns false if they're already in the same set
	// these find will perform path compression
	rootX := uf.find(x)
	rootY := uf.find(y)

	fmt.Printf("root of %d: %d, root of %d: %d\n", x, rootX, y, rootY)

	// x and y are already in the same set -> nothing to merge
	if rootX == rootY {
		fmt.Printf("Element %v and %v already belong to the same root %v. Nothing to merge. \n", x, y, rootX)
		return false
	}

	// merge the smaller group into the bigger group
	// todo: ideally, we should merge the tree with smaller depth into the tree with bigger depth
	if uf.sizes[rootX] < uf.sizes[rootY] { // merge x into y
		fmt.Printf("sizes[%v] = %v < sizes[%v] = %v. Merging root %v into root %v \n", rootX, uf.sizes[rootX], rootY, uf.sizes[rootY], rootX, rootY)

		uf.parents[rootX] = rootY
		uf.sizes[rootY] += uf.sizes[rootX]
	} else { // merge y into x
		fmt.Printf("sizes[%v] = %v >= sizes[%v] = %v. Merging root %v into root %v \n", rootX, uf.sizes[rootX], rootY, uf.sizes[rootY], rootY, rootX)

		uf.parents[rootY] = rootX
		uf.sizes[rootX] += uf.sizes[rootY]
	}

	return true
}

func accountsMerge(accounts [][]string) [][]string {
	// accounts[0] is name
	// accounts[1...] are emails
	n := len(accounts)

	// uf will be initialized with account indexes [0; n-1], each accountIndex a root pointing to itself
	uf := newUnionFind(n)

	fmt.Printf("Initialized the union-find structure of size %v. \n", n)
	uf.print()
	// iterate emails and map them to accountId
	// union the accounts if current email is already pointing to another account
	emailsToAccounts := make(map[string]int)

	for accountIndex, arr := range accounts {
		fmt.Println()
		fmt.Println("========================")
		fmt.Printf("Handling account %v (%v) \n", accountIndex, getAccountName(accounts, accountIndex))

		for emailIndex := 1; emailIndex < len(arr); emailIndex++ {
			email := arr[emailIndex]

			existingAccountForEmail, ok := emailsToAccounts[email]

			if ok {
				fmt.Printf("Email %v is already mapped to account %v. Merging account %v into account %v. \n", email, existingAccountForEmail, accountIndex, existingAccountForEmail)

				uf.union(accountIndex, existingAccountForEmail)
			} else {
				fmt.Printf("Email %v is not yet mapped to any account. Adding new email %v to account %v. \n", email, email, accountIndex)

				emailsToAccounts[email] = accountIndex
			}
		}
	}

	uf.print()

	// all grouped account to emails
	// no Set<String> in Go -> use a map to avoid emails duplication
	accountToEmail := make(map[int]map[string]bool) // only contains root accounts, bool to take just 1 extra byte per set value

	for accountIndex, arr := range accounts {
		// all accounts of root are merged into root
		rootIndex := uf.find(accountIndex)

		// set of emails for current root account
		var rootAccountEmails map[string]bool

		m, ok := accountToEmail[rootIndex]
		if ok {
			rootAccountEmails = m
		} else {
			rootAccountEmails = make(map[string]bool)
			accountToEmail[rootIndex] = rootAccountEmails
		}

		// iterate emails of the current account (not necessary root)
		// add emails of the current account to the root
		for emailIndex := 1; emailIndex < len(arr); emailIndex++ {
			email := arr[emailIndex]

			rootAccountEmails[email] = true
		}
	}

	fmt.Printf("Root accounts to unique emails map (emails are not sorted):\n%v \n", accountToEmail)

	totalRootAccounts := len(accountToEmail)
	fmt.Printf("Total root accounts: %v \n", totalRootAccounts)

	// todo: extract result collection to a separate function
	// collect the result - rootAccountName
	var result = make([][]string, totalRootAccounts)

	resultIndex := 0
	for accountIndex, emails := range accountToEmail {
		// emails is map[string]bool, so that it is a set with non-repeating emails
		accountName := getAccountName(accounts, accountIndex)

		accountArray := make([]string, len(emails)+1) // 0-th will be account name
		accountArray[0] = accountName

		emailIndex := 1
		for email := range emails { // emails is a email->bool map
			accountArray[emailIndex] = email

			emailIndex++
		}

		// we can sort just the part of the array with emails!
		slices.Sort(accountArray[1:])

		result[resultIndex] = accountArray

		resultIndex++
	}

	return result
}

func getAccountName(accounts [][]string, i int) string {
	return accounts[i][0]
}

func printAccounts(accounts [][]string) {
	for accountIndex, arr := range accounts {
		name := getAccountName(accounts, accountIndex)
		fmt.Printf("Account[%v] (%v). Emails: %v \n", accountIndex, name, arr[1:])
	}
}

func test1() {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}

	fmt.Printf("Initial accounts: \n")
	printAccounts(accounts)

	mergedAccounts := accountsMerge(accounts)

	fmt.Printf("Merged accounts: \n")
	printAccounts(mergedAccounts)
}

func main() {
	// 721. Accounts Merge
	test1()
}
