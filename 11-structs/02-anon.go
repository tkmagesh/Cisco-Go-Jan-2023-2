package main

import "fmt"

func main() {
	/*
		type Employee struct {
			Id     int
			Name   string
			Salary float32
		}
	*/
	emp := struct {
		Id     int
		Name   string
		Salary float32
	}{100, "Magesh", 10000}
	// fmt.Println(emp)
	fmt.Printf("%#v\n", emp)

	empty := struct{}{}
	fmt.Println(empty)

}
