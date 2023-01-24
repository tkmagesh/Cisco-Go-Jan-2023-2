package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating a slice (using indexer)")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println("Iterating a slice (using range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
		// fmt.Println(val)
	}

	//adding new values
	nos = append(nos, 10)
	fmt.Println(nos)

	nos = append(nos, 20, 30, 40)
	fmt.Println(nos)

	hundreds := []int{100, 200, 300}
	// nos = append(nos, hundreds[0], hundreds[1], hundreds[2])
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	//slicing
	/*
		[lo:hi] => from index "lo" to "hi-1"
		[:hi] => from index 0 to "hi-1"
		[lo:] => from index "lo" to end
	*/

	fmt.Println("nos[2:5] =", nos[2:5])
	fmt.Println("nos[2:] =", nos[2:])
	fmt.Println("nos[:5] =", nos[:5])

	fmt.Println("original list :", nos)
	sort(nos)
	fmt.Println("after sorting, list :", nos)

	/*
		subset := nos[2:5]
		subset[0] = 10000
		fmt.Println("subset = ", subset)
		fmt.Println("nos = ", nos)
	*/

	//to create a copy
	fmt.Println("Creating a copy of the slice")
	subset := make([]int, 3)
	copy(subset, nos[2:5])
	subset[0] = 10000
	fmt.Println("subset = ", subset)
	fmt.Println("nos = ", nos)

}

func sort(list []int) /* do not return anything */ {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
