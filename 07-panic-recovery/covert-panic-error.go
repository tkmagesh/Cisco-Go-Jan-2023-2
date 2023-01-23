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
	q, r, err := divideWrapper(100, 0)
	if err != nil {
		fmt.Printf("err : %v, try valid inputs\n", err)
		return
	}
	fmt.Printf("quotient = %d and remainder = %d\n", q, r)
}

func divideWrapper(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error) // typecast to an error
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

//3rd party library
func divide(x, y int) (quotient, remainder int) {
	fmt.Println("[calculating quotient]")
	if y == 0 {
		panic(DivideByZeroError) // creating application specific panic
	}
	quotient = x / y // panics if y == 0
	fmt.Println("[calculating remainder]")
	remainder = x % y
	return
}
