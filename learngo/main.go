package main

import (
	"fmt"

	"github.com/chancebro/learngo/mydict"
)

func main() {
	// Dictionary 생성
	dictionary := mydict.Dictionary{"first": "first word"}
	word := "first"
	dictionary.Add(word, "first word")
	dictionary.Add("hi", "first word")
	error := dictionary.Update("hi", "second word")
	if error != nil {
		fmt.Println(error)
	}
	dictionary.Delete("hi")

	fmt.Println(dictionary["hi"])
}
