package main

import (
	"fmt"
)

func main() {
	foods := []string{"panda", "cookie", "coke"}
	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}
	fmt.Println(len(foods))
}
