package numbers

import (
	"fmt"
)

// Fast Exponentiation :
// Ask the user to enter 2 integers a and b and output a^b (i.e. pow(a,b)) in O(lg n) time complexity.
func FastExponentiation() {
	var (
		a      int
		b      int
		result int
	)

	fmt.Printf("Enter a :")
	fmt.Scanln(&a)

	fmt.Printf("Enter b :")
	fmt.Scanln(&b)

	result = lognPow(a, b)

	fmt.Printf("%d^%d = %d\n", a, b, result)

}

func lognPow(a, b int) int {
	if b == 0 {
		return 1
	}

	temp := lognPow(a, b/2)
	if b%2 == 0 {
		return temp * temp
	} else {
		return a * temp * temp
	}
}
