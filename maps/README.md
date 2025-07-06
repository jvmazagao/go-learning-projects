# Maps in Go 🗺️

**Directory**: `/maps` | **Difficulty**: ⭐⭐ Intermediate

This project demonstrates Go's map data structure, which provides an efficient way to store key-value pairs. Maps are one of Go's most useful built-in data structures for associative data.

## 🎯 Learning Objectives

- Understand Go's map data structure
- Learn map declaration and initialization
- Master map operations (insert, update, delete, lookup)
- Explore map iteration with `range`
- Understand map behavior and best practices
- Apply maps to practical problems (word counting)

## 📁 Project Structure

```
maps/
├── README.md              # This documentation file
├── maps.go               # Basic map operations
└── count-words/          # Practical map application
    └── main.go          # Character counting example
```

## 🚀 Getting Started

### Prerequisites

Before starting this project, you should be familiar with:
- Basic Go syntax (see [Hello World Project](../hello-world/README.md))
- Variables and data types (see [Variables Project](../variables/README.md))
- Control flow (if statements)
- Loops and iteration

### Running the Projects

**Basic Maps:**
```bash
cd maps
go run maps.go
```

**Word Counter:**
```bash
cd maps/count-words
go run main.go
```

## 📚 Key Concepts

### What are Maps?

Maps are Go's built-in associative data type (hash table/dictionary). They store key-value pairs where:
- **Keys** must be comparable types (strings, numbers, booleans, etc.)
- **Values** can be any type
- **Lookup** is very fast (O(1) average case)

### Map Declaration

```go
// Declaration with make
ages := make(map[string]int)

// Declaration with literal
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
}

// Zero value is nil
var ages map[string]int // nil map
```

### Map Operations

```go
// Insert/Update
ages["Charlie"] = 35

// Lookup
age := ages["Alice"]

// Check if key exists
age, exists := ages["David"]

// Delete
delete(ages, "Bob")

// Iterate
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}
```

## 🔍 Code Walkthrough

### Basic Maps (`maps.go`)

```go
package main

import "fmt"

func main() {
    ages := map[string]int{
        "Bob":  42,
        "Jhon": 35,
        "Mary": 27,
    }

    fmt.Println(ages["Jhon"])
    fmt.Println(ages["David"])

    age, exists := ages["David"]

    if exists {
        fmt.Println(age)
    } else {
        fmt.Println("Not Exists")
    }

    delete(ages, "Bob")

    for name, age := range ages {
        fmt.Printf("The age of %s is %d\n", name, age)
    }
}
```

**Step-by-step analysis:**

1. **Map Initialization (Lines 5-9)**:
   - Creates a map with string keys and int values
   - Uses literal syntax for initialization
   - Contains three key-value pairs

