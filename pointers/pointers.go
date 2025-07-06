package main

import "fmt"

func changeValue(val int) {
	val = 100
}

func changePointer(val *int) {
	*val = 100
}

func main() {
	x := 42
	var y *int = &x

	fmt.Println(x)
	fmt.Println("Pointer: ", y)
	fmt.Println("Address: ", &x)
	fmt.Println("Address: ", &y)

	*y = 30

	fmt.Println(x)
	changeValue(x)
	fmt.Println(x)
	fmt.Println(y)
	changePointer(y)
	fmt.Println(x)
	fmt.Println(y)
}
