package main

import "fmt"

func main() {
	var chan1 = make(chan int, 10)

	// go func() {
		for i := 0; i < 10; i++ {
			chan1 <- i

			fmt.Println(<-chan1)
		}
		// }()
		
		
		go func ()  {
			for v := range chan1 {
				fmt.Println(v)
			}
			}()
			close(chan1)
}
