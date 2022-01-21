package main

import (
	"fmt"
	"sync"
)

type person struct {
	Firstname string
	Lastname  string
}

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go foo()
	bar()

	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