2. **Direct Lookup (Lines 11-12)**:
   - `ages["Jhon"]` returns `35` (key exists)
   - `ages["David"]` returns `0` (key doesn't exist, zero value)

3. **Safe Lookup (Lines 14-20)**:
   - Uses comma ok idiom: `age, exists := ages["David"]`
   - `exists` is `false` for non-existent key
   - Prints "Not Exists" since David isn't in the map

4. **Deletion (Line 22)**:
   - `delete(ages, "Bob")` removes Bob from the map
   - Safe operation even if key doesn't exist

5. **Iteration (Lines 24-26)**:
   - `range` iterates over all key-value pairs
   - Order is not guaranteed (maps are unordered)

### Word Counter (`count-words/main.go`)

```go
package main

import "fmt"

func main() {
    sentence := "the quick brown fox jumps over the lazy fox"

    counter := make(map[rune]int)

    for _, c := range sentence {
        counter[c]++
    }

    for key, val := range counter {
        fmt.Printf("This it the counter %d of the char %c\n", val, key)
    }
}
```

**Analysis:**

1. **String Processing**: Iterates over each character (`rune`) in the sentence
2. **Character Counting**: Uses each character as a key in the map
3. **Increment**: `counter[c]++` increments the count for each character
4. **Output**: Prints each character and its frequency

## 🎯 Map Patterns and Best Practices

### 1. Safe Map Access

```go
// Always check if key exists
if value, exists := myMap[key]; exists {
    // Use value
} else {
    // Handle missing key
}
```

### 2. Map Initialization

```go
// For known size
m := make(map[string]int, 100)

// For small maps
m := map[string]int{
    "a": 1,
    "b": 2,
}
```

### 3. Map as Set

```go
// Using map as a set
set := make(map[string]bool)
set["item1"] = true
set["item2"] = true

// Check membership
if set["item1"] {
    fmt.Println("item1 is in set")
}
```

### 4. Nested Maps

```go
// Map of maps
users := map[string]map[string]string{
    "alice": {
        "email": "alice@example.com",
        "role":  "admin",
    },
    "bob": {
        "email": "bob@example.com",
        "role":  "user",
    },
}
```

## 🧪 Practice Exercises

### Exercise 1: Word Frequency Counter

Create a function that counts word frequency in a text:

```go
func countWords(text string) map[string]int {
    words := strings.Fields(text)
    frequency := make(map[string]int)
    
    for _, word := range words {
        frequency[strings.ToLower(word)]++
    }
    
    return frequency
}
```

### Exercise 2: Map as Cache

Implement a simple cache using maps:

```go
type Cache struct {
    data map[string]interface{}
    mu   sync.RWMutex
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]interface{}),
    }
}

func (c *Cache) Set(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    value, exists := c.data[key]
    return value, exists
}
```

### Exercise 3: Grouping Data

Group items by a common property:

```go
type Person struct {
    Name string
    Age  int
    City string
}

func groupByCity(people []Person) map[string][]Person {
    groups := make(map[string][]Person)
    
    for _, person := range people {
        groups[person.City] = append(groups[person.City], person)
    }
    
    return groups
}
```

## 🔗 Related Projects

### Prerequisites
1. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax
2. **[Variables Project](../variables/README.md)** - Variable handling
3. **[Slices Project](../slices/README.md)** - Array and slice concepts

### Next Steps
1. **[Structs Project](../structs/README.md)** - Complex data structures
2. **[Interfaces Project](../interfaces/README.md)** - Interface design with maps

## 📖 Additional Resources

### Official Documentation
- [Go Maps](https://golang.org/ref/spec#Map_types)
- [Effective Go - Maps](https://golang.org/doc/effective_go.html#maps)

### Map Characteristics
- **Unordered**: Iteration order is not guaranteed
- **Reference Type**: Maps are passed by reference
- **Nil Maps**: Zero value is `nil`, cannot be written to
- **Concurrency**: Maps are not safe for concurrent access

## 🎯 Key Takeaways

- **Maps are hash tables** providing O(1) average lookup time
- **Keys must be comparable** (strings, numbers, booleans, etc.)
- **Use comma ok idiom** for safe lookups
- **Maps are unordered** - iteration order is not guaranteed
- **Maps are reference types** - passed by reference
- **Use `make()`** to create maps or literal syntax
- **Check for existence** before using values
- **Maps are not thread-safe** - use sync.Map for concurrent access

## 🚨 Common Mistakes

### ❌ Accessing Nil Maps
```go
// Bad - will panic
var m map[string]int
m["key"] = 1 // panic: assignment to entry in nil map

// Good
m := make(map[string]int)
m["key"] = 1
```

### ❌ Ignoring the "ok" Value
```go
// Bad - might get zero value
value := myMap["key"]

// Good - check if key exists
if value, exists := myMap["key"]; exists {
    // Use value
}
```

### ❌ Assuming Map Order
```go
// Bad - order is not guaranteed
for key, value := range myMap {
    // Don't rely on order
}

// Good - sort keys if order matters
keys := make([]string, 0, len(myMap))
for key := range myMap {
    keys = append(keys, key)
}
sort.Strings(keys)
for _, key := range keys {
    value := myMap[key]
    // Process in sorted order
}
```

---

*Ready to explore more data structures? Check out the [Slices project](../slices/README.md)!* 