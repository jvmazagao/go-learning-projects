# Structs - Go Structs and Methods 🏗️

A Go program that demonstrates structs, methods, and pointers in Go programming language.

## 📋 Table of Contents

1. [Overview](#overview)
2. [Files](#files)
3. [Code Structure](#code-structure)
4. [Key Concepts](#key-concepts)
5. [How to Run](#how-to-run)
6. [Expected Output](#expected-output)
7. [Learning Objectives](#learning-objectives)
8. [Prerequisites](#prerequisites)
9. [Next Steps](#next-steps)

## 🎯 Overview

This project demonstrates fundamental object-oriented concepts in Go using structs, methods, and pointers. It shows how to create custom data types, define methods on structs, and work with pointers for mutable operations.

Understanding structs is essential for building complex data structures and implementing object-oriented patterns in Go.

## 📁 Files

- `person.go` - The main Go source file demonstrating structs and methods

## 💻 Code Structure

```go
package main

import "fmt"

type Person struct {
	name  string
	age   int
	email string
}

func (p Person) greetings() string {
	return "Greetings " + p.name
}

func (p *Person) HaveBirthday() {
	p.age = p.age + 1
}

func NewPerson(name string, age int, email string) *Person {
	return &Person{name, age, email}
}

func main() {
	person := NewPerson("jhon", 10, "jvmazagao@gmail.com")
	fmt.Println(person.greetings())
	person.HaveBirthday()
	fmt.Println(person.age)
}
```

### Code Explanation

- **`type Person struct`**: Defines a custom data type with fields
- **`func (p Person) greetings()`**: Method with value receiver (read-only)
- **`func (p *Person) HaveBirthday()`**: Method with pointer receiver (can modify)
- **`NewPerson()`**: Constructor function that returns a pointer
- **`main()`**: Creates a person and demonstrates method calls

## 🔑 Key Concepts

### 1. Struct Definition
```go
type Person struct {
	name  string
	age   int
	email string
}
```
- Defines a custom data type with multiple fields
- Fields are accessed using dot notation: `person.name`

### 2. Value Receiver Methods
```go
func (p Person) greetings() string {
	return "Greetings " + p.name
}
```
- Method receives a copy of the struct
- Cannot modify the original struct
- Good for read-only operations

### 3. Pointer Receiver Methods
```go
func (p *Person) HaveBirthday() {
	p.age = p.age + 1
}
```
- Method receives a pointer to the struct
- Can modify the original struct
- Good for operations that change state

### 4. Constructor Function
```go
func NewPerson(name string, age int, email string) *Person {
	return &Person{name, age, email}
}
```
- Factory function to create new instances
- Returns a pointer for efficiency
- Follows Go conventions

## 🚀 How to Run

### Method 1: Direct Execution

1. Navigate to the structs directory:
   ```bash
   cd structs
   ```

2. Run the program:
   ```bash
   go run person.go
   ```

### Method 2: Build and Run

1. Build the executable:
   ```bash
   go build person.go
   ```

2. Run the executable:
   ```bash
   # On Unix-like systems (Linux, macOS)
   ./person
   
   # On Windows
   person.exe
   ```

## 📤 Expected Output

```
Greetings jhon
11
```

**Explanation**:
- First line: Output from the `greetings()` method
- Second line: Age after calling `HaveBirthday()` method (10 + 1 = 11)

## 🎓 Learning Objectives

By completing this project, you will understand:

- ✅ **Struct definition** - How to create custom data types in Go
- ✅ **Method receivers** - Value vs pointer receivers and when to use each
- ✅ **Pointer operations** - Working with pointers for mutable operations
- ✅ **Constructor patterns** - Factory functions for creating instances
- ✅ **Object-oriented concepts** - Encapsulation and method binding in Go

## 🔧 Prerequisites

Before starting this project, ensure you have:

- **Go installed** on your system (version 1.16 or later recommended)
- **Basic understanding** of Go syntax from previous projects
- **Understanding of variables** and basic types
- **Text editor or IDE** (VS Code, GoLand, Vim, etc.)

### Recommended Prerequisites

Complete these projects first:
1. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax
2. **[Variables Project](../variables/README.md)** - Variable handling
3. **[Calculator Project](../calculator/README.md)** - Functions and error handling

## 🔄 Next Steps

After understanding structs, explore:

### Advanced Struct Concepts
- **Struct composition** - Embedding structs within structs
- **Interface implementation** - Using structs with interfaces
- **Struct tags** - Adding metadata to struct fields
- **Anonymous structs** - Creating structs without names

### Related Projects in This Repository
1. **[Composition Project](../composition/README.md)** - Advanced struct composition
2. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax (prerequisite)
3. **[Variables Project](../variables/README.md)** - Variable handling (prerequisite)

### Additional Learning Resources
- [Go Tour - Structs](https://tour.golang.org/moretypes/2)
- [Go by Example - Structs](https://gobyexample.com/structs)
- [Effective Go - Structs](https://golang.org/doc/effective_go.html#structs)

## 🐛 Troubleshooting

### Common Issues

**Issue**: `cannot assign to field in map`
- **Solution**: Use pointer receivers for methods that modify struct fields

**Issue**: `undefined: field_name`
- **Solution**: Make sure field names are capitalized for public access

**Issue**: Method doesn't modify the struct
- **Solution**: Use pointer receiver `(p *Person)` instead of value receiver `(p Person)`

## 📚 Additional Notes

### Field Visibility
- **Capitalized fields**: Public (accessible from other packages)
- **Lowercase fields**: Private (only accessible within the same package)

### Method Naming Conventions
- Use descriptive names that indicate the action
- Follow Go naming conventions (camelCase)
- Be consistent with receiver naming

### Pointer vs Value Receivers
```go
// Use value receiver when:
// - Method doesn't modify the struct
// - Struct is small (copying is cheap)

// Use pointer receiver when:
// - Method modifies the struct
// - Struct is large (avoid copying)
// - Need to modify the original
```

### Struct Literals
```go
// Field names
person := Person{name: "John", age: 25, email: "john@example.com"}

// Positional (must be in order)
person := Person{"John", 25, "john@example.com"}

// Mixed
person := Person{name: "John", age: 25, email: "john@example.com"}
```

---

**🎉 Excellent! You now understand Go structs and methods!**

*Ready to explore composition? Check out the [Composition project](../composition/README.md)!* 