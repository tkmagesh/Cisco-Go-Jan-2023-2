package main

import "fmt"

func main() {
	var no int
	no = 100

	// from value to address
	var noPtr *int
	noPtr = &no
	fmt.Println(no, noPtr)

	//dereferencing - from address to value
	fmt.Println(*noPtr)

	fmt.Println(no == *(&no))

	fmt.Println("Before incrementing, no = ", no)
	increment(no)
	fmt.Println("After incrementing, no = ", no)
}

func increment(no int) /* no return result */ {
	no++
}
