package main

import "fmt"

func main() {
	sentence := "the quick brown fox jumps over the lazy fox"

	counter := make(map[rune]int)

	for _, c := range sentence {
		counter[c]++
	}

	for key, val := range counter {
		fmt.Printf("This it the counter %d of the char %c\n", val, key)
	}
}
