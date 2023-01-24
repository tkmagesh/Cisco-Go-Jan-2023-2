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

	fmt.Println("Before incrementing, no = ", no) // 100
	increment(&no)
	fmt.Println("After incrementing, no = ", no) // 101

	x, y := 100, 200
	fmt.Printf("Before swapping, x = %d and y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swapping, x = %d and y = %d\n", x, y)
}

func increment(n *int) /* no return result */ {
	fmt.Println("&n =", n)
	*n++
}

func swap(n1, n2 *int) /* no return result */ {
	*n1, *n2 = *n2, *n1
}
