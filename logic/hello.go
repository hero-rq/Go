package main

import (
	"fmt"
	"strings"
	)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("hello! ")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, upperName := lenAndUpper("panda")
	fmt.Println(totalLength, upperName) 
}
