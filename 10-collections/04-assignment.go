/* create a function that returns the prime numbers between given start and end
print the prime numbers in the main function
*/

package main

import "fmt"

func main() {
	primeNos := genPrimeNos(3, 100)
	fmt.Println(primeNos)
}

func genPrimeNos(start, end int) []int {
	result := []int{}
	for i := start; i <= end; i++ {
		if isPrime(i) {
			result = append(result, i)
		}
	}
	return result
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
