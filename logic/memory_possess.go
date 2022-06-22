package main

import (
	"fmt"
)

func main() {
	a := 2
	b := &a
	fmt.Println(b)
	a = 23
	a = 2425
	fmt.Println(*b)
}
