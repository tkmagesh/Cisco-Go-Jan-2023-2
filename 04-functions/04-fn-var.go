/* Higher order functions - functions as values to variables */
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

	var getGreetMsg func(string) string
	getGreetMsg = func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a good day!", userName)
	}
	msg := getGreetMsg("Suresh")
	fmt.Println(msg)

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	result := add(100, 200)
	fmt.Println(result)

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}
	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

}
