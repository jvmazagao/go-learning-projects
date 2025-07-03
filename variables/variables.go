package main

import "fmt"

func main() {
	var age int = 27
	var name string = "John"
	var x, y int = 10, 10
	pi := 3.14
	lastName := "Doe"
	isTrue := true

	fmt.Println(age)
	fmt.Println(name, lastName)
	fmt.Println(x, y)
	fmt.Println(pi)
	fmt.Println(isTrue)
	fmt.Printf("pi is type: %T\n", pi)
}
