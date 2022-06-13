package main

import (
	"fmt"
	"github.com/hero-rq/learngo/mydict"
	)

func main() {
	dictionary := mydict.Dictionary{}
	baseword := "hello"
	dictionary.Add(baseword, "alphaka")
	err := dictionary.Update(baseword, "panda")
	dictionary.Delete(baseword)
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseword)
	fmt.Println(word)
}
