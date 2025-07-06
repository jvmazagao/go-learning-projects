# Error Handling in Go ⚠️

**Directory**: `/error-handling` | **Difficulty**: ⭐⭐ Intermediate

This project demonstrates Go's error handling patterns, which are fundamental to writing robust and reliable Go programs. Go uses explicit error handling rather than exceptions, making error handling a first-class concern.

## 🎯 Learning Objectives

- Understand Go's error handling philosophy
- Learn how to create and return errors
- Master error checking patterns
- Explore error handling best practices
- Understand the difference between errors and panics

## 📁 Project Structure

```
error-handling/
├── README.md              # This documentation file
└── main.go               # Error handling examples
```

## 🚀 Getting Started

### Prerequisites

Before starting this project, you should be familiar with:
- Basic Go syntax (see [Hello World Project](../hello-world/README.md))
- Functions and return values
- Control flow (if statements)
- The `fmt` package

### Running the Project

```bash
cd error-handling
go run main.go
```

**Expected Output:**
```
Result: 5
Error: Divided by 0
```

## 📚 Key Concepts

### Go's Error Philosophy

Go takes a different approach to error handling compared to many other languages:

1. **Explicit Error Handling**: Errors are returned as values, not thrown as exceptions
2. **No Exceptions**: Go doesn't have try/catch blocks
3. **Error as Interface**: Errors implement the `error` interface
4. **Multiple Return Values**: Functions can return both a result and an error

### The Error Interface

```go
type error interface {
    Error() string
}
```

Any type that implements the `Error() string` method satisfies the error interface.

### Error Handling Pattern

The standard pattern in Go is:

```go
result, err := someFunction()
if err != nil {
    // Handle the error
    return
}
// Use the result
```

## 🔍 Code Walkthrough

### Main Function Analysis

```go
package main

import "fmt"

func calculateDivision(number, divider float32) (float32, error) {
    if divider == 0 {
        return 0, fmt.Errorf("Divided by 0")
    }

    return number / divider, nil
}

func main() {
    result, err := calculateDivision(10, 2)

    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)

    // Try with zero
    result2, err2 := calculateDivision(10, 0)
    if err2 != nil {
        fmt.Println("Error:", err2)
        return
    }

    fmt.Println("Result:", result2)
}
```

**Step-by-step analysis:**

1. **Function Definition (Lines 5-11)**:
   - `calculateDivision` takes two `float32` parameters
   - Returns `(float32, error)` - multiple return values
   - Checks for division by zero
   - Returns error using `fmt.Errorf()` or success with `nil` error

2. **First Call (Lines 13-19)**:
   - Calls `calculateDivision(10, 2)`
   - Checks if `err != nil`
   - Since divider is not zero, no error occurs
   - Prints the result: `5`

3. **Second Call (Lines 21-27)**:
   - Calls `calculateDivision(10, 0)`
   - Division by zero triggers error
   - Error is checked and handled
   - Prints the error message

## 🎯 Error Handling Patterns

### 1. Creating Errors

```go
// Using fmt.Errorf
return fmt.Errorf("invalid input: %s", input)

// Using errors.New
import "errors"
return errors.New("something went wrong")

// Custom error types
type ValidationError struct {
    Field string
    Value string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Value)
}
```

### 2. Checking Errors

```go
// Basic error check
if err != nil {
    return err
}

// Error check with logging
if err != nil {
    log.Printf("operation failed: %v", err)
    return err
}

// Error check with custom handling
if err != nil {
    switch err.(type) {
    case *ValidationError:
        // Handle validation error
    default:
        // Handle other errors
    }
}
```

### 3. Wrapping Errors

```go
// Go 1.13+ error wrapping
import "fmt"

func processData(data string) error {
    if err := validate(data); err != nil {
        return fmt.Errorf("processing failed: %w", err)
    }
    return nil
}
```

## 🧪 Practice Exercises

### Exercise 1: Custom Error Types

Create a custom error type for different types of errors:

```go
type MathError struct {
    Operation string
    Message   string
}

func (e MathError) Error() string {
    return fmt.Sprintf("math error in %s: %s", e.Operation, e.Message)
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, MathError{
            Operation: "division",
            Message:   "division by zero",
        }
    }
    return a / b, nil
}
```

### Exercise 2: Error Handling with Defer

Combine error handling with defer for resource cleanup:

```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    // Process file...
    return nil
}
```

### Exercise 3: Multiple Error Checks

Handle multiple potential error conditions:

```go
func validateUser(name, email string) error {
    if name == "" {
        return fmt.Errorf("name cannot be empty")
    }
    
    if email == "" {
        return fmt.Errorf("email cannot be empty")
    }
    
    if !strings.Contains(email, "@") {
        return fmt.Errorf("invalid email format")
    }
    
    return nil
}
```

## 🔗 Related Projects

### Prerequisites
1. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax
2. **[Variables Project](../variables/README.md)** - Variable handling
3. **[Defer Project](../defer/README.md)** - Defer statements

### Next Steps
1. **[Calculator Project](../calculator/README.md)** - Practical error handling application
2. **[Structs Project](../structs/README.md)** - Object-oriented error handling

## 📖 Additional Resources

### Official Documentation
- [Go Error Handling](https://golang.org/doc/effective_go.html#errors)
- [Error Handling and Go](https://blog.golang.org/error-handling-and-go)
- [Working with Errors in Go 1.13](https://blog.golang.org/go1.13-errors)

### Best Practices
- Always check errors immediately after function calls
- Don't ignore errors (avoid `_` when error is returned)
- Provide meaningful error messages
- Use error wrapping for context
- Consider custom error types for complex scenarios

## 🎯 Key Takeaways

- **Errors are values** in Go, not exceptions
- **Always check errors** returned by functions
- **Use `fmt.Errorf()`** to create formatted error messages
- **Return `nil`** for success cases
- **Error handling is explicit** and part of the function signature
- **Combine with defer** for resource cleanup
- **Custom error types** provide more context and control

## 🚨 Common Mistakes

### ❌ Don't Ignore Errors
```go
// Bad
file, _ := os.Open("file.txt")

// Good
file, err := os.Open("file.txt")
if err != nil {
    return err
}
```

### ❌ Don't Return Errors from Main
```go
// Bad
func main() {
    err := doSomething()
    if err != nil {
        return err // This won't work
    }
}

// Good
func main() {
    err := doSomething()
    if err != nil {
        log.Fatal(err)
    }
}
```

### ❌ Don't Over-Engineer Simple Errors
```go
// Bad - over-engineered for simple case
type SimpleError struct {
    message string
}

func (e SimpleError) Error() string {
    return e.message
}

// Good - use standard library
return fmt.Errorf("simple error message")
```

---

*Ready to build something practical? Check out the [Calculator project](../calculator/README.md)!* 