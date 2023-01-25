package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := time.After(10 * time.Second)
	ch := genFib(stopCh)

	fmt.Println("Runs for 10 seconds!")

	for no := range ch {
		fmt.Println(no)
	}
}

func genFib(stopCh <-chan time.Time) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-stopCh:
				break LOOP
			default:
				ch <- x
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}
