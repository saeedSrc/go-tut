package main

import (
	"fmt"
)

func main() {

	c1 := make(chan string)
	// c2 := make(chan string)
	// go func() {

	// }()

	c1 <- "one"
	go func() {

		fmt.Println(<-c1)
	}()

	fmt.Println("hi")

	// for i := 0; i < 2; i++ {
	//     select {
	//     case msg1 := <-c1:
	//         fmt.Println("received", msg1)
	//     case msg2 := <-c2:
	//         fmt.Println("received", msg2)
	//     }
	// }
	fmt.Println("about to exit")
}
