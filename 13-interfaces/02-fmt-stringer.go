package main

import (
	"fmt"
	"strings"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

//fmt.Stringer implementation
func (p Product) String() string {
	return p.Format()
}

/*
Write the apis for the following

"IndexOf" => return the index of the given product (otherwise return -1)
	ex:  returns the index of the given product

"Includes" => return true if the given product is present in the collection else return false
	ex:  returns true if the given product is present in the collection

"Filter" => return a new collection of products that satisfy the given condition
	some of the use cases:
		1. filter all costly products (cost > 1000)
			OR
		2. filter all stationary products (category = "Stationary")
			OR
		3. filter all costly stationary products
		etc
"Any" => return true if any of the product in the collections satifies the given criteria
		use cases:
			1. are there any costly product (cost > 1000)?
			OR
			2. are there any stationary product (category = "Stationary")?
			ORF
			3. are there any costly stationary product?
			etc
"All" => return true if all the products in the collections satifies the given criteria
		use cases:
			1. are all the products costly products (cost > 1000)?
			OR
			2. are all the products stationary products (category = "Stationary")?
			OR
			3. are all the products costly stationary products?
			etc
*/

type Products []Product

func (products Products) IndexOf(product Product) int {
	for idx, p := range products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products Products) Filter(predicate func(Product) bool) Products {
	var result Products
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

/*
func (products Products) Format() string {
	var sb strings.Builder
	for _, p := range products {
		sb.WriteString(fmt.Sprintf("%s\n", p.Format()))
	}
	return sb.String()
}
*/

// fmt.Stringer implementation
func (products Products) String() string {
	var sb strings.Builder
	for _, p := range products {
		sb.WriteString(fmt.Sprintf("%s\n", p))
	}
	return sb.String()
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	// fmt.Println(products.Format())
	// using the fmt.Stringer implementation
	fmt.Println(products)

	kettle := Product{101, "Kettle", 2500, 10, "Utencil"}
	fmt.Println("Index of kettle :", products.IndexOf(kettle))

	fmt.Println("Costly Products")
	costlyProductPredicate := func(p Product) bool {
		return p.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	// fmt.Println(costlyProducts.Format())
	fmt.Println(costlyProducts)
}
