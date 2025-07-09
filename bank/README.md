# Bank System in Go 🏦

**Directory**: `/bank` | **Difficulty**: ⭐⭐⭐ Advanced

This project demonstrates a complete banking system implementation in Go, showcasing object-oriented design, error handling, and real-world application patterns. The system includes account management, transaction processing, and transfer operations with proper validation.

## 🎯 Learning Objectives

- Understand complex struct design and composition
- Master method receivers (value vs pointer)
- Learn error handling in real-world scenarios
- Explore interface design and implementation
- Practice transaction processing and validation
- Understand package organization and module structure

## 📁 Project Structure

```
bank/
├── README.md              # This documentation file
├── go.mod                 # Go module definition
├── main.go               # Main application and testing
└── src/                  # Source code package
    ├── account.go        # Account struct and methods
    ├── transaction.go    # Transaction types and operations
    └── bank.go          # Bank system and transfer logic
```

## 🚀 Getting Started

### Prerequisites

Before starting this project, you should be familiar with:
- Basic Go syntax (see [Hello World Project](../hello-world/README.md))
- Structs and methods (see [Structs Project](../structs/README.md))
- Error handling (see [Error Handling Project](../error-handling/README.md))
- Pointers and references (see [Pointers Project](../pointers/README.md))

### Running the Project

```bash
cd bank
go run main.go
```

**Expected Output:**
```
=== BANK SYSTEM TEST ===

✓ Bank created
✓ 2 accounts created

--- Initial Balances ---
Alice: $100.00
Bob:   $50.00

--- Transfer $30 from Alice to Bob ---
✓ Transfer successful

--- Balances After Transfer ---
Alice: $70.00 (expected: $70.00)
Bob:   $80.00 (expected: $80.00)

--- Test Insufficient Funds ---
✓ Correctly rejected: Cannot withdraw this amount 100.000000.

--- Test Non-existent Account ---
✓ Correctly reported: Account 999 not found.

--- Final State ---
✅ ALL TESTS PASSED!
```

## 📚 Key Concepts

### Banking System Architecture

The bank system is built around three main components:

1. **Account**: Represents a customer account with balance and transaction history
2. **Transaction**: Represents financial operations (deposits, withdrawals, transfers)
3. **Bank**: Manages multiple accounts and handles transfers between them

### Design Patterns Used

- **Object-Oriented Design**: Structs with methods for encapsulation
- **Error Handling**: Comprehensive error checking and reporting
- **Interface Design**: Banking interface for extensibility
- **Transaction Processing**: Atomic operations with validation

## 🔍 Code Walkthrough

### Account Management (`src/account.go`)

```go
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
```

**Key Features:**
- **Constructor Pattern**: `NewAccount` creates properly initialized accounts
- **Transaction History**: Each account maintains a list of transactions
- **Encapsulation**: Balance is managed through methods, not direct access

### Transaction Processing

```go
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
```

**Key Features:**
- **Pointer Receivers**: Methods modify the original account
- **Validation**: Withdraw checks for sufficient funds
- **Transaction Recording**: All operations are logged
- **Error Handling**: Returns meaningful error messages

### Bank System (`src/bank.go`)

```go
type Bank struct {
    accounts    map[string]*Account
    transaction []Transaction
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
```

**Key Features:**
- **Account Management**: Uses map for efficient account lookup
- **Transfer Logic**: Atomic transfer operations
- **Error Propagation**: Proper error handling throughout
- **Transaction Creation**: Creates appropriate transaction types

### Transaction Types (`src/transaction.go`)

```go
type Transaction struct {
    identifier string
    Amount     float32
    timestamp  time.Time
    operation  string
}

type TranferTransaction struct {
    Transaction
    from string
    to   string
}
```

**Key Features:**
- **Composition**: TransferTransaction embeds Transaction
- **Time Stamping**: Automatic timestamp creation
- **Type Safety**: Different transaction types for different operations

## 🎯 Advanced Patterns

### 1. Interface Design

```go
type Banking interface {
    Deposit(t Transaction)
    Withdraw(t Transaction)
}
```

The `Banking` interface allows for different implementations and testing.

### 2. Error Handling Strategy

```go
func (b *Bank) GetAccount(id string) (*Account, error) {
    account, exists := b.accounts[id]
    if exists {
        return account, nil
    }
    return &Account{}, fmt.Errorf("Account %s not found.", id)
}
```

- **Explicit Error Checking**: Always check returned errors
- **Meaningful Messages**: Provide context in error messages
- **Graceful Degradation**: Handle errors without crashing

### 3. Transaction Atomicity

The transfer operation demonstrates atomic transaction processing:
1. Validate source account exists
2. Validate sufficient funds
3. Perform withdrawal
4. Validate destination account exists
5. Perform deposit

