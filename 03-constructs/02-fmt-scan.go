package main

import "fmt"

func main() {
	var firstName, lastName string
	fmt.Print("Enter your name (firstname lastname):")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s, %s  Have a nice day!\n", lastName, firstName)

	var no int
	fmt.Print("Enter a number : ")
	fmt.Scanf("%d\n", &no)
	fmt.Println("The given number is :", no)
}
