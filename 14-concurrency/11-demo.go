package main

import (
	"fmt"
	"time"
)

//consumer
func main() {

	ch := add(100, 200)
	ch <- 100
	result := <-ch
	fmt.Println("Result :", result)
}

//producer
func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- x + y
	}()
	return ch
}
