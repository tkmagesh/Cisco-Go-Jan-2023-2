package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 10
	fmt.Println("Len(ch) =", len(ch))
	ch <- 20
	fmt.Println("Len(ch) =", len(ch))
	ch <- 30
	fmt.Println("Len(ch) =", len(ch))

	fmt.Println(<-ch)
	fmt.Println("Len(ch) =", len(ch))
	fmt.Println(<-ch)
	fmt.Println("Len(ch) =", len(ch))
	fmt.Println(<-ch)
	fmt.Println("Len(ch) =", len(ch))
}
