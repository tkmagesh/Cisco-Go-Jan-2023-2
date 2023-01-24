package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
	Org    *Organization // association
}

type Organization struct {
	Name string
	City string
}

func main() {
	cisco := Organization{
		Name: "Cisco",
		City: "Bangalore",
	}

	e1 := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
		Org:    &cisco,
	}

	e2 := Employee{
		Id:     200,
		Name:   "Suresh",
		Salary: 20000,
		Org:    &cisco,
	}

	cisco.City = "Pune"
	fmt.Println("e1.Org.City =", e1.Org.City)
	fmt.Println("e2.Org.City =", e2.Org.City)
}