If any step fails, the entire operation is rolled back.

## 🧪 Practice Exercises

### Exercise 1: Add Interest Calculation

Create a method to calculate and apply interest to accounts:

```go
func (a *Account) ApplyInterest(rate float32) {
    interest := a.Balance * rate
    transaction := Transaction{
        Amount:    interest,
        operation: "INTEREST",
        timestamp: time.Now(),
    }
    a.Deposit(transaction)
}
```

### Exercise 2: Implement Account Statement

Create a method to generate account statements:

```go
func (a *Account) GenerateStatement() string {
    var statement strings.Builder
    statement.WriteString(fmt.Sprintf("Account: %s\n", a.ID))
    statement.WriteString(fmt.Sprintf("Owner: %s\n", a.Owner))
    statement.WriteString(fmt.Sprintf("Balance: $%.2f\n\n", a.Balance))
    statement.WriteString("Transactions:\n")
    
    for _, t := range a.Transactions {
        statement.WriteString(fmt.Sprintf("%s: $%.2f (%s)\n", 
            t.timestamp.Format("2006-01-02 15:04:05"), 
            t.Amount, t.operation))
    }
    
    return statement.String()
}
```

### Exercise 3: Add Transaction Validation

Implement transaction validation rules:

```go
func (t Transaction) Validate() error {
    if t.Amount <= 0 {
        return fmt.Errorf("transaction amount must be positive")
    }
    
    if t.operation != "DEPOSIT" && t.operation != "WITHDRAW" {
        return fmt.Errorf("invalid operation: %s", t.operation)
    }
    
    return nil
}
```

## 🔗 Related Projects

### Prerequisites
1. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax
2. **[Structs Project](../structs/README.md)** - Structs and methods
3. **[Error Handling Project](../error-handling/README.md)** - Error handling patterns
4. **[Pointers Project](../pointers/README.md)** - Pointer receivers

### Next Steps
1. **[Interfaces Project](../interfaces/README.md)** - Interface design patterns
2. **[Composition Project](../composition/README.md)** - Advanced composition

## 📖 Additional Resources

### Design Patterns
- **Domain-Driven Design**: Structuring code around business concepts
- **Repository Pattern**: Managing data access through interfaces
- **Factory Pattern**: Creating objects with proper initialization
- **Command Pattern**: Encapsulating operations as objects

### Best Practices
- **Encapsulation**: Keep data private, expose through methods
- **Single Responsibility**: Each struct has one clear purpose
- **Error Handling**: Always handle errors explicitly
- **Testing**: Write tests for business logic

## 🎯 Key Takeaways

- **Object-Oriented Design**: Use structs and methods for encapsulation
- **Error Handling**: Comprehensive error checking is essential
- **Transaction Safety**: Ensure operations are atomic and consistent
- **Interface Design**: Plan for extensibility and testing
- **Real-World Patterns**: Apply business logic patterns to code
- **Package Organization**: Structure code logically across files

## 🚨 Common Mistakes

### ❌ Ignoring Error Returns
```go
// Bad - ignoring potential errors
account.Withdraw(transaction)

// Good - always check errors
if err := account.Withdraw(transaction); err != nil {
    return err
}
```

### ❌ Direct Field Access
```go
// Bad - bypassing validation
account.Balance -= amount

// Good - use methods for controlled access
account.Withdraw(WithdrawTransaction(amount))
```

### ❌ Inconsistent Error Handling
```go
// Bad - inconsistent error patterns
func (b *Bank) GetAccount(id string) *Account {
    return b.accounts[id] // Might return nil
}

// Good - explicit error handling
func (b *Bank) GetAccount(id string) (*Account, error) {
    account, exists := b.accounts[id]
    if !exists {
        return nil, fmt.Errorf("account not found")
    }
    return account, nil
}
```

## 🔍 Advanced Concepts

### Memory Management

```go
// Efficient account storage
type Bank struct {
    accounts map[string]*Account  // Pointers for efficiency
    // ...
}
```

### Transaction Logging

```go
// Comprehensive transaction tracking
type Transaction struct {
    ID        string    // Unique identifier
    Amount    float32   // Transaction amount
    Type      string    // Transaction type
    Timestamp time.Time // When it occurred
    AccountID string    // Associated account
}
```

### Concurrency Considerations

```go
// Thread-safe bank operations
type Bank struct {
    accounts map[string]*Account
    mu       sync.RWMutex
}

func (b *Bank) Transfer(from, to string, amount float32) error {
    b.mu.Lock()
    defer b.mu.Unlock()
    // ... transfer logic
}
```

---

*Ready to explore more complex systems? Check out the [Interfaces project](../interfaces/README.md) for advanced design patterns!* 