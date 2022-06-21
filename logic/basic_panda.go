package main

import (
	"fmt"
)

func add(a ...int) int {
	var total int
	for _, item := range a {
		total += item
	}
	return total
}

func main() {
	pandora := add(1, 24, 2526, 12525, 1225, 12000, 1001010)
	fmt.Println(pandora)
}
