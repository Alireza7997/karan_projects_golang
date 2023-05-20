package numbers

import (
	"fmt"
	"strings"
)

// Mortgage Calculator :
// Calculate the monthly payments of a fixed term mortgage over given Nth terms at a given interest rate.
// Also figure out how long it will take the user to pay back the loan
// For added complexity, add an option for users to select the compounding interval (Monthly, Weekly, Daily, Continually).
func MortgageCalculator() {
	var (
		interval           string
		intervalType       string
		mortgageAmount     int32
		maxAmount          int32 = 10000000 // you can set maximum amount for the loan
		interestRate       float32
		totalPaybackAmount float32
		paybackYears       int32
		maximumPaybackTime int32 = 30
		numberOfPayments   int32
		eachPaymentAmount  float32
	)

	fmt.Print("Please enter the amount you wish to borrow for your mortgage. ")
	fmt.Scanln(&mortgageAmount)
	if mortgageAmount > maxAmount {
		fmt.Printf("The maximum amount you can borrow is : %d", maxAmount)
		fmt.Println()
		MortgageCalculator()
	}

	fmt.Print("How many years do you want to take to pay back the mortgage(max 30)? ")
	fmt.Scan(&paybackYears)
	if paybackYears > maximumPaybackTime {
		fmt.Printf("The maximum pay back time is : %d", maximumPaybackTime)
		fmt.Println()
		return
	}

	fmt.Print("Enter the interest rate(%): ")
	fmt.Scanln(&interestRate)
	totalPaybackAmount += float32(mortgageAmount) * interestRate

	fmt.Print("Please choose the payment interval you prefer - Monthly(m), Weekly(w) or Daily(d): ")
	fmt.Scanln(&interval)
	interval = strings.ToLower(interval)
	if interval != "m" && interval != "w" && interval != "d" {
		fmt.Println("The payment interval you entered is not valid. Please choose a valid interval (m/w/d)")
		fmt.Println()
		return
	}
	switch interval {
	case "m":
		intervalType = "month"
		numberOfPayments = paybackYears * 12
	case "w":
		intervalType = "week"
		numberOfPayments = paybackYears * 360 / 7
	case "d":
		intervalType = "day"
		numberOfPayments = paybackYears * 360
	}

	eachPaymentAmount = totalPaybackAmount / float32(numberOfPayments)
	fmt.Printf("You have to pay $%.2f each %s", eachPaymentAmount, intervalType)
	fmt.Println()
	fmt.Printf("It takes you %d %ss to pay the loan back", numberOfPayments, intervalType)
}
