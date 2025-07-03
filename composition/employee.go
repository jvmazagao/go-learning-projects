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
