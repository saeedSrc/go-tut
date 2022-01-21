package main

import (
	"fmt"
	// "sync"
)

type person struct {
	Firstname string
	Lastname  string
}

// var wg sync.WaitGroup

func main() {


	counter := 0
	// wg.Add(1)

	for i :=0;i<10;i++ {
		go func(){
			v := counter
			v++
			counter = v
			fmt.Println(counter)
	
		}()
	}
	
	// wg.Wait()
}
