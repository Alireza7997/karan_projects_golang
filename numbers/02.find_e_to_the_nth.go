package numbers

import "fmt"

// Find e to the Nth Digit :
// Enter a number and have the program generate e up to that many decimal places. Keep a limit to how far the program will go.
func FindE() {
	var (
		nthDigit    int
		e           string = "2.718281828459045235360287471352662497757247093699959574966967627724076630353547594571382178525166427427466391932003059921817413596629043572900334295260595630738132328627943490763233829880753195251019011573834187930702154089149934884167509244761460668"
		eToNthDigit string
		limit       int = 250
	)

	fmt.Printf("Enter 'n' to generate 'e' to  the nth digit(up to %d) : ", limit)
	fmt.Scan(&nthDigit)

	if nthDigit < 1 || nthDigit > limit {
		fmt.Printf("'n' must be from 1 to %d", limit)
		fmt.Println()
		FindE()
	} else {
		for _, digit := range e[:nthDigit+1] {
			eToNthDigit = eToNthDigit + string(digit)
		}

		fmt.Println(eToNthDigit)
	}

}
