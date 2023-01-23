/* Higher order functions - functions as arguments (applied) */
package main

import (
	"fmt"
	"log"
	"time"
)

type MathOperation func(int, int)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
	*/

	//logging
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	//profiling
	/*
		profileAdd := getProfileOperation(add)
		profileAdd(100, 200)

		profileSubtract := getProfileOperation(subtract)
		profileSubtract(100, 200)
	*/

	//profiling & logging
	logAdd := getLogOperation(add)
	profileLogAdd := getProfileOperation(logAdd)
	profileLogAdd(100, 200)

	logSubtract := getLogOperation(subtract)
	profileLogSubtract := getProfileOperation(logSubtract)
	profileLogSubtract(100, 200)

}

func getProfileOperation(oper MathOperation) MathOperation {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		elapsed := time.Since(start)
		log.Println("Time Taken :", elapsed)
	}
}

func getLogOperation(oper MathOperation) MathOperation {
	return func(x, y int) {
		log.Println("Operation started")
		oper(100, 200)
		log.Println("Operation completed")
	}
}

/*
func getProfileOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		elapsed := time.Since(start)
		log.Println("Time Taken :", elapsed)
	}
}

func getLogOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started")
		oper(100, 200)
		log.Println("Operation completed")
	}
}
*/

//3rd party code
func add(x, y int) {
	fmt.Println("Add result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result :", x-y)
}
