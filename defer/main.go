package main

import "fmt"

func main() {
	fmt.Println("Starting")

	defer fmt.Println("This prints...") // When does this run?

	fmt.Println("Middle")
	fmt.Println("Ending")
}
