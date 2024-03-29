/* constants */
package main

import "fmt"

func main() {
	// const pi float64 = 3.14
	// const pi = 3.14

	/*
		const pi = 31.4
		const app_version = "1.0.0"
	*/

	const (
		pi          = 31.4
		app_version = "1.0.0"
	)

	/* iota */
	/*
		const red int = 0
		const green int = 1
		const blue int = 2
	*/

	/*
		const (
			red   int = iota
			green int = iota
			blue  int = iota
		)
	*/

	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota + 2
			green
			blue
		)
	*/

	/*
		const (
			red = (iota * 2) + 1
			green
			blue
		)
	*/

	const (
		red = (iota * 2) + 1
		green
		_
		blue
	)
	fmt.Printf("red = %d, green = %d and blue = %d\n", red, green, blue)

	fmt.Println("iota applied")
	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)
}
