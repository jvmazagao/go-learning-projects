# Calculator Command-Line Tool 🧮

A comprehensive command-line calculator written in Go that demonstrates advanced Go concepts including error handling, command-line argument parsing, and unit testing.

## 📋 Table of Contents

1. [Overview](#overview)
2. [Features](#features)
3. [Files](#files)
4. [Installation](#installation)
5. [Usage](#usage)
6. [Error Handling](#error-handling)
7. [Code Breakdown](#code-breakdown)
8. [Testing](#testing)
9. [Prerequisites](#prerequisites)
10. [Next Steps](#next-steps)

## 🎯 Overview

This is a sophisticated command-line calculator written in Go that takes two integers and an arithmetic operator as input and returns the result of the corresponding operation. It supports addition, subtraction, multiplication, and division operations with comprehensive error handling using panic and recover mechanisms.

This project demonstrates real-world Go programming concepts and is perfect for understanding how to build practical applications.

## ✨ Features

- **Four basic arithmetic operations**: 
  - Addition (`+`)
  - Subtraction (`-`)
  - Multiplication (`*`)
  - Division (`/`)
- **Robust error handling** for invalid inputs and edge cases
- **User-friendly output** with clear formatting
- **Comprehensive unit tests** for all functionality
- **Command-line interface** for easy integration with scripts

## 📁 Files

- `calculator.go` - Main calculator program with all logic
- `calculator_test.go` - Comprehensive unit tests
- `go.mod` - Go module definition
- `README.md` - This documentation file

## 🔧 Installation

### Prerequisites

To get started with the calculator, you'll need to have Go installed on your machine. You can download and install Go from [the official Go website](https://golang.org/dl/).

### Setup

1. **Clone the repository** (if not already done):
   ```bash
   git clone https://github.com/your-username/go-learn.git
   cd go-learn/calculator
   ```

2. **Verify Go installation**:
   ```bash
   go version
   ```

3. **Install dependencies** (if any):
   ```bash
   go mod tidy
   ```

## 🚀 Usage

### Command Format

```bash
go run calculator.go <num1> <operator> <num2>
```

### Example Usage

#### 1. Addition
```bash
go run calculator.go 5 + 3
```
**Output**:
```
The result from (5 + 3) is 8
```

#### 2. Subtraction
```bash
go run calculator.go 7 - 2
```
**Output**:
```
The result from (7 - 2) is 5
```

#### 3. Multiplication
```bash
go run calculator.go 6 * 4
```
**Output**:
```
The result from (6 * 4) is 24
```

#### 4. Division
```bash
go run calculator.go 10 / 2
```
**Output**:
```
The result from (10 / 2) is 5
```

### Error Examples

#### Division by Zero
```bash
go run calculator.go 10 / 0
```
**Output**:
```
💥 Error: division by zero
```

#### Invalid Number
```bash
go run calculator.go abc + 5
```
**Output**:
```
💥 Error: failed to convert number
```

#### Missing Arguments
```bash
go run calculator.go 5 +
```
**Output**:
```
💥 Error: Please provide a number
```

## ⚠️ Error Handling

The program uses **panic** and **recover** to handle errors gracefully:

### Error Types Handled

- **Invalid input conversion**: Non-numeric values trigger `"failed to convert number"`
- **Division by zero**: Attempts to divide by zero trigger `"division by zero"`
- **Unsupported operations**: Invalid operators trigger `"operation not permitted"`
- **Missing arguments**: Insufficient command-line arguments trigger `"Please provide a number"`

### Error Recovery

The `GlobalErrorHandler` function ensures that any panics are caught and displayed in a user-friendly way, and the program exits cleanly with an error code.

## 💻 Code Breakdown

### 1. Calculator Function
Performs the specified operation (`+`, `-`, `*`, `/`) on the two numbers provided.

```go
func Calculator(a int, b int, operation string) int {
	switch operation {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("division by zero")
		}
		return a / b
	default:
		panic("operation not permitted")
	}
}
```

**Key Features**:
- Uses `switch` statement for clean operation handling
- Validates division by zero before performing operation
- Panics for unsupported operations

### 2. ConvertNumber Function
Converts a string to an integer with error handling.

```go
func ConvertNumber(num string) int {
	res, err := strconv.Atoi(num)
	if err != nil {
		panic("failed to convert number")
	}
	return res
}
```

**Key Features**:
- Uses `strconv.Atoi()` for string-to-int conversion
- Handles conversion errors with panic
- Returns integer value on success

### 3. GlobalErrorHandler Function
A global error handler that recovers from panics and provides user-friendly error messages.

```go
func GlobalErrorHandler() {
	if r := recover(); r != nil {
		fmt.Fprintln(os.Stderr, "💥 Error:", r)
		os.Exit(1)
	}
}
```

**Key Features**:
- Uses `defer` to ensure it runs even if main function panics
- Writes errors to stderr for proper error stream handling
- Exits with error code 1 to indicate failure

### 4. Main Function
The entry point that orchestrates the entire program.

```go
func main() {
	defer GlobalErrorHandler()

	if len(os.Args) < 4 {
		panic("Please provide a number")
	}

	num1 := ConvertNumber(os.Args[1])
	num2 := ConvertNumber(os.Args[3])
	operator := os.Args[2]

	result := Calculator(num1, num2, operator)

	fmt.Printf("The result from (%d %s %d) is %d \n", num1, operator, num2, result)
}
```

**Key Features**:
- Validates command-line argument count
- Parses arguments in correct order
- Uses formatted output for clear results

## 🧪 Testing

The project includes comprehensive unit tests in `calculator_test.go`.

### Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover

# Run tests with coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Test Coverage

The tests cover:
- All arithmetic operations
- Error conditions (division by zero, invalid operations)
- Input validation
- Edge cases

## 🔧 Prerequisites

Before starting this project, ensure you have:

- **Go installed** on your system (version 1.16 or later recommended)
- **Basic understanding** of Go syntax from previous projects
- **Command line experience** for running the calculator
- **Text editor or IDE** (VS Code, GoLand, Vim, etc.)

### Recommended Prerequisites

Complete these projects first:
1. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax
2. **[Variables Project](../variables/README.md)** - Variable handling

## 🔄 Next Steps

After mastering this calculator, explore:

### Advanced Go Concepts
- **Structs and methods** - Object-oriented programming in Go
- **Interfaces** - Polymorphism and abstraction
- **Goroutines and channels** - Concurrency in Go
- **Web development** - Building HTTP servers and APIs

### Related Projects in This Repository
1. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax (prerequisite)
2. **[Variables Project](../variables/README.md)** - Variable handling (prerequisite)
3. **[Structs Project](../structs/README.md)** - Object-oriented concepts

### Additional Learning Resources
- [Go Tour - Methods](https://tour.golang.org/methods/1)
- [Go by Example - Functions](https://gobyexample.com/functions)
- [Effective Go - Error Handling](https://golang.org/doc/effective_go.html#errors)

## 🐛 Troubleshooting

### Common Issues

**Issue**: `command not found: go`
- **Solution**: Install Go from [golang.org/dl](https://golang.org/dl/)

**Issue**: `cannot find package "strconv"`
- **Solution**: This is a standard library package - check your Go installation

**Issue**: Program exits immediately
- **Solution**: Make sure you're providing all three arguments: `go run calculator.go 5 + 3`

**Issue**: Tests fail
- **Solution**: Run `go test -v` to see detailed test output

## 📚 Additional Notes

### Command-Line Arguments
- `os.Args[0]` is the program name
- `os.Args[1]` is the first number
- `os.Args[2]` is the operator
- `os.Args[3]` is the second number

### Error Handling Philosophy
- Use panic for unrecoverable errors
- Use recover to handle panics gracefully
- Provide clear, user-friendly error messages

### Testing Best Practices
- Test both success and failure cases
- Use descriptive test names
- Test edge cases and boundary conditions

---

**🎉 Congratulations! You've built a functional calculator in Go!**

*You now understand error handling, testing, and building practical Go applications!* 