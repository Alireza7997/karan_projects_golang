package numbers

import (
	"fmt"
	"strings"
)

// Next Prime Number :
// Have the program find prime numbers until the user chooses to stop asking for the next one.

var (
	primeNumbers []int
	lastPrime    int = 0
)

func NextPrime() {
	var answer string

	fmt.Print("Generate prime number? (y/n): ")
	fmt.Scanln(&answer)
	answer = strings.ToLower(answer)

	if answer == "y" {
		for i := lastPrime; ; i++ {
			if isPrime(i) {
				primeNumbers = append(primeNumbers, i)
				lastPrime = i + 1
				break
			}
		}
		fmt.Println(primeNumbers)
		NextPrime()
	} else {
		fmt.Printf("Finished, prime numbers: %v", primeNumbers)
	}
}
