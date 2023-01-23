/* complex numbers */
package main

import "fmt"

func main() {
	// var c1 complex128 = 4 + 5i
	c1 := 4 + 5i
	fmt.Println(c1)

	c2 := 11 + 16i
	c3 := c1 + c2
	fmt.Println("c3 = ", c3)

	fmt.Println("real(c3) =", real(c3))
	fmt.Println("imag(c3) =", imag(c3))
}
