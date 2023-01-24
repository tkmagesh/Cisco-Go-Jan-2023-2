package main

import (
	"fmt"

	"github.com/tkmagesh/Cisco-Go-Jan-2023-2/09-modularity/calculator"

	// _ "github.com/tkmagesh/Cisco-Go-Jan-2023-2/09-modularity/calculator"

	"github.com/fatih/color"
	app_utils "github.com/tkmagesh/Cisco-Go-Jan-2023-2/09-modularity/calculator/utils"
)

func main() {
	// fmt.Println("app invoked")

	color.Red("app invoked")

	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("Op count :", calculator.OpCount())

	// fmt.Println("is 21 even ?:", utils.IsEven(21))
	fmt.Println("is 21 even ?:", app_utils.IsEven(21))
}
