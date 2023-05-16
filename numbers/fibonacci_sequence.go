package numbers

import "fmt"

// Fibonacci Sequence :
// Enter a number and have the program generate the Fibonacci sequence to that number or to the Nth number.
func Fibonacci() {
	var (
		n     int
		i     int = 2
		limit int = 999999999
	)

	fmt.Printf("Enter a number to generate the Fibonacci sequence to that number (limit : %d): ", limit)
	fmt.Scan(&n)
	sequence := make([]int, n)

	if n < 0 || n > 999999999 {
		fmt.Printf("the number must be between 0 and %d", limit)
		fmt.Println()
		Fibonacci()
	} else {
		sequence[0] = 0
		if n == 0 {
			fmt.Println(sequence)
		}
		sequence[1] = 1
		if n == 1 {
			fmt.Println(sequence)

		}
		for {
			sequence[i] = sequence[i-2] + sequence[i-1]

			if sequence[i] == n {
				fmt.Println(sequence)
				break
			} else if sequence[i] > n {
				fmt.Println(sequence[:i])
				break
			}
			i++
		}
	}
}
