package numbers

import "fmt"

// Factorial Finder :
// The Factorial of a positive integer, n, is defined as the product of the sequence n, n-1, n-2, ...1 and the factorial of zero, 0, is defined as being 1.
// Solve this using both loops and recursion.
func FactorialFinder() {
	var (
		number    int
		factorial int = 1
	)

	fmt.Printf("Enter a number to calculate the factorial: ")
	fmt.Scanln(&number)

	if number == 0 {
		fmt.Printf("Factorial of %d is %d\n", number, 0)
	} else {
		for i := number; i > 0; i-- {
			factorial *= i
		}

		fmt.Printf("Factorial of %d is %d (by loop)\n", number, factorial)
	}

	factorial = recursiveFactorial(number)
	fmt.Printf("Factorial of %d is %d (by recursion)\n", number, factorial)
}

func recursiveFactorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * recursiveFactorial(n-1)
}
