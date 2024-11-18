package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	people := [5]string{"nico", "flunn", "cho", "norm", "ai"}
	for _, person := range people {
		go isSexy(person, channel)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println("waiting for ", i)
		fmt.Println(<-channel)
	}
}
func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 5)
	//fmt.Println(person)
	channel <- person + " is sexy"
}
