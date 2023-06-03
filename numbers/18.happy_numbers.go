package numbers

import (
	"fmt"
)

// Happy Numbers :
// A happy number is defined by the following process:
// Starting with any positive integer, replace the number by the sum of the squares of its digits,
// and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy numbers, while those that do not end in 1 are unhappy numbers
// Display an example of your output here. Find first 8 happy numbers.
func HappyNumbers() {
	var (
		baseNumber      int
		number          int
		isHappyNumber   bool = false
		repetition      uint8
		repetitionLimit uint8 = 30
	)

	fmt.Print("Enter a number: ")
	fmt.Scanln(&baseNumber)
	number = baseNumber

	for {
		if repetition > repetitionLimit {
			break
		}
		if number == 1 {
			isHappyNumber = true
			break
		}

		numberStr := fmt.Sprintf("%d", number)
		number = 0
		for _, n := range numberStr {
			number += int(n-'0') * int(n-'0')
		}

		repetition++
	}

	fmt.Printf("%d is happy: %v\n", baseNumber, isHappyNumber)

}
