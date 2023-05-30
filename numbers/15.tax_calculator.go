package numbers

import "fmt"

// Tax Calculator :
// Asks the user to enter a cost and either a country or state tax. It then returns the tax plus the total cost with tax.

func TaxCalculator() {
	var (
		cost      float64
		taxRate   float64
		taxAmount float64
		totalCost float64
	)

	fmt.Printf("Enter a cost: ")
	fmt.Scanln(&cost)

	fmt.Printf("Enter the tax rate (e.g. 0.09): ")
	fmt.Scanln(&taxRate)

	taxAmount = cost * taxRate
	totalCost = cost + taxAmount

	fmt.Printf("\033[31mTax Amount: \033[0m%.2f\n\033[32mTotal Amount: \033[0m%.2f\n", taxAmount, totalCost)
}
