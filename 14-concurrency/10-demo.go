package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println("Result :", result)
}

//producer
func add(x, y int, ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- x + y
}
