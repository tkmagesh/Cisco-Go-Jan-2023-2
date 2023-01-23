/* Higher order functions - functions as return values */
package main

import "fmt"

func main() {
	fn := getFn()
	fn()
}

func getFn() func() {
	return f2
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
