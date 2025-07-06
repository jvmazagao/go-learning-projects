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
