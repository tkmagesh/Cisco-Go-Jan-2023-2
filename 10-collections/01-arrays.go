package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating an array (using indexer)")
	for idx := 0; idx < 5; idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println("Iterating an array (using range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
		// fmt.Println(val)
	}

	list1 := [3]string{"Pen", "Pencil", "Marker"}
	list2 := [3]string{"Pen", "Pencil", "Marker"}
	fmt.Printf("&list1 = %p, &list2 = %p\n", &list1, &list2)

	fmt.Println(list1 == list2)

	n1 := 100
	n2 := 100
	fmt.Printf("&n1 = %p, &n2 = %p\n", &n1, &n2)
	fmt.Println(n1 == n2)

	sort(&nos)
	fmt.Println("After sorting, nos = ", nos)

}

func sort(list *[5]int) /* do not return anything */ {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			/*
				if (*list)[i] > (*list)[j] {
					(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
				}
			*/
			if list[i] > list[j] { // accessing the elements using the pointer is allowed
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
