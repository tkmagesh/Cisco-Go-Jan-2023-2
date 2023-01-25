package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genFib()
	for no := range ch {
		fmt.Println(no)
	}
}

func genFib() <-chan int {
	count := 15
	ch := make(chan int)
	go func() {
		x, y := 0, 1
		for i := 0; i < count; i++ {
			ch <- x
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		}
		close(ch)
	}()
	return ch
}
