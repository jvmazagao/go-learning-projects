# Slices in Go 📏

**Directory**: `/slices` | **Difficulty**: ⭐⭐ Intermediate

This project demonstrates Go's slice data structure, which provides a dynamic, flexible view into arrays. Slices are one of Go's most commonly used data structures, offering dynamic sizing and efficient operations.

## 🎯 Learning Objectives

- Understand the relationship between arrays and slices
- Learn slice declaration and initialization
- Master slice operations (append, copy, slicing)
- Explore slice capacity and length
- Understand slice internals and memory management
- Learn best practices for slice usage

## 📁 Project Structure

```
slices/
├── README.md              # This documentation file
└── slices.go             # Slice examples and demonstrations
```

## 🚀 Getting Started

### Prerequisites

Before starting this project, you should be familiar with:
- Basic Go syntax (see [Hello World Project](../hello-world/README.md))
- Variables and data types (see [Variables Project](../variables/README.md))
- Arrays and basic data structures
- Control flow (loops and conditionals)

### Running the Project

```bash
cd slices
go run slices.go
```

**Expected Output:**
```
[1 2 3]
[1 2 3]
Size of slice 3 and capacity 3
Size of slice 4 and capacity 6
Size of slice 5 and capacity 6
Size of slice 6 and capacity 6
Size of slice 7 and capacity 12
Size of slice 8 and capacity 12
Size of slice 9 and capacity 12
Size of slice 10 and capacity 12
Size of slice 11 and capacity 12
Size of slice 12 and capacity 12
Size of slice 13 and capacity 24
```

## 📚 Key Concepts

### What are Slices?

Slices are dynamic views into arrays. They provide:
- **Dynamic sizing**: Can grow and shrink as needed
- **Efficient operations**: Fast append, copy, and slice operations
- **Array backing**: Underlying array provides storage
- **Three components**: Pointer, length, and capacity

### Slice Structure

```go
type slice struct {
    ptr *T      // Pointer to underlying array
    len int     // Length (number of elements)
    cap int     // Capacity (maximum length without reallocation)
}
```

### Slice Declaration

```go
// From array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // [2, 3, 4]

// Direct creation
slice := []int{1, 2, 3, 4, 5}

// Using make
slice := make([]int, 5)     // Length 5, capacity 5
slice := make([]int, 5, 10) // Length 5, capacity 10
```

## 🔍 Code Walkthrough

### Main Function Analysis

```go
package main

import (
    "fmt"
)

func analyze(slice []int) {
    fmt.Printf("Size of slice %d and capacity %d\n", len(slice), cap(slice))
}

func main() {
    arr := [3]int{1, 2, 3}
    slice := []int{1, 2, 3}

    fmt.Println(arr)
    fmt.Println(slice)
    analyze(slice)
    slice = append(slice, 4)
    analyze(slice)
    for i := 5; i <= 10; i++ {
        slice = append(slice, i)
        analyze(slice)
    }
}
```

**Step-by-step analysis:**

1. **Array and Slice Creation (Lines 10-11)**:
   - `arr := [3]int{1, 2, 3}` creates a fixed-size array
   - `slice := []int{1, 2, 3}` creates a slice with same values

2. **Initial Analysis (Lines 13-15)**:
   - Prints both array and slice (they look the same)
   - `analyze(slice)` shows length=3, capacity=3

