package main

import "fmt"

// var greetMsg string = "Hi %s, Have a nice day!\n"

var greetMsg = "Hi %s, Have a nice day!\n"

// greetMsg := "Hi %s, Have a nice day!\n" // INVALID

var i = 100

func main() {
	/*
		var name string
		name = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	/*
		var name string = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	//type inference
	/*
		var name = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	/* := syntax is application ONLY to function scoped variables (NOT package scoped ones) */
	name := "Magesh"
	// greetMsg := "Hi %s, Have a nice day!\n"
	fmt.Printf(greetMsg, name)

	/* multiple variables */
	/*
		var x int
		var y int
		var result int
		var str string
		x = 100
		y = 200
		result = x + y
		str = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		result = x + y
		str = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		result = x + y
		str = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	/* Declaration & initialization */
	/*
		var x int = 100
		var y int = 200
		var result int = x + y
		var str string = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/
	/*
		var x, y int = 100, 200
		var result int = x + y
		var str string = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			result int    = x + y
			str    string = "Sum of %d and %d is %d\n"
		)
		fmt.Printf(str, x, y, result)
	*/

	/* type inference */
	/*
		var (
			x, y   = 100, 200
			result = x + y
			str    = "Sum of %d and %d is %d\n"
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/* using := */

	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)
}
