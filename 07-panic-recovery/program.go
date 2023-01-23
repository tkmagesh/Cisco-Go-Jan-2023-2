package main

import (
	"errors"
	"fmt"
	"log"
)

var DivideByZeroError error = errors.New("divide by 0 error")

func main() {
	defer func() {
		if err := recover(); err != nil {
			//the deferred function is executed because of a panic
			log.Println(err)
			return
		}
		fmt.Println("Thank you!")
	}()
	q, r := divide(100, 7)
	fmt.Printf("quotient = %d and remainder = %d\n", q, r)
}

func divide(x, y int) (quotient, remainder int) {
	fmt.Println("[calculating quotient]")
	quotient = x / y // panics if y == 0
	fmt.Println("[calculating remainder]")
	remainder = x % y
	return
}
