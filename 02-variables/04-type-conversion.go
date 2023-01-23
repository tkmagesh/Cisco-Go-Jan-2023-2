/* type conversion */

package main

import "fmt"

func main() {
	/*
		var f1 float32 = 10.998987
		var i1 int32 = f1
	*/

	var i1 int32 = 100
	// var f1 float32 = i1
	/* type conversion is done by using the type name like a function */
	var f1 float32 = float32(i1)
	fmt.Println(f1)
}
