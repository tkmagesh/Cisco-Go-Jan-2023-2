package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func (emp Employee) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %.2f", emp.Id, emp.Name, emp.Salary)
}

func (emp *Employee) AwardBonus(bonus float32) {
	// (*emp).Salary += bonus
	emp.Salary += bonus //accessing fields using pointer is allowed
}

func main() {

	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}

	// fmt.Println(Format(emp))
	fmt.Println(emp.Format())

	fmt.Println("After awarding 2000 bonus")
	// AwardBonus(&emp, 2000)
	// (&emp).AwardBonus(2000)
	emp.AwardBonus(2000)
	// fmt.Println(Format(emp))
	fmt.Println(emp.Format())
}
