package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go populate(ch1)

	go fanoutin(ch1, ch2)

	for v:=range ch2 {
		fmt.Println(v)
	}

    fmt.Println("exit")
}


func populate(c chan int) { 
	for i:=0;i<10;i++ {
		     c<-i
	}
	close(c)
}


func fanoutin(c1, c2 chan int) {


	var wg  sync.WaitGroup

	for v:= range c1 {
		wg.Add(1)
		go func (v2 int)  {
			c2<-timeConsumingWork(v2)
			wg.Done()
		}(v)
   
	}
 
	wg.Wait()
	close(c2)
}

func timeConsumingWork(v int) int {

	time.Sleep(time.Microsecond * 10)

	return v + rand.Intn(500)
}