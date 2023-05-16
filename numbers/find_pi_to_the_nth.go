package numbers

import (
	"fmt"
)

// Find PI to the Nth Digit :
// Enter a number and have the program generate PI up to that many decimal places. Keep a limit to how far the program will go.
func FindPi() {
	var (
		nthDigit     int
		pi           string = "3.141592653589793238462643383279502884197169399375105820974944592307816406286208998628034825342117067982148086513282306647093844609550582231725359408128481117450284102701938521105559644622948954930381964428810975665933446128475648233786783165271201909"
		piToNthDigit string
		limit        int = 250
	)

	fmt.Printf("Enter 'n' to generate 'Pi' to  the nth digit(up to %d) : ", limit)
	fmt.Scan(&nthDigit)

	if nthDigit < 1 || nthDigit > limit {
		fmt.Printf("'n' must be from 1 to %d", limit)
		fmt.Println()
		FindPi()
	} else {
		for _, digit := range pi[:nthDigit+1] {
			piToNthDigit = piToNthDigit + string(digit)
		}

		fmt.Println(piToNthDigit)
	}

}
