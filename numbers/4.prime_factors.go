package numbers

import (
	"fmt"
	"math"
)

// Prime Factorization :
// Have the user enter a number and find all Prime Factors (if there are any) and display them.
func PrimeFactors() {

	var number int
	fmt.Println("Enter a number to get prime factors : ")
	fmt.Scan(&number)

	var divisionNumbers []int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			divisionNumbers = append(divisionNumbers, i)
		}
	}

	var primeFactors []int
	for _, factor := range divisionNumbers {
		if isPrime(factor) {
			primeFactors = append(primeFactors, factor)
		}
	}

	fmt.Println(primeFactors)

}

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}

	maxDivistor := int(math.Sqrt(float64(number)))
	for i := 2; i <= maxDivistor; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
