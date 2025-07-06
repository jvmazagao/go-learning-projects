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
