package main

import (
	"fmt"
)

func main() {
	pandora := 12324252515
	cookie := fmt.Sprintf("%b", pandora)
	fmt.Println(pandora, cookie)
}
