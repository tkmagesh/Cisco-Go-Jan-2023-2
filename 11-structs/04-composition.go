package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

type PerishableProduct struct {
	Product
	Expiry string
}

// utility constructor function
func NewPerishableProduct(id int, name string, cost float32, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	// grapes := PerishableProduct{Product{100, "Grapes", 40}, "2 Days"}
	/*
		grapes := PerishableProduct{
			Product: Product{
				Id:   100,
				Name: "Grapes",
				Cost: 40,
			},
			Expiry: "2 Days",
		}
	*/
	grapes := NewPerishableProduct(100, "Grapes", 40, "2 Days")

	// fmt.Println(grapes.Product.Id, grapes.Product.Cost)
	fmt.Println(grapes.Id, grapes.Name, grapes.Cost, grapes.Expiry)
}

/*
Write a "Format" function that returns a formatted string of the given "Product"
Write an "ApplyDiscount" function that updates the given "Product's" cost by applying the given discount (%)

Use the above functions for a "PerishableProduct" object
*/
