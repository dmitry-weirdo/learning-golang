package main

import "fmt"

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{
		balance: balance, // todo: it's better to make an array copy, then we can even add the 0-th element and not always make -1 to indices
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.IsValidAccount(account1) || !this.IsValidAccount(account2) {
		return false
	}

	// from account 1
	if this.balance[this.GetAccountIndex(account1)] < money {
		return false
	}

	this.balance[this.GetAccountIndex(account1)] -= money

	// to account 2
	this.balance[this.GetAccountIndex(account2)] += money

	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.IsValidAccount(account) {
		return false
	}

	// add money is always allowed
	// todo: think about overflow
	this.balance[this.GetAccountIndex(account)] += money

	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.IsValidAccount(account) {
		return false
	}

	if this.balance[this.GetAccountIndex(account)] < money {
		return false
	}

	this.balance[this.GetAccountIndex(account)] -= money

	return true
}

func (this *Bank) IsValidAccount(account int) bool {
	return (1 <= account) && (account <= len(this.balance))
}

func (this *Bank) GetAccountIndex(account int) int {
	return account - 1 // account number 1 is in index 0
}

func (this *Bank) GetAccountMoney(account int) int64 {
	if !this.IsValidAccount(account) {
		panic(fmt.Sprintf("Incorrect account number: %v", account))
	}

	return this.balance[this.GetAccountIndex(account)]
}

func testDeposit(b Bank, account int, money int64, expectedResult bool) {
	// this method assumes that the account number is value
	balanceBefore := b.GetAccountMoney(account)

	result := b.Deposit(account, money)

	balanceAfter := b.GetAccountMoney(account)

	fmt.Println()
	fmt.Printf("Account %d: added deposit $%v from $%v to $%v. Success: %v. \n", account, money, balanceBefore, balanceAfter, result)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}

	var expectedBalanceAfter int64

	if !result { // balance must not change
		expectedBalanceAfter = balanceBefore
	} else { // balance must increase on money
		expectedBalanceAfter = balanceBefore + money
	}

	if balanceAfter != expectedBalanceAfter {
		fmt.Printf("FAILURE: expected balanceAfter = %v, actual balanceAfter = %v \n", expectedBalanceAfter, balanceAfter)
	}
}

func testWithdraw(b Bank, account int, money int64, expectedResult bool) {
	// this method assumes that the account number is value
	balanceBefore := b.GetAccountMoney(account)

	result := b.Withdraw(account, money)

	balanceAfter := b.GetAccountMoney(account)

	fmt.Println()
	fmt.Printf("Account %d: withdrew $%v from $%v to $%v. Success: %v. \n", account, money, balanceBefore, balanceAfter, result)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}

	var expectedBalanceAfter int64

	if !result { // balance must not change
		expectedBalanceAfter = balanceBefore
	} else { // balance must decrease on money
		expectedBalanceAfter = balanceBefore - money
	}

	if balanceAfter != expectedBalanceAfter {
		fmt.Printf("FAILURE: expected balanceAfter = %v, actual balanceAfter = %v \n", expectedBalanceAfter, balanceAfter)
	}
}

func testTransfer(b Bank, accountFrom, accountTo int, money int64, expectedResult bool) {
	// this method assumes that the account1 and account2 numbers is value
	balanceFromBefore := b.GetAccountMoney(accountFrom)
	balanceToBefore := b.GetAccountMoney(accountTo)

	result := b.Transfer(accountFrom, accountTo, money)

	balanceFromAfter := b.GetAccountMoney(accountFrom)
	balanceToAfter := b.GetAccountMoney(accountTo)

	fmt.Println()
	fmt.Printf("Tried to transfer money %v from account %v to account %v. Success: %v. \n", money, accountFrom, accountTo, result)
	fmt.Printf("Account %v (from): balance changed $%v -> $%v. \n", accountFrom, balanceFromBefore, balanceFromAfter)
	fmt.Printf("Account %v (to): balance changed $%v -> $%v. \n", accountTo, balanceToBefore, balanceToAfter)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}

	// check balance of accountFrom
	var expectedBalanceFromAfter int64

	if !result { // balance must not change
		expectedBalanceFromAfter = balanceFromBefore
	} else { // balance must decrease on money
		expectedBalanceFromAfter = balanceFromBefore - money
	}

	if balanceFromAfter != expectedBalanceFromAfter {
		fmt.Printf("FAILURE: expected balanceFromAfter = %v, actual balanceFromAfter = %v \n", expectedBalanceFromAfter, balanceFromAfter)
	}

	// check balance of accountTo
	var expectedBalanceToAfter int64

	if !result { // balance must not change
		expectedBalanceToAfter = balanceToBefore
	} else { // balance must increase on money
		expectedBalanceToAfter = balanceToBefore + money
	}

	if balanceToAfter != expectedBalanceToAfter {
		fmt.Printf("FAILURE: expected balanceToAfter = %v, actual balanceToAfter = %v \n", expectedBalanceToAfter, balanceToAfter)
	}
}

func testWithdrawWithInvalidAccount(b Bank, account int, money int64) {
	expectedValidAccount := false
	validAccount := b.IsValidAccount(account)

	if validAccount != expectedValidAccount {
		fmt.Printf("FAILURE: expected validAccount = %v, actual validAccount = %v \n", expectedValidAccount, validAccount)
	}

	expectedResult := false
	result := b.Withdraw(account, money)

	fmt.Println()
	fmt.Printf("Tried to withdraw money %d from an invalid account %d. Success: %v. Is valid account: %v. \n", money, account, result, validAccount)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test() {
	initialBalances := []int64{10, 100, 20, 50, 30}
	b := Constructor(initialBalances)
	testWithdraw(b, 3, 10, true)
	testTransfer(b, 5, 1, 20, true)
	testDeposit(b, 5, 20, true)
	testTransfer(b, 3, 4, 15, false)          // fail -> current balance of account 3 is 10 < 15
	testWithdrawWithInvalidAccount(b, 10, 50) // account 10 does not exist
}

func main() {
	// 2043. Simple Bank System
	test()
}
