# Defer Statement in Go 🕐

**Directory**: `/defer` | **Difficulty**: ⭐⭐ Intermediate

This project demonstrates the `defer` statement in Go, which is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.

## 🎯 Learning Objectives

- Understand how the `defer` statement works
- Learn when deferred functions are executed
- Explore common use cases for defer
- Understand the LIFO (Last In, First Out) execution order

## 📁 Project Structure

```
defer/
├── README.md              # This documentation file
└── main.go               # Defer statement examples
```

## 🚀 Getting Started

### Prerequisites

Before starting this project, you should be familiar with:
- Basic Go syntax (see [Hello World Project](../hello-world/README.md))
- Functions and control flow
- Console output with `fmt` package

### Running the Project

```bash
cd defer
go run main.go
```

**Expected Output:**
```
Starting
Middle
Ending
This prints...
```

## 📚 Key Concepts

### What is `defer`?

The `defer` statement postpones the execution of a function until the surrounding function returns. The deferred function's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

### Key Characteristics

1. **LIFO Order**: Multiple defer statements are executed in Last In, First Out order
2. **Argument Evaluation**: Arguments are evaluated when the defer statement is encountered
3. **Function Scope**: Deferred functions are executed when the surrounding function returns
4. **Cleanup Pattern**: Commonly used for resource cleanup (file closing, mutex unlocking, etc.)

### Example Analysis

```go
func main() {
    fmt.Println("Starting")
    
    defer fmt.Println("This prints...") // Deferred until main() returns
    
    fmt.Println("Middle")
    fmt.Println("Ending")
}
```

**Execution Order:**
1. "Starting" is printed
2. The defer statement is encountered (but not executed yet)
3. "Middle" is printed
4. "Ending" is printed
5. `main()` function is about to return
6. The deferred function executes: "This prints..." is printed

## 🔍 Code Walkthrough

### Main Function Analysis

```go
package main

import "fmt"

func main() {
    fmt.Println("Starting")

    defer fmt.Println("This prints...") // When does this run?

    fmt.Println("Middle")
    fmt.Println("Ending")
}
```

**Step-by-step execution:**

1. **Line 6**: Prints "Starting" immediately
2. **Line 8**: The `defer` statement is encountered
   - The function call `fmt.Println("This prints...")` is deferred
   - It will execute when `main()` returns
3. **Line 10**: Prints "Middle" immediately
4. **Line 11**: Prints "Ending" immediately
5. **Function Return**: `main()` is about to return
6. **Deferred Execution**: The deferred function executes, printing "This prints..."

## 🎯 Common Use Cases

### 1. Resource Cleanup

```go
func processFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        return
    }
    defer file.Close() // Ensures file is closed when function returns
    
    // Process the file...
}
```

### 2. Mutex Unlocking

```go
func safeOperation() {
    mu.Lock()
    defer mu.Unlock() // Ensures mutex is unlocked when function returns
    
    // Perform operations...
}
```

### 3. Multiple Defer Statements

```go
func multipleDefers() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    defer fmt.Println("Third")
    
    fmt.Println("Main execution")
}
// Output: Main execution, Third, Second, First
```

## 🧪 Practice Exercises

### Exercise 1: Understanding Defer Order

Create a function that demonstrates the LIFO order of defer statements:

```go
func deferOrder() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")
    
    fmt.Println("Main execution")
}
```

### Exercise 2: Defer with Arguments

Demonstrate that arguments are evaluated when defer is encountered:

```go
func deferWithArgs() {
    x := 1
    defer fmt.Println("Deferred with x =", x)
    
    x = 2
    fmt.Println("Main execution with x =", x)
}
```

### Exercise 3: Defer in Loops

Explore defer behavior in loops:

```go
func deferInLoop() {
    for i := 1; i <= 3; i++ {
        defer fmt.Printf("Deferred: %d\n", i)
    }
    fmt.Println("Loop finished")
}
```

## 🔗 Related Projects

### Prerequisites
1. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax
2. **[Variables Project](../variables/README.md)** - Variable handling

### Next Steps
1. **[Error Handling Project](../error-handling/README.md)** - Error handling patterns
2. **[Calculator Project](../calculator/README.md)** - Practical application of defer

## 📖 Additional Resources

### Official Documentation
- [Go Defer Statement](https://golang.org/ref/spec#Defer_statements)
- [Effective Go - Defer](https://golang.org/doc/effective_go.html#defer)

### Related Concepts
- Function scope and lifetime
- Resource management
- Error handling patterns
- Cleanup patterns

## 🎯 Key Takeaways

- **Defer** postpones function execution until the surrounding function returns
- **LIFO order** applies to multiple defer statements
- **Arguments are evaluated immediately** when defer is encountered
- **Common use cases** include resource cleanup and mutex unlocking
- **Defer is essential** for proper resource management in Go

---

*Ready to explore error handling? Check out the [Error Handling project](../error-handling/README.md)!* 