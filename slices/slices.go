package main

import (
	"fmt"
)

func analyze(slice []int) {
	fmt.Printf("Size of slice %d and capacity %d\n", len(slice), cap(slice))
}

func main() {
	arr := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}

	fmt.Println(arr)
	fmt.Println(slice)
	analyze(slice)
	slice = append(slice, 4)
	analyze(slice)
	for i := 5; i <= 10; i++ {
		slice = append(slice, i)
		analyze(slice)
	}
}
