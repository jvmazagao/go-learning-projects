# Variables - Go Variable Declaration 📊

A Go program that demonstrates different ways to declare and use variables in Go programming language.

## 📋 Table of Contents

1. [Overview](#overview)
2. [Files](#files)
3. [Code Structure](#code-structure)
4. [Variable Declaration Examples](#variable-declaration-examples)
5. [How to Run](#how-to-run)
6. [Expected Output](#expected-output)
7. [Learning Objectives](#learning-objectives)
8. [Go Zero Values](#go-zero-values)
9. [Prerequisites](#prerequisites)
10. [Next Steps](#next-steps)

## 🎯 Overview

This project showcases various variable declaration techniques in Go, including:

- Explicit type declaration
- Variable initialization
- Multiple variable declaration
- Zero value behavior

Understanding variables is fundamental to Go programming and this project provides hands-on experience with different declaration methods.

## 📁 Files

- `variables.go` - The main Go source file demonstrating variable usage

## 💻 Code Structure

```go
package main

import "fmt"

func main() {
    var age int = 27
    var name string
    var x, y int = 10, 10

    fmt.Println(age)
    fmt.Println(name)
    fmt.Println(x, y)
}
```

### Code Explanation

- **`var age int = 27`**: Declares an integer variable with explicit type and initial value
- **`var name string`**: Declares a string variable without initialization (gets zero value)
- **`var x, y int = 10, 10`**: Declares multiple variables of the same type in one statement
- **`fmt.Println()`**: Prints each variable to see their values

## 📝 Variable Declaration Examples

### 1. Explicit Type Declaration with Initialization
```go
var age int = 27
```
- Declares an integer variable `age` with explicit type and initial value
- The `int` type can store whole numbers (positive, negative, or zero)

### 2. Zero Value Declaration
```go
var name string
```
- Declares a string variable `name` without initialization
- Go automatically assigns the zero value (empty string `""`)
- This is a key feature of Go's type system

### 3. Multiple Variable Declaration
```go
var x, y int = 10, 10
```
- Declares multiple variables of the same type in a single statement
- Both `x` and `y` are initialized to `10`

## 🚀 How to Run

### Method 1: Direct Execution

1. Navigate to the variables directory:
   ```bash
   cd variables
   ```

2. Run the program:
   ```bash
   go run variables.go
   ```

### Method 2: Build and Run

1. Build the executable:
   ```bash
   go build variables.go
   ```

2. Run the executable:
   ```bash
   # On Unix-like systems (Linux, macOS)
   ./variables
   
   # On Windows
   variables.exe
   ```

## 📤 Expected Output

```
27

10 10
```

**Note**: The empty line represents the zero value of the string variable `name` (an empty string).

## 🎓 Learning Objectives

By completing this project, you will understand:

- ✅ **Variable declaration syntax** - Different ways to declare variables in Go
- ✅ **Type system basics** - How Go's static typing works
- ✅ **Zero values** - What happens when variables aren't initialized
- ✅ **Multiple declarations** - How to declare multiple variables efficiently
- ✅ **Type safety** - How Go prevents type-related errors

## 🔢 Go Zero Values

In Go, variables declared without an explicit initial value are given their type's zero value:

| Type | Zero Value | Example |
|------|------------|---------|
| `int` | `0` | `var x int` → `x` is `0` |
| `string` | `""` (empty string) | `var s string` → `s` is `""` |
| `bool` | `false` | `var b bool` → `b` is `false` |
| `float64` | `0.0` | `var f float64` → `f` is `0.0` |
| `*T` | `nil` | `var p *int` → `p` is `nil` |

### Why Zero Values Matter

- **Predictable behavior**: You always know what a variable contains
- **No undefined values**: Unlike some languages, Go variables always have a value
- **Memory efficiency**: Zero values use minimal memory

## 🔧 Prerequisites

Before starting this project, ensure you have:

- **Go installed** on your system (version 1.16 or later recommended)
- **Basic understanding** of Go syntax from the hello-world example
- **Text editor or IDE** (VS Code, GoLand, Vim, etc.)

### Verify Go Installation

```bash
go version
```

You should see output like: `go version go1.21.0 darwin/amd64`

## 🔄 Next Steps

After understanding variables, explore:

### Immediate Next Steps
- **Constants** - Learn about immutable values
- **Type inference with `:=`** - Shorter variable declaration syntax
- **Variable scope** - Understanding where variables are accessible
- **Type conversions** - Converting between different types

### Related Projects in This Repository
1. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax (prerequisite)
2. **[Calculator Project](../calculator/README.md)** - Apply variables in a practical program

### Additional Learning Resources
- [Go Tour - Variables](https://tour.golang.org/basics/8)
- [Go by Example - Variables](https://gobyexample.com/variables)
- [Effective Go - Variables](https://golang.org/doc/effective_go.html#variables)

## 🐛 Troubleshooting

### Common Issues

**Issue**: `undefined: variable_name`
- **Solution**: Make sure the variable is declared before use

**Issue**: `cannot use string as int`
- **Solution**: Go is statically typed - you can't assign different types

**Issue**: Variable seems to have wrong value
- **Solution**: Check if you're seeing the zero value instead of an initialized value

## 📚 Additional Notes

### Variable Naming Conventions
- Use camelCase: `userName`, `totalCount`
- Start with a letter or underscore
- Avoid reserved keywords
- Be descriptive but concise

### Type Inference
```go
// These are equivalent:
var age int = 27
age := 27  // Type is inferred as int
```

### Block Declaration
```go
var (
    name string = "John"
    age  int    = 25
)
```

---

**🎉 Great job! You now understand Go variables!**

*Ready to build something practical? Check out the [Calculator project](../calculator/README.md)!* 