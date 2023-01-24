package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f", product.Id, product.Name, product.Cost)
}

func (product *Product) ApplyDiscount(discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}

//PerishableProduct (inheriting Product)
type PerishableProduct struct {
	Product
	Expiry string
}

// overriding Product's Format method
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
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

	grapes := NewPerishableProduct(100, "Grapes", 40, "2 Days")

	// fmt.Println(Format(grapes.Product))
	// fmt.Println(grapes.Product.Format())
	fmt.Println(grapes.Format())

	fmt.Println("After applying 10% discount")
	// ApplyDiscount(&grapes.Product, 10)
	// grapes.Product.ApplyDiscount(10)
	grapes.ApplyDiscount(10)

	// fmt.Println(Format(grapes.Product))
	// fmt.Println(grapes.Product.Format())
	fmt.Println(grapes.Format())
}
