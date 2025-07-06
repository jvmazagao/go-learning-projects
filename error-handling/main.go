package main

import "fmt"

func calculateDivision(number, divider float32) (float32, error) {
	if divider == 0 {
		return 0, fmt.Errorf("Divided by 0")
	}

	return number / divider, nil
}

func main() {
	result, err := calculateDivision(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)

	// Try with zero
	result2, err2 := calculateDivision(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}

	fmt.Println("Result:", result2)
}