3. **First Append (Lines 16-18)**:
   - `append(slice, 4)` adds element
   - Capacity doubles to 6 (Go's growth strategy)

4. **Growth Pattern (Lines 19-22)**:
   - Each append triggers capacity analysis
   - Capacity doubles when needed: 3→6→12→24

### Function Analysis

**`analyze` (Lines 6-8)**:
```go
func analyze(slice []int) {
    fmt.Printf("Size of slice %d and capacity %d\n", len(slice), cap(slice))
}
```
- Takes a slice parameter
- Prints length and capacity
- Shows how slice grows over time

## 🎯 Slice Operations and Patterns

### 1. Basic Slice Operations

```go
// Creating slices
slice1 := []int{1, 2, 3, 4, 5}
slice2 := make([]int, 3, 5)

// Slicing
firstThree := slice1[:3]    // [1, 2, 3]
lastThree := slice1[2:]     // [3, 4, 5]
middle := slice1[1:4]       // [2, 3, 4]

// Appending
slice1 = append(slice1, 6)  // [1, 2, 3, 4, 5, 6]
```

### 2. Slice Capacity Management

```go
// Pre-allocate for known size
slice := make([]int, 0, 1000) // Avoids multiple reallocations

// Copy slices
original := []int{1, 2, 3}
copy := make([]int, len(original))
copy(copy, original)
```

### 3. Slice as Function Parameters

```go
func processSlice(slice []int) {
    // Modifies the original slice (slices are reference types)
    slice[0] = 100
}

func returnNewSlice(slice []int) []int {
    // Returns a new slice
    return append(slice, 999)
}
```

### 4. Multi-dimensional Slices

```go
// 2D slice
matrix := make([][]int, 3)
for i := range matrix {
    matrix[i] = make([]int, 3)
}
```

## 🧪 Practice Exercises

### Exercise 1: Slice Filtering

Create a function that filters a slice based on a condition:

```go
func filterInts(slice []int, predicate func(int) bool) []int {
    result := make([]int, 0)
    for _, value := range slice {
        if predicate(value) {
            result = append(result, value)
        }
    }
    return result
}

// Usage
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evens := filterInts(numbers, func(n int) bool {
    return n%2 == 0
})
```

### Exercise 2: Slice Rotation

Implement a function that rotates a slice:

```go
func rotateSlice(slice []int, positions int) []int {
    length := len(slice)
    if length == 0 {
        return slice
    }
    
    positions = positions % length
    if positions < 0 {
        positions += length
    }
    
    result := make([]int, length)
    for i := 0; i < length; i++ {
        newIndex := (i + positions) % length
        result[newIndex] = slice[i]
    }
    
    return result
}
```

### Exercise 3: Slice Deduplication

Create a function that removes duplicates from a slice:

```go
func deduplicate(slice []int) []int {
    seen := make(map[int]bool)
    result := make([]int, 0)
    
    for _, value := range slice {
        if !seen[value] {
            seen[value] = true
            result = append(result, value)
        }
    }
    
    return result
}
```

## 🔗 Related Projects

### Prerequisites
1. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax
2. **[Variables Project](../variables/README.md)** - Variable handling
3. **[Maps Project](../maps/README.md)** - Map data structures

### Next Steps
1. **[Structs Project](../structs/README.md)** - Complex data structures
2. **[Interfaces Project](../interfaces/README.md)** - Interface design with slices

## 📖 Additional Resources

### Official Documentation
- [Go Slices](https://golang.org/ref/spec#Slice_types)
- [Effective Go - Slices](https://golang.org/doc/effective_go.html#slices)
- [Go Slices: usage and internals](https://blog.golang.org/slices)

### Slice Characteristics
- **Reference Type**: Slices are passed by reference
- **Dynamic**: Can grow and shrink as needed
- **Efficient**: Fast append and slice operations
- **Array Backing**: Underlying array provides storage

## 🎯 Key Takeaways

- **Slices are dynamic views** into arrays
- **Three components**: pointer, length, and capacity
- **Append can trigger reallocation** when capacity is exceeded
- **Slices are reference types** - passed by reference
- **Use `make()`** to pre-allocate with specific capacity
- **Length vs Capacity**: Length is current size, capacity is maximum without reallocation
- **Growth strategy**: Capacity typically doubles when needed

## 🚨 Common Mistakes

### ❌ Modifying Shared Underlying Array
```go
// Bad - unexpected side effects
original := []int{1, 2, 3, 4, 5}
slice1 := original[1:4] // [2, 3, 4]
slice2 := original[2:5] // [3, 4, 5]

slice1[0] = 999 // Modifies both slice1 and slice2!

// Good - create independent copies
slice1 := make([]int, len(original[1:4]))
copy(slice1, original[1:4])
```

### ❌ Ignoring Append Return Value
```go
// Bad - might not modify original slice
func appendValue(slice []int, value int) {
    append(slice, value) // Return value ignored!
}

// Good - return the new slice
func appendValue(slice []int, value int) []int {
    return append(slice, value)
}
```

### ❌ Inefficient Slice Growth
```go
// Bad - many reallocations
var slice []int
for i := 0; i < 1000; i++ {
    slice = append(slice, i) // Many reallocations
}

// Good - pre-allocate
slice := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    slice = append(slice, i) // No reallocations
}
```

### ❌ Confusing Length and Capacity
```go
// Bad - confusion about slice size
slice := make([]int, 5, 10)
fmt.Println(len(slice)) // 5 (current size)
fmt.Println(cap(slice)) // 10 (maximum without reallocation)

// Remember: length is current size, capacity is maximum
```

## 🔍 Advanced Concepts

### Slice Internals

```go
// Understanding slice internals
func sliceInternals() {
    arr := [5]int{1, 2, 3, 4, 5}
    slice := arr[1:4] // Points to arr[1], length 3, capacity 4
    
    fmt.Printf("Array: %v\n", arr)
    fmt.Printf("Slice: %v (len=%d, cap=%d)\n", slice, len(slice), cap(slice))
}
```

### Slice Expressions

```go
// Full slice expression: [low:high:max]
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:3:4] // [2, 3] with capacity 3
```

### Nil Slices

```go
// Nil slice behavior
var slice []int // nil slice
fmt.Println(slice == nil) // true
fmt.Println(len(slice))   // 0
fmt.Println(cap(slice))   // 0

// Safe to append to nil slice
slice = append(slice, 1) // Creates new slice
```

---

*Ready to explore more complex data structures? Check out the [Structs project](../structs/README.md)!* 