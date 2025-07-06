# Player-Recorder - Advanced Interface Composition 🎵

A Go program that demonstrates advanced interface composition, type assertions, and complex interface hierarchies in Go programming language.

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

This project demonstrates advanced interface concepts in Go, including interface composition, type assertions, and complex interface hierarchies. It shows how to build sophisticated abstractions by combining multiple interfaces and how to work with different types through complex interface relationships.

This is an advanced project that builds upon basic interface concepts to show real-world interface design patterns.

## 📁 Files

- `recorder.go` - The main Go source file demonstrating advanced interface composition

## 💻 Code Structure

```go
package main

import "fmt"

type Player interface {
	Play()
	Stop()
}

type Recorder interface {
	Record()
	StopRecording()
}

type PlayerRecorder interface {
	Player
	Recorder
}

type Microphone struct {
	tone string
}

func (m Microphone) Record() {
	fmt.Println("Recording from microphone", m.tone)
}

func (m Microphone) StopRecording() {
	fmt.Println("Stopping recording from microphone.")
}

type Radio struct {
	station string
}

func (r Radio) Play() {
	fmt.Println("Playing from radio: ", r.station)
}

func (r Radio) Stop() {
	fmt.Println("Stoppint radio: ", r.station)
}

type Smartphone struct {
	model string
}

func (s Smartphone) Play() {
	fmt.Println("Playing music on", s.model)
}

func (s Smartphone) Stop() {
	fmt.Println("Music stopped")
}

func (s Smartphone) Record() {
	fmt.Println("Recording video on", s.model)
}

func (s Smartphone) StopRecording() {
	fmt.Println("Recording saved")
}

func PlayMusic(p Player) {
	p.Play()
}

func RecordAudio(r Recorder) {
	r.Record()
}

func Execute(pr PlayerRecorder) {
	pr.Play()
	pr.Record()
	pr.Stop()
	pr.StopRecording()
}

func SmartPlay(p Player) {
	p.Play()

	if recorder, ok := p.(Recorder); ok {
		fmt.Println("I'm can record")
		recorder.Record()
	}
}

func main() {
	smartphone := Smartphone{"Iphone"}
	radio := Radio{"98 FM"}
	microphone := Microphone{"C"}
	Execute(smartphone)
	PlayMusic(radio)
	PlayMusic(smartphone)
	RecordAudio(microphone)
	RecordAudio(smartphone)
	SmartPlay(smartphone)
}
```

### Code Explanation

- **`Player` and `Recorder` interfaces**: Basic interfaces for playing and recording
- **`PlayerRecorder` interface**: Composed interface that embeds both `Player` and `Recorder`
- **`Microphone`, `Radio`, `Smartphone`**: Different types implementing various interfaces
- **`Execute()`**: Function that works with types implementing both interfaces
- **`SmartPlay()`**: Demonstrates type assertions to check interface capabilities

## 🔑 Key Concepts

### 1. Interface Composition
```go
type PlayerRecorder interface {
	Player
	Recorder
}
```
- Combines multiple interfaces into one
- Any type implementing `PlayerRecorder` must implement both `Player` and `Recorder`
- Enables complex interface hierarchies

### 2. Interface Implementation
```go
// Smartphone implements both Player and Recorder
func (s Smartphone) Play() {
	fmt.Println("Playing music on", s.model)
}

func (s Smartphone) Record() {
	fmt.Println("Recording video on", s.model)
}
```
- `Smartphone` implements both `Player` and `Recorder` interfaces
- Can be used wherever `PlayerRecorder` is expected
- Demonstrates multiple interface satisfaction

### 3. Type Assertions
```go
func SmartPlay(p Player) {
	p.Play()

	if recorder, ok := p.(Recorder); ok {
		fmt.Println("I'm can record")
		recorder.Record()
	}
}
```
- Checks if a `Player` also implements `Recorder`
- Safe type assertion with `ok` pattern
- Conditional behavior based on interface capabilities

### 4. Interface Segregation
```go
// Microphone only implements Recorder
func (m Microphone) Record() {
	fmt.Println("Recording from microphone", m.tone)
}

// Radio only implements Player
func (r Radio) Play() {
	fmt.Println("Playing from radio: ", r.station)
}
```
- Different types implement different interfaces
- Follows interface segregation principle
- Allows for focused, single-purpose types

### 5. Polymorphic Functions
```go
func PlayMusic(p Player) {
	p.Play()
}

func RecordAudio(r Recorder) {
	r.Record()
}

func Execute(pr PlayerRecorder) {
	pr.Play()
	pr.Record()
	pr.Stop()
	pr.StopRecording()
}
```
- Functions work with any type implementing the required interface
- Same function works with different concrete types
- Demonstrates polymorphic behavior

