package main

import "fmt"

func main() {
	// var productRanks map[string]int = make(map[string]int)
	// var productRanks map[string]int = map[string]int{"Pen": 5, "Pencil": 1}
	/*
		var productRanks map[string]int = map[string]int{
			"Pen":    5,
			"Pencil": 1,
		}
	*/
	/*
		var productRanks = map[string]int{
			"Pen":    5,
			"Pencil": 1,
		}
	*/

	productRanks := map[string]int{
		"Pen":    5,
		"Pencil": 1,
	}
	productRanks["Marker"] = 3
	fmt.Println(productRanks)

	fmt.Println("Iterating a map (using range)")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	keyToCheck := "Pencil"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %v\n", keyToCheck, rank)
	} else {
		fmt.Printf("key %q does not exist\n", keyToCheck)
	}

	keyToDelete := "Notepad"
	delete(productRanks, keyToDelete)
	fmt.Printf("After deleting %q\n", keyToDelete)
	fmt.Println(productRanks)

}
