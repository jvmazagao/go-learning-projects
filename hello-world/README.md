# Hello World - Go Basics 🌟

A simple Go program that demonstrates the basic structure of a Go application and how to print output to the console.

## 📋 Table of Contents

1. [Overview](#overview)
2. [Files](#files)
3. [Code Structure](#code-structure)
4. [How to Run](#how-to-run)
5. [Expected Output](#expected-output)
6. [Learning Objectives](#learning-objectives)
7. [Prerequisites](#prerequisites)
8. [Next Steps](#next-steps)

## 🎯 Overview

This project contains a basic "Hello World" program written in Go. It serves as an introduction to Go programming fundamentals, including:

- Package declaration
- Import statements
- Main function
- Basic output using `fmt.Println()`

This is the perfect starting point for anyone new to Go programming!

## 📁 Files

- `hello.go` - The main Go source file containing the hello world program

## 💻 Code Structure

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello João")
    fmt.Println("Welcome to Go")
}
```

### Code Explanation

- **`package main`**: Declares this as the main package (entry point of the program)
- **`import "fmt"`**: Imports the formatting package for input/output operations
- **`func main()`**: The main function that gets executed when the program runs
- **`fmt.Println()`**: Prints text to the console with a newline at the end

## 🚀 How to Run

### Method 1: Direct Execution

1. Navigate to the hello-world directory:
   ```bash
   cd hello-world
   ```

2. Run the program:
   ```bash
   go run hello.go
   ```

### Method 2: Build and Run

1. Build the executable:
   ```bash
   go build hello.go
   ```

2. Run the executable:
   ```bash
   # On Unix-like systems (Linux, macOS)
   ./hello
   
   # On Windows
   hello.exe
   ```

## 📤 Expected Output

```
Hello João
Welcome to Go
```

## 🎓 Learning Objectives

By completing this project, you will understand:

- ✅ **Go package structure** - How packages are organized in Go
- ✅ **Import statements** - How to use external packages
- ✅ **Main function** - The entry point of Go programs
- ✅ **Console output** - How to print text to the terminal
- ✅ **Basic syntax** - Go's fundamental syntax rules

## 🔧 Prerequisites

Before starting this project, ensure you have:

- **Go installed** on your system (version 1.16 or later recommended)
- **Basic understanding** of command line operations
- **Text editor or IDE** (VS Code, GoLand, Vim, etc.)

### Verify Go Installation

```bash
go version
```

You should see output like: `go version go1.21.0 darwin/amd64`

## 🔄 Next Steps

After understanding this basic example, you can explore:

### Immediate Next Steps
- **Variables and data types** - Learn how to store and manipulate data
- **Functions and parameters** - Create reusable code blocks
- **Control structures** - Make decisions in your programs

### Related Projects in This Repository
1. **[Variables Project](../variables/README.md)** - Learn about variable declaration
2. **[Calculator Project](../calculator/README.md)** - Build a functional program

### Additional Learning Resources
- [Go Tour - Basics](https://tour.golang.org/basics/1)
- [Go by Example - Hello World](https://gobyexample.com/hello-world)
- [Effective Go - Packages](https://golang.org/doc/effective_go.html#package-names)

## 🐛 Troubleshooting

### Common Issues

**Issue**: `command not found: go`
- **Solution**: Install Go from [golang.org/dl](https://golang.org/dl/)

**Issue**: `cannot find package "fmt"`
- **Solution**: This usually means Go is not properly installed or configured

**Issue**: Permission denied when running executable
- **Solution**: On Unix-like systems, run `chmod +x hello`

## 📚 Additional Notes

- The `fmt` package is one of Go's most commonly used packages
- `Println` automatically adds a newline character at the end
- Use `Printf` for formatted output with placeholders
- The main function must be in the `main` package

---

**🎉 Congratulations! You've written your first Go program!**

*Ready for the next challenge? Check out the [Variables project](../variables/README.md)!* 