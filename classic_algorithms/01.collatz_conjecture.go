package classic_algorithms

import "fmt"

// Collatz Conjecture :
// Start with a number n > 1. Find the number of steps it takes to reach one using the following process:
// If n is even, divide it by 2. If n is odd, multiply it by 3 and add 1.
func CollatzConjecture() {
	var (
		number uint
		steps  uint16
	)

	fmt.Printf("Enter a number(n > 1): ")
	_, err := fmt.Scanln(&number)
	if err != nil || number < 1 {
		fmt.Println("Invalid input. Please enter a number greater than 1.")
		CollatzConjecture()
		return
	}

	for {
		if number == 1 {
			break
		}

		if number%2 == 0 {
			number /= 2
			steps++
			// fmt.Printf("step %d: %d\n", steps, number)
		} else if number%2 != 0 {
			number = number*3 + 1
			steps++
			// fmt.Printf("step %d: %d\n", steps, number)
		} else if number%2 != 0 {
		}
	}

	fmt.Printf("Number of steps to reach 1: %d\n", steps)
}
