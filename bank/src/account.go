package src

import "fmt"

type Account struct {
	ID           string
	owner        string
	balance      float32
	transactions []Transaction
}

func (d Account) Deposit(transaction Transaction) error {
	d.balance += transaction.amount
	d.transactions = append(d.transactions, transaction)
	return nil
}

func (d Account) Withdraw(transaction Transaction) error {
	if d.balance-transaction.amount < 0 {
		return fmt.Errorf("Cannot withdraw this amount %f.", transaction.amount)
	}

	d.balance -= transaction.amount
	d.transactions = append(d.transactions, transaction)

	return nil
}

func (d Account) GetBalance() float32 {
	return float32(d.balance)
}

func (d Account) PrintStatement() []Transaction {
	return d.transactions
}
