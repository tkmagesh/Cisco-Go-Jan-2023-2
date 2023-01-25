package main

import (
	"fmt"
	"time"
)

func main() {
	go f1()
	f2()
	time.Sleep(time.Second) // DO NOT DO THIS
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
