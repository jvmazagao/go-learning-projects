# Interfaces - Go Interfaces and Polymorphism 🔌

A Go program that demonstrates interfaces, type assertions, and polymorphism in Go programming language.

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

This project demonstrates fundamental interface concepts in Go, including interface definition, implementation, type assertions, and polymorphism. It shows how to create flexible, reusable code using interfaces and how to work with different types through a common interface.

Understanding interfaces is crucial for writing flexible and maintainable Go code.

## 📁 Files

- `interfaces.go` - The main Go source file demonstrating interfaces and polymorphism

## 💻 Code Structure

```go
package main

import "fmt"

type Greetings interface {
	Greets() string
}

type Naming struct {
	name string
}

type Person struct {
	Naming
}

func (p Person) Greets() string {
	return fmt.Sprintf("Hi! I'm %s", p.name)
}

type Dog struct {
	Naming
}

func (d Dog) Greets() string {
	return fmt.Sprintf("Woof, woof!")
}

type Robot struct {
	Naming
}

func (r Robot) Greets() string {
	return fmt.Sprintf("Beep! Boop! I'm %s", r.name)
}

func MakeGreetings(g Greetings) string {
	return g.Greets()
}

type empty interface{}

func PrintAnything(thing empty) {
	fmt.Println("I received:", thing)
}

func WhatThingIsIt(thing interface{}) {
	switch t := thing.(type) {
	case string:
		fmt.Println("this is string:", t)
	case int:
		fmt.Println("this is a integer:", t)
	case bool:
		fmt.Println("this is a boolean:", t)
	default:
		fmt.Println("this is a default value:", t)
	}
}

func main() {
	person := Person{Naming{"Jhon"}}
	dog := Dog{Naming{"Vilma"}}
	robot := Robot{Naming{"Bender"}}
	fmt.Println(MakeGreetings(person))
	fmt.Println(MakeGreetings(dog))
	fmt.Println(MakeGreetings(robot))

	mixedTypeBag := []interface{}{3, "any", true}

	for _, item := range mixedTypeBag {
		WhatThingIsIt(item)
	}
}
```

### Code Explanation

- **`Greetings interface`**: Defines a contract for types that can greet
- **`Person`, `Dog`, `Robot`**: Different types that implement the `Greetings` interface
- **`MakeGreetings()`**: Function that works with any type implementing `Greetings`
- **`empty interface{}`**: Interface that accepts any type
- **`WhatThingIsIt()`**: Demonstrates type assertions with switch statements

## 🔑 Key Concepts

### 1. Interface Definition
```go
type Greetings interface {
	Greets() string
}
```
- Defines a contract that types must implement
- Only specifies method signatures, not implementations
- Any type with a `Greets() string` method implements this interface

### 2. Interface Implementation
```go
func (p Person) Greets() string {
	return fmt.Sprintf("Hi! I'm %s", p.name)
}

func (d Dog) Greets() string {
	return fmt.Sprintf("Woof, woof!")
}

func (r Robot) Greets() string {
	return fmt.Sprintf("Beep! Boop! I'm %s", r.name)
}
```
- Each type implements the interface in its own way
- Implementation is implicit (no explicit declaration needed)
- Different behaviors for the same interface method

### 3. Polymorphism
```go
func MakeGreetings(g Greetings) string {
	return g.Greets()
}
```
- Function accepts any type that implements `Greetings`
- Same function works with `Person`, `Dog`, or `Robot`
- Demonstrates polymorphic behavior

### 4. Empty Interface
```go
type empty interface{}

func PrintAnything(thing empty) {
	fmt.Println("I received:", thing)
}
```
- `interface{}` accepts any type
- Useful for functions that need to handle unknown types
- Also called the "empty interface"

### 5. Type Assertions
```go
func WhatThingIsIt(thing interface{}) {
	switch t := thing.(type) {
	case string:
		fmt.Println("this is string:", t)
	case int:
		fmt.Println("this is a integer:", t)
	case bool:
		fmt.Println("this is a boolean:", t)
	default:
		fmt.Println("this is a default value:", t)
	}
}
```
- Uses type switch to determine the actual type
- Safe way to work with `interface{}` values
- Handles different types appropriately

