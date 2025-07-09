package src

import "testing"

func TestNewBank(t *testing.T) {
	b := NewBank()

	if b.accounts == nil {
		t.Error("Expected accounts map to be initialized")
	}
}

func TestCreateAndGetAccount(t *testing.T) {
	b := NewBank()

	acc := NewAccount("1", "owner", 100)

	b.CreateAccount("1", acc)

	got, err := b.GetAccount(acc.ID)

	if err != nil {
		t.Fatal("Failed to get account")
	}

	if got.ID != "1" {
		t.Errorf("ID = %v, want %v", got.ID, "1")
	}
	if got.Balance != 100.0 {
		t.Errorf("Balance = %v, want %v", got.Balance, 100.0)
	}
}

func TestTransfer(t *testing.T) {
	b := NewBank()

	acc1 := NewAccount("1", "Alice", 100.0)
	acc2 := NewAccount("2", "Bob", 50.0)
	b.CreateAccount("1", acc1)
	b.CreateAccount("2", acc2)

	err := b.Transfer("1", "2", 30.0)
	if err != nil {
		t.Fatalf("Transfer failed: %v", err)
	}

	a1, _ := b.GetAccount("1")
	a2, _ := b.GetAccount("2")

	if a1.Balance != 70.0 {
		t.Errorf("Alice balance = %v, want %v", a1.Balance, 70.0)
	}
	if a2.Balance != 80.0 {
		t.Errorf("Bob balance = %v, want %v", a2.Balance, 80.0)
	}
}

func TestInsufficientFunds(t *testing.T) {
	b := NewBank()

	acc1 := NewAccount("1", "Alice", 50.0)
	acc2 := NewAccount("2", "Bob", 0.0)
	b.CreateAccount("1", acc1)
	b.CreateAccount("2", acc2)

	err := b.Transfer("1", "2", 100.0)
	if err == nil {
		t.Error("Expected error for insufficient funds, got nil")
	}
}
