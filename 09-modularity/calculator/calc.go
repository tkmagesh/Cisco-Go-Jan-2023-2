package calculator

import "fmt"

var op_count int

func init() {
	fmt.Println("[calc.go] calculator initialized - 1")
}

func init() {
	fmt.Println("[calc.go] calculator initialized - 2")
}

func OpCount() int {
	return op_count
}
