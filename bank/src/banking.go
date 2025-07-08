package src

import (
	"fmt"
	"time"
)

type Bank struct {
	accounts map[string]*Account
}

func (b Bank) CreateAccount(id string, account Account) {
	b.accounts[id] = &account
}

func (b Bank) GetAccount(id string) (Account, error) {
	account, exists := b.accounts[id]
	if exists {
		return *account, nil
	}
	return Account{}, fmt.Errorf("Account %s not found.", id)
}

func (b Bank) Transfer(from string, to string, amount float32) error {
	fromAccount, err := b.GetAccount(from)
	if err != nil {
		return fmt.Errorf("Account %s not foud.", from)
	}

	toAccount, err := b.GetAccount(to)
	if err != nil {
		return fmt.Errorf("Account %s not found.", to)
	}

	err = fromAccount.Withdraw(Transaction{
		amount:    amount,
		timestamp: time.Now(),
		from:      from,
		to:        to,
		operation: "WITHDRAW",
	})

	if err != nil {
		return err
	}

	toAccount.Deposit(Transaction{
		amount:    amount,
		timestamp: time.Now(),
		from:      from,
		to:        to,
		operation: "DEPOSIT",
	})

	return nil
}
