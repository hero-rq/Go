package main

import (
	"fmt"
	)

func main() {
	a := 2
	b := &a
	*b = 13
	fmt.Println(a)
}
