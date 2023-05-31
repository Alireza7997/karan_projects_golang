package numbers

import (
	"fmt"
	"math"
)

// complex Number Algebra :
// Show addition, multiplication, negation, and inversion of complex numbers in separate functions.
// (Subtraction and division operations can be made with pairs of these operations.)
// Print the results for each operation tested.
func ComplexAlgera() {
	var (
		availableOperations  = []string{"Addition", "Subtraction", "Multiplication", "Negation", "Inversion"}
		complex1             complex
		complex2             complex
		additionResult       complex
		subtractionResult    complex
		multiplicationResult complex
		negationResult1      complex
		negationResult2      complex
		inversionResult1     complex
		inversionResult2     complex
	)
	fmt.Printf("\033[033mAvailable operations: %v\033[0m\n", availableOperations)

	fmt.Printf("Enter a complex number(e.g. 2 -3): ")
	fmt.Scanln(&complex1.Real, &complex1.Imaginary)

	fmt.Printf("Enter the second complex number: ")
	fmt.Scanln(&complex2.Real, &complex2.Imaginary)

	additionResult = complex1.Add(complex2)
	subtractionResult = complex1.Subtract(complex2)
	multiplicationResult = complex1.Multiply(complex2)
	negationResult1 = complex1.Negate()
	negationResult2 = complex2.Negate()
	inversionResult1 = complex1.Invert()
	inversionResult2 = complex2.Invert()

	green := "\033[32m"
	reset := "\033[0m"
	fmt.Printf("Addition of %v and %v: %v%s%s\n", complex1, complex2, green, additionResult, reset)
	fmt.Printf("Subtraction of %v and %v: %v%s%s\n", complex1, complex2, green, subtractionResult, reset)
	fmt.Printf("Multiplication of %v and %v: %v%s%s\n", complex1, complex2, green, multiplicationResult, reset)
	fmt.Printf("Negation of %v: %v%s%s\n", complex1, green, negationResult1, reset)
	fmt.Printf("Negation of %v: %v%s%s\n", complex2, green, negationResult2, reset)
	fmt.Printf("Inversion of %v: %v%s%s\n", complex1, green, inversionResult1, reset)
	fmt.Printf("Inversion of %v: %v%s%s\n", complex2, green, inversionResult2, reset)

}

type complex struct {
	Real      float64
	Imaginary float64
}

func (z1 complex) Add(z2 complex) complex {
	return complex{z1.Real + z2.Real, z1.Imaginary + z2.Imaginary}
}

func (z1 complex) Subtract(z2 complex) complex {
	return complex{z1.Real - z2.Real, z1.Imaginary - z2.Imaginary}
}

func (z1 complex) Multiply(z2 complex) complex {
	real := z1.Real*z2.Real - z1.Imaginary*z2.Imaginary
	imaginary := z1.Real*z2.Imaginary + z1.Imaginary*z2.Real
	return complex{real, imaginary}
}

func (z complex) Negate() complex {
	return complex{-z.Real, -z.Imaginary}
}

func (z complex) Invert() complex {
	norm := z.Real*z.Real + z.Imaginary*z.Imaginary
	if norm == 0 {
		panic("Cannot invert a complex number with zero norm.")
	}
	return complex{z.Real / norm, -z.Imaginary / norm}
}

func (z complex) String() string {
	sign := "+"
	if z.Imaginary < 0 {
		sign = "-"
	}
	return fmt.Sprintf("%v %s %vi", z.Real, sign, math.Abs(z.Imaginary))
}