## 🚀 How to Run

### Method 1: Direct Execution

1. Navigate to the interfaces directory:
   ```bash
   cd interfaces
   ```

2. Run the program:
   ```bash
   go run interfaces.go
   ```

### Method 2: Build and Run

1. Build the executable:
   ```bash
   go build interfaces.go
   ```

2. Run the executable:
   ```bash
   # On Unix-like systems (Linux, macOS)
   ./interfaces
   
   # On Windows
   interfaces.exe
   ```

## 📤 Expected Output

```
Hi! I'm Jhon
Woof, woof!
Beep! Boop! I'm Bender
this is a integer: 3
this is string: any
this is a boolean: true
```

**Explanation**:
- First three lines: Polymorphic behavior with different types implementing `Greetings`
- Last three lines: Type assertions showing different types in the mixed slice

## 🎓 Learning Objectives

By completing this project, you will understand:

- ✅ **Interface definition** - How to define contracts for types
- ✅ **Interface implementation** - How types implicitly implement interfaces
- ✅ **Polymorphism** - How interfaces enable polymorphic behavior
- ✅ **Empty interfaces** - How to work with any type using `interface{}`
- ✅ **Type assertions** - How to safely work with interface values
- ✅ **Type switches** - How to handle different types in a switch statement

## 🔧 Prerequisites

Before starting this project, ensure you have:

- **Go installed** on your system (version 1.16 or later recommended)
- **Understanding of structs** and methods from previous projects
- **Basic knowledge** of composition and embedding
- **Text editor or IDE** (VS Code, GoLand, Vim, etc.)

### Recommended Prerequisites

Complete these projects first:
1. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax
2. **[Variables Project](../variables/README.md)** - Variable handling
3. **[Structs Project](../structs/README.md)** - Structs and methods
4. **[Composition Project](../composition/README.md)** - Struct composition

## 🔄 Next Steps

After understanding interfaces, explore:

### Advanced Interface Concepts
- **Interface composition** - Embedding interfaces within interfaces
- **Interface satisfaction** - Ensuring types implement all required methods
- **Interface design** - Best practices for interface design
- **Mocking** - Using interfaces for testing

### Related Projects in This Repository
1. **[Player-Recorder Project](../player-recorder/README.md)** - Advanced interface composition
2. **[Structs Project](../structs/README.md)** - Basic structs and methods (prerequisite)
3. **[Composition Project](../composition/README.md)** - Struct composition (prerequisite)

### Additional Learning Resources
- [Go Tour - Interfaces](https://tour.golang.org/methods/9)
- [Go by Example - Interfaces](https://gobyexample.com/interfaces)
- [Effective Go - Interfaces](https://golang.org/doc/effective_go.html#interfaces)

## 🐛 Troubleshooting

### Common Issues

**Issue**: `cannot use type as interface type`
- **Solution**: Make sure the type implements all methods in the interface

**Issue**: `interface conversion: interface is nil`
- **Solution**: Check if the interface value is nil before type assertion

**Issue**: `impossible type assertion`
- **Solution**: The type assertion is impossible because the types don't match

## 📚 Additional Notes

### Interface Design Principles

**Keep interfaces small**:
```go
// Good - focused interface
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Bad - too many methods
type BigInterface interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
    Close() error
    Flush() error
    // ... many more methods
}
```

### Interface Satisfaction

Interfaces are satisfied implicitly:
```go
type Greetings interface {
    Greets() string
}

// Person automatically implements Greetings
// No explicit declaration needed
type Person struct {
    name string
}

func (p Person) Greets() string {
    return "Hello, " + p.name
}
```

### Type Assertions vs Type Switches

```go
// Type assertion (single type)
if str, ok := thing.(string); ok {
    fmt.Println("It's a string:", str)
}

// Type switch (multiple types)
switch t := thing.(type) {
case string:
    fmt.Println("String:", t)
case int:
    fmt.Println("Int:", t)
default:
    fmt.Println("Unknown type")
}
```

---

**🎉 Excellent! You now understand Go interfaces and polymorphism!**

*Ready for advanced interface concepts? Check out the [Player-Recorder project](../player-recorder/README.md)!* 