## 🚀 How to Run

### Method 1: Direct Execution

1. Navigate to the player-recorder directory:
   ```bash
   cd player-recorder
   ```

2. Run the program:
   ```bash
   go run recorder.go
   ```

### Method 2: Build and Run

1. Build the executable:
   ```bash
   go build recorder.go
   ```

2. Run the executable:
   ```bash
   # On Unix-like systems (Linux, macOS)
   ./recorder
   
   # On Windows
   recorder.exe
   ```

## 📤 Expected Output

```
Playing from radio:  98 FM
Playing music on Iphone
Recording from microphone C
Recording video on Iphone
Playing music on Iphone
I'm can record
Recording video on Iphone
```

**Explanation**:
- First line: Radio playing (only Player interface)
- Second line: Smartphone playing music
- Third line: Microphone recording (only Recorder interface)
- Fourth line: Smartphone recording video
- Fifth line: Smartphone playing music again
- Sixth and seventh lines: SmartPlay function demonstrating type assertion

## 🎓 Learning Objectives

By completing this project, you will understand:

- ✅ **Interface composition** - How to combine multiple interfaces
- ✅ **Type assertions** - How to safely check interface capabilities
- ✅ **Interface segregation** - How to design focused interfaces
- ✅ **Polymorphic behavior** - How interfaces enable flexible code
- ✅ **Advanced interface patterns** - Real-world interface design
- ✅ **Interface hierarchies** - Building complex interface relationships

## 🔧 Prerequisites

Before starting this project, ensure you have:

- **Go installed** on your system (version 1.16 or later recommended)
- **Understanding of basic interfaces** from the interfaces project
- **Knowledge of structs and methods** from previous projects
- **Text editor or IDE** (VS Code, GoLand, Vim, etc.)

### Recommended Prerequisites

Complete these projects first:
1. **[Hello World Project](../hello-world/README.md)** - Basic Go syntax
2. **[Variables Project](../variables/README.md)** - Variable handling
3. **[Structs Project](../structs/README.md)** - Structs and methods
4. **[Composition Project](../composition/README.md)** - Struct composition
5. **[Interfaces Project](../interfaces/README.md)** - Basic interfaces (essential prerequisite)

## 🔄 Next Steps

After understanding advanced interfaces, explore:

### Advanced Go Concepts
- **Goroutines and channels** - Concurrency in Go
- **Web development** - Building HTTP servers and APIs
- **Database integration** - Working with databases
- **Testing patterns** - Advanced testing techniques

### Related Projects in This Repository
1. **[Interfaces Project](../interfaces/README.md)** - Basic interfaces (prerequisite)
2. **[Structs Project](../structs/README.md)** - Structs and methods (prerequisite)
3. **[Composition Project](../composition/README.md)** - Struct composition (prerequisite)

### Additional Learning Resources
- [Go Tour - Interface Composition](https://tour.golang.org/methods/14)
- [Go by Example - Interfaces](https://gobyexample.com/interfaces)
- [Effective Go - Interfaces](https://golang.org/doc/effective_go.html#interfaces)

## 🐛 Troubleshooting

### Common Issues

**Issue**: `impossible type assertion`
- **Solution**: The type doesn't implement the interface you're asserting

**Issue**: `cannot use type as interface type`
- **Solution**: Make sure the type implements all required methods

**Issue**: Interface composition not working
- **Solution**: Ensure all embedded interfaces are properly implemented

## 📚 Additional Notes

### Interface Design Patterns

**Composition over inheritance**:
```go
// Good - interface composition
type PlayerRecorder interface {
    Player
    Recorder
}

// Bad - large monolithic interface
type BigDevice interface {
    Play()
    Stop()
    Record()
    StopRecording()
    // ... many more methods
}
```

### Type Assertion Best Practices

```go
// Safe type assertion
if recorder, ok := device.(Recorder); ok {
    recorder.Record()
} else {
    fmt.Println("Device cannot record")
}

// Type switch for multiple interfaces
switch v := device.(type) {
case PlayerRecorder:
    v.Play()
    v.Record()
case Player:
    v.Play()
case Recorder:
    v.Record()
default:
    fmt.Println("Unknown device type")
}
```

### Interface Segregation Principle

```go
// Good - focused interfaces
type Player interface {
    Play()
    Stop()
}

type Recorder interface {
    Record()
    StopRecording()
}

// Bad - large interface
type Device interface {
    Play()
    Stop()
    Record()
    StopRecording()
    // Forces all types to implement everything
}
```

---

**🎉 Fantastic! You now understand advanced interface composition in Go!**

*You've mastered the fundamentals of interface design and polymorphism!* 