package main

import (
	"fmt"
	"strings"
	)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, upperName := lenAndUpper("panda")
	fmt.Println(totalLength, upperName) 
}
