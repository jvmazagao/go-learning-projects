# Composition - Go Struct Composition 🔗

A Go program that demonstrates struct composition and embedding, showing how to build complex data structures by combining simpler ones.

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

This project demonstrates advanced object-oriented concepts in Go using struct composition and embedding. It shows how to create complex data structures by embedding simpler structs, implement method overriding, and build hierarchical relationships between types.

Struct composition is Go's way of achieving code reuse and building complex abstractions without traditional inheritance.

## 📁 Files

- `employee.go` - The main Go source file demonstrating struct composition

## 💻 Code Structure

```go
package main

import "fmt"

type Address struct {
	street string
	city   string
}

type Person struct {
	Address
	name  string
	age   int
	email string
}

func (p Person) Greetings() string {
	return fmt.Sprintf("Greetings %s", p.name)
}

func NewPerson(name string, age int, email string) *Person {
	return &Person{Address{"bla", "ble"}, name, age, email}
}

type Employee struct {
	Person
	jobTitle    string
	salary      float32
	workAddress Address
}

func (e Employee) Greetings() string {
	return fmt.Sprintf("Hello, I'm %s and I work as a %s", e.name, e.jobTitle)
}

func main() {
	employee := Employee{
		Person:   *NewPerson("jhon", 31, "jhon@gmail.com"),
		jobTitle: "master",
		salary:   100.00,
	}
	fmt.Println(employee.Greetings())
	fmt.Println(employee.Person.Greetings())
}
```

### Code Explanation

- **`Address struct`**: Simple struct representing a physical address
- **`Person struct`**: Embeds `Address` and adds personal information
- **`Employee struct`**: Embeds `Person` and adds work-related information
- **Method overriding**: `Employee.Greetings()` overrides `Person.Greetings()`
- **Composition**: Building complex types from simpler ones

## 🔑 Key Concepts

### 1. Struct Embedding
```go
type Person struct {
	Address  // Embedded struct (composition)
	name  string
	age   int
	email string
}
```
- `Address` is embedded directly (no field name)
- All `Address` fields become accessible on `Person`
- This is Go's way of achieving composition over inheritance

### 2. Method Promotion
```go
// Address fields are promoted to Person
person := Person{Address{"123 Main St", "City"}, "John", 25, "john@example.com"}
fmt.Println(person.street)  // Access embedded field directly
fmt.Println(person.city)    // Access embedded field directly
```
- Embedded struct's fields and methods are promoted to the outer struct
- Can access embedded fields directly without qualification

### 3. Method Overriding
```go
func (p Person) Greetings() string {
	return fmt.Sprintf("Greetings %s", p.name)
}

func (e Employee) Greetings() string {
	return fmt.Sprintf("Hello, I'm %s and I work as a %s", e.name, e.jobTitle)
}
```
- `Employee.Greetings()` overrides `Person.Greetings()`
- Can still access the original method via `employee.Person.Greetings()`

### 4. Multiple Embedding
```go
type Employee struct {
	Person        // Embedded Person (includes Address)
	jobTitle    string
	salary      float32
	workAddress Address  // Separate Address field
}
```
- `Employee` embeds `Person` (which embeds `Address`)
- Also has its own `workAddress` field
- Demonstrates hierarchical composition

## 🚀 How to Run

### Method 1: Direct Execution

1. Navigate to the composition directory:
   ```bash
   cd composition
   ```

2. Run the program:
   ```bash
   go run employee.go
   ```

### Method 2: Build and Run

1. Build the executable:
   ```bash
   go build employee.go
   ```

2. Run the executable:
   ```bash
   # On Unix-like systems (Linux, macOS)
   ./employee
   
   # On Windows
   employee.exe
   ```

## 📤 Expected Output

```
Hello, I'm jhon and I work as a master
Greetings jhon
```

**Explanation**:
- First line: Output from `Employee.Greetings()` (overridden method)
- Second line: Output from `Person.Greetings()` (original method accessed via composition)

## 🎓 Learning Objectives

By completing this project, you will understand:

- ✅ **Struct embedding** - How to embed one struct within another
- ✅ **Method promotion** - How embedded methods become available on the outer struct
- ✅ **Method overriding** - How to override methods from embedded structs
- ✅ **Composition patterns** - Building complex types from simpler ones
- ✅ **Hierarchical relationships** - Creating type hierarchies without inheritance

## 🔧 Prerequisites

Before starting this project, ensure you have:

- **Go installed** on your system (version 1.16 or later recommended)
- **Understanding of structs** and methods from the structs project
- **Basic knowledge** of pointers and constructors
- **Text editor or IDE** (VS Code, GoLand, Vim, etc.)

### Recommended Prerequisites

Complete these projects first:
1. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax
2. **[Variables Project](../variables/README.md)** - Variable handling
3. **[Structs Project](../structs/README.md)** - Structs and methods (essential prerequisite)

## 🔄 Next Steps

After understanding composition, explore:

### Advanced Composition Concepts
- **Interface composition** - Embedding interfaces within interfaces
- **Type assertions** - Working with embedded types
- **Polymorphism** - Using composition for polymorphic behavior
- **Design patterns** - Implementing common patterns with composition

### Related Projects in This Repository
1. **[Structs Project](../structs/README.md)** - Basic structs and methods (prerequisite)
2. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax (prerequisite)
3. **[Variables Project](../variables/README.md)** - Variable handling (prerequisite)

### Additional Learning Resources
- [Go Tour - Struct Embedding](https://tour.golang.org/moretypes/10)
- [Go by Example - Structs](https://gobyexample.com/structs)
- [Effective Go - Embedding](https://golang.org/doc/effective_go.html#embedding)

## 🐛 Troubleshooting

### Common Issues

**Issue**: `ambiguous selector`
- **Solution**: When multiple embedded structs have the same field/method, you must qualify the access

**Issue**: `cannot use embedded type as field`
- **Solution**: Embedded types must be struct types, not interfaces or other types

**Issue**: Method not found on embedded struct
- **Solution**: Make sure the embedded struct has the method and it's properly promoted

## 📚 Additional Notes

### Composition vs Inheritance

**Go uses composition over inheritance**:
- **Composition**: "has-a" relationship (Employee has a Person)
- **Inheritance**: "is-a" relationship (not directly supported in Go)
- **Benefits**: More flexible, avoids inheritance pitfalls

### Field Access Rules

```go
type Employee struct {
	Person
	workAddress Address
}

employee := Employee{...}

// These are equivalent:
employee.street        // Promoted from Person.Address
employee.Person.street // Explicit access

// These are different:
employee.street        // Home address (from Person)
employee.workAddress.street // Work address
```

### Method Resolution

When a method is called on a struct with embedded types:
1. Look for method on the struct itself
2. Look for method on embedded types (in order of embedding)
3. If multiple embedded types have the same method, it's ambiguous

### Best Practices

- **Use composition for code reuse**
- **Keep embedded structs simple and focused**
- **Be careful with method name conflicts**
- **Use explicit qualification when needed**
- **Document the composition relationships**

---

**🎉 Fantastic! You now understand Go struct composition!**

*You've mastered the fundamentals of object-oriented programming in Go!* 