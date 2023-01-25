package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genFib(10)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func genFib(count int) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
		for i := 0; i < count; i++ {
			ch <- x
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		}
	}()
	return ch
}
