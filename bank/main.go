package main

import (
	"bank/src"
	"fmt"
)

func main() {
	b := src.NewBank()
	a_1 := src.NewAccount("1", "a_1")
	a_2 := src.NewAccount("2", "a_2")

	b.CreateAccount(a_1.ID, &a_1)
	b.CreateAccount(a_2.ID, &a_2)

	deposit := src.DepositTransaction(100)

	a_1.Deposit(deposit)
	a_2.Deposit(deposit)

	fmt.Println(a_1.GetBalance())
	fmt.Println(a_2.GetBalance())

	b.Transfer(a_1.ID, a_2.ID, 50)

	a1, _ := b.GetAccount(a_1.ID)
	a2, _ := b.GetAccount(a_2.ID)

	fmt.Println(a1.GetBalance())
	fmt.Println(a2.GetBalance())

}
