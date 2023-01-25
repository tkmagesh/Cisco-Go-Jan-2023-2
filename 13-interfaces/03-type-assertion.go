package main

import "fmt"

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "This is a string"
	x = true
	x = 99.98
	x = struct{}{}
	fmt.Println(x)

	// x = "This is a string"
	x = 100
	// y := x.(int) + 200
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	x = 100
	// x = "This is a string"
	// x = true
	// x = 99.98
	// x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 =", val+200)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, 99.5% of x is ", val*(99.5/100))
	case struct{}:
		fmt.Println("x is an empty struct")
	default:
		fmt.Println("Unknown type")
	}

}
