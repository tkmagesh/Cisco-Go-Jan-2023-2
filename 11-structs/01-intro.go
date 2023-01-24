package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {
	// var emp Employee
	// var emp Employee = Employee{100, "Magesh", 10000} // NOT advisable

	// var emp Employee = Employee{Id: 100, Name: "Magesh", Salary: 10000}
	/*
		var emp Employee = Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
	*/

	/*
		var emp = Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
	*/
	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	// fmt.Println(emp)
	// fmt.Printf("%#v\n", emp)
	fmt.Printf("%+v\n", emp)

	fmt.Println("Accessing struct fields")
	// fmt.Printf("Id = %d, Name = %q, Salary = %.2f\n", emp.Id, emp.Name, emp.Salary)
	fmt.Println(Format(emp))

	fmt.Println("Structs are VALUES")
	e1 := Employee{200, "Suresh", 20000}
	e2 := Employee{200, "Suresh", 20000}
	fmt.Printf("&e1 = %p and &e2 = %p\n", &e1, &e2)
	fmt.Println("e1 ==  e2 ? ", e1 == e2)

	fmt.Println("After awarding 2000 bonus")
	AwardBonus(&emp, 2000)
	fmt.Printf("emp = %+v\n", emp)
}

func Format(emp Employee) string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %.2f", emp.Id, emp.Name, emp.Salary)
}

func AwardBonus(emp *Employee, bonus float32) {
	// (*emp).Salary += bonus
	emp.Salary += bonus //accessing fields using pointer is allowed
}
