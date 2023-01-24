package main

import (
	"fmt"

	_ "github.com/tkmagesh/Cisco-Go-Jan-2023-2/09-modularity/calculator"
	app_utils "github.com/tkmagesh/Cisco-Go-Jan-2023-2/09-modularity/calculator/utils"
)

func main() {
	fmt.Println("app invoked")

	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println("Op count :", calculator.OpCount())
	*/

	// fmt.Println("is 21 even ?:", utils.IsEven(21))
	fmt.Println("is 21 even ?:", app_utils.IsEven(21))
}
