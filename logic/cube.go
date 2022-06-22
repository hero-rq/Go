package main

import (
	"fmt"
)

type alien struct {
	name string
	age  int
	job  string
}

func (p alien) sayHello() {
	fmt.Printf("Hi I am %s, I am %d years old and my job is %s\n", p.name, p.age, p.job)
}

func (p alien) Pacific() {
	fmt.Printf("Hi, your age is %d here\n", (p.age - 100))
}

func main() {
	panda := alien{"panda", 256, "helicopter"}
	panda.sayHello()
	panda.Pacific()
}
