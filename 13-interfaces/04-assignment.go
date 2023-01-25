package main

import (
	"fmt"
	"sort"
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
Write the sort method for the Products using which the user can sort the products by any attribute
	- By Id
	- By Name
	- By Cost
	etc

IMPORTANT: using sort.Sort() function



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

func (products Products) sortByField(field string) {
	switch field {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	default:
		sort.Sort(ByCost{products})
	}
}

/* sort.Interface implementation */
func (a Products) Len() int {
	return len(a)
}

func (a Products) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Products) Less(i, j int) bool {
	return a[i].Id < a[j].Id
}

type ByName struct {
	Products
}

func (a ByName) Less(i, j int) bool {
	return a.Products[i].Name < a.Products[j].Name
}

type ByCost struct {
	Products
}

func (a ByCost) Less(i, j int) bool {
	return a.Products[i].Cost < a.Products[j].Cost
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

	fmt.Println("Sort By Id")
	products.sortByField("Id")
	fmt.Println(products)

	fmt.Println("Sort By Name")
	products.sortByField("Name")
	fmt.Println(products)

	fmt.Println("Sort By Cost")
	products.sortByField("Cost")
	fmt.Println(products)
}
