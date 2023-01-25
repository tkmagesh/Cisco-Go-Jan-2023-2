package main

import (
	"fmt"
)

func main() {
	go f1()
	f2()
	fmt.Scanln() // DO NOT DO THIS
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
