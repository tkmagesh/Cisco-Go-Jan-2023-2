package calculator

import "fmt"

func init() {
	fmt.Println("[subtract.go] calculator initialized")
}

func Subtract(x, y int) int {
	op_count++
	return x - y
}
