package src

import "fmt"

type Account struct {
	ID           string
	Owner        string
	Balance      float32
	Transactions []Transaction
}

func NewAccount(id string, owner string, balance float32) Account {
	return Account{
		ID:           id,
		Owner:        owner,
		Balance:      balance,
		Transactions: make([]Transaction, 0),
	}
}

func (d *Account) Deposit(transaction Transaction) error {
	d.Balance += transaction.Amount
	d.Transactions = append(d.Transactions, transaction)
	return nil
}

func (d *Account) Withdraw(transaction Transaction) error {
	if d.Balance-transaction.Amount < 0 {
		return fmt.Errorf("Cannot withdraw this amount %f.", transaction.Amount)
	}

	d.Balance -= transaction.Amount
	d.Transactions = append(d.Transactions, transaction)

	return nil
}

func (d *Account) GetBalance() float32 {
	return float32(d.Balance)
}

func (d Account) PrintStatement() []Transaction {
	return d.Transactions
}
