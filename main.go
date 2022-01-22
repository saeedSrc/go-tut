package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type person struct {
	Firstname string
	Lastname  string
}

var wg sync.WaitGroup

// var mu sync.Mutex

func main() {

	var counter int64
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {

			// mu.Lock()

			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))

			// mu.Unlock()
			wg.Done()
		}()

	}

	wg.Wait()
}
