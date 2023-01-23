/* Higher order functions - functions as arguments (applied) */
package main

import (
	"fmt"
	"log"
)

func main() {
	// step - 1
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	// step - 2
	/*
		log.Println("Operation started")
		add(100, 200)
		log.Println("Operation completed")

		log.Println("Operation started")
		subtract(100, 200)
		log.Println("Operation completed")
	*/

	// step - 3
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	// step - 4
	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)

}

func logOperation(oper func(int, int), x, y int) {
	log.Println("Operation started")
	oper(100, 200)
	log.Println("Operation completed")
}

func logAdd(x, y int) {
	log.Println("Operation started")
	add(100, 200)
	log.Println("Operation completed")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(100, 200)
	log.Println("Operation completed")
}

//3rd party code
func add(x, y int) {
	fmt.Println("Add result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result :", x-y)
}
