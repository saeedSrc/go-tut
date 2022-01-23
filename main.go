package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	var ctx = context.Background()

	var ctx2, cancel = context.WithCancel(ctx)


	go func() {
		n := 0
		for {
			select {
			case <-ctx2.Done():
				return
			default:
				n++
				time.Sleep(time.Microsecond * 100)
                fmt.Println("working:", n)
			}
		}
	} ()



	time.Sleep(2 * time.Second)

	cancel()



	fmt.Println(runtime.NumGoroutine())

    fmt.Println("exit")
}
