package main

import (
	"fmt"
	"sync"
)

func main() {

	var chan1 = make(chan int)
	var chan2 = make(chan int)
	var chan3 = make(chan int)


	go send(chan1, chan2)


	go receive(chan1, chan2, chan3)

	for v:= range chan3 {
		fmt.Println(v)
	}


	fmt.Println("abou to ecxit")

}

func send(e, o chan<- int){
	for i:=0; i<11; i++ {
		if i%2 == 0 {
          e<- i
		} else {
			o<-i
		}

		
	}
	close(e)
		close(o)
}

func receive(e, o <-chan int, fan chan<- int){
	var wg  sync.WaitGroup

	wg.Add(2)

	go func(){
		for v:= range e {
			fan <- v
		}

		wg.Done()

	}()

	go func(){
		for v:= range o {
			fan <- v
		}

		wg.Done()
		}()

		wg.Wait()

		close(fan)

}

	
