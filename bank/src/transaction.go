package src

import "time"

type Transaction struct {
	identifier string
	amount     float32
	timestamp  time.Time
	// TODO create operations enum
	operation string
}

type TranferTransaction struct {
	Transaction
	from string
	to   string
}

func newTransaction(amount float32, operation string) Transaction {
	return Transaction{
		amount:    amount,
		operation: operation,
		timestamp: time.Now(),
	}
}

func NewTransferTransaction(amount float32, from, to string) TranferTransaction {
	return TranferTransaction{Transaction: newTransaction(amount, "WITHDRAW"), from: from, to: to}
}

func DepositTransaction(amount float32) Transaction {
	return newTransaction(amount, "DEPOSIT")
}

func WithdrawTransaction(amount float32) Transaction {
	return newTransaction(amount, "WITHDRAW")
}
