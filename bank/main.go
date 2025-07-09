package main

import (
	"bank/src"
	"fmt"
)

func main() {
	fmt.Println("=== BANK SYSTEM TEST ===\n")

	// Create bank
	b := src.NewBank()
	fmt.Println("✓ Bank created")

	// Create accounts
	acc1 := src.NewAccount("1", "Alice", 100.0)
	acc2 := src.NewAccount("2", "Bob", 50.0)

	b.CreateAccount("1", acc1)
	b.CreateAccount("2", acc2)
	fmt.Println("✓ 2 accounts created")

	// Check initial balances
	fmt.Println("\n--- Initial Balances ---")
	a1, err1 := b.GetAccount("1")
	a2, err2 := b.GetAccount("2")

	if err1 != nil || err2 != nil {
		fmt.Println("ERROR: Could not get accounts")
		return
	}

	fmt.Printf("Alice: $%.2f\n", a1.Balance)
	fmt.Printf("Bob:   $%.2f\n", a2.Balance)

	// Test transfer
	fmt.Println("\n--- Transfer $30 from Alice to Bob ---")
	err := b.Transfer("1", "2", 30.0)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("✓ Transfer successful")
	}

	// Check balances after transfer
	fmt.Println("\n--- Balances After Transfer ---")
	a1, _ = b.GetAccount("1")
	a2, _ = b.GetAccount("2")
	fmt.Printf("Alice: $%.2f (expected: $70.00)\n", a1.Balance)
	fmt.Printf("Bob:   $%.2f (expected: $80.00)\n", a2.Balance)

	// Test insufficient funds
	fmt.Println("\n--- Test Insufficient Funds ---")
	err = b.Transfer("1", "2", 100.0)
	if err != nil {
		fmt.Println("✓ Correctly rejected:", err)
	} else {
		fmt.Println("ERROR: Should have failed!")
	}

	// Test non-existent account
	fmt.Println("\n--- Test Non-existent Account ---")
	_, err = b.GetAccount("999")
	if err != nil {
		fmt.Println("✓ Correctly reported:", err)
	} else {
		fmt.Println("ERROR: Should have failed!")
	}

	// Final check
	fmt.Println("\n--- Final State ---")
	success := a1.Balance == 70.0 && a2.Balance == 80.0
	if success {
		fmt.Println("✅ ALL TESTS PASSED!")
	} else {
		fmt.Println("❌ BALANCE MISMATCH!")
	}
}
