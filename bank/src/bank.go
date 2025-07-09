package src

import (
	"fmt"
)

type Banking interface {
	Deposit(t Transaction)
	Withdraw(t Transaction)
}

type Bank struct {
	accounts    map[string]*Account
	transaction []Transaction
}

func NewBank() Bank {
	return Bank{
		accounts: make(map[string]*Account),
	}
}

func (b *Bank) CreateAccount(id string, account Account) {
	b.accounts[id] = &account
}

func (b *Bank) GetAccount(id string) (*Account, error) {
	account, exists := b.accounts[id]
	if exists {
		return account, nil
	}
	return &Account{}, fmt.Errorf("Account %s not found.", id)
}

func (b *Bank) Transfer(from string, to string, amount float32) error {
	transfer := NewTransferTransaction(amount, from, to)
	fromAccount, err := b.GetAccount(transfer.from)
	if err != nil {
		return err
	}

	err = fromAccount.Withdraw(WithdrawTransaction(transfer.Amount))

	if err != nil {
		return err
	}

	toAccount, err := b.GetAccount(transfer.to)

	if err != nil {
		return err
	}

	toAccount.Deposit(DepositTransaction(amount))

	return nil
}
