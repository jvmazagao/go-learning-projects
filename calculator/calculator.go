package main

import (
	"fmt"
	"os"
	"strconv"
)

func Calculator(a int, b int, operation string) int {
	switch operation {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("division by zero")
		}
		return a / b
	default:
		panic("operation not permited")
	}
}

func ConvertNumber(num string) int {
	res, err := strconv.Atoi(num)
	if err != nil {
		panic("failed to convert number")
	}

	return res
}

func GlobalErrorHandler() {
	if r := recover(); r != nil {
		fmt.Fprintln(os.Stderr, "💥 Error:", r)
		os.Exit(1)
	}
}

func main() {

	defer GlobalErrorHandler()

	if len(os.Args) < 4 {
		panic("Please provide a number")
	}

	num1 := ConvertNumber(os.Args[1])
	num2 := ConvertNumber(os.Args[3])
	operator := os.Args[2]

	result := Calculator(num1, num2, operator)

	fmt.Printf("The result from (%d %s %d) is %d \n", num1, operator, num2, result)
}
