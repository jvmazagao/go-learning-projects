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
