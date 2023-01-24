package calculator

import "fmt"

func init() {
	fmt.Println("[add.go] calculator initialized")
}

func Add(x, y int) int {
	op_count++
	return x + y
}
