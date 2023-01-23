package main

import "fmt"

func main() {
	var fn func()
	fn = func() {
		fmt.Println("Hi")
	}
	fn()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	// Try converting the following as above
	/*
		msg := func(userName string) string {
			return fmt.Sprintf("Hi %s, Have a good day!", userName)
		}("Suresh")
		fmt.Println(msg)

		result := func(x, y int) int {
			return x + y
		}(100, 200)
		fmt.Println(result)

		q, r := func(x, y int) (quotient, remainder int) {
			quotient = x / y
			remainder = x % y
			return
		}(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
	*/
}
