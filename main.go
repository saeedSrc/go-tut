package main

import "fmt"

func main() {

	var chan1 = make(chan int)
	var chan2 = make(chan int)
	var chan3 = make(chan int)

	go send(chan1, chan2, chan3)
	receive(chan1, chan2, chan3)

	fmt.Println("about to exit")
}

func receive(e, o, q <-chan int) {

	for {
		select {
		case v := <-e:
			fmt.Println(v)

		case v := <-o:
			fmt.Println(v)

		case v := <-q:
			fmt.Println(v)
			return
		}
	}

}

func send(e, o, q chan<- int) {

	for i := 0; i < 11; i++ {

		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
	close(q)

}
