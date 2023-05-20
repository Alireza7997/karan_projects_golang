package numbers

import "fmt"

// Change Return Program :
// The user enters a cost and then the amount of money given.
// The program will figure out the change and the number of quarters, dimes, nickels, pennies needed for the change.

func ChangeReturn() {
	var (
		money     float64
		cost      float64
		change    float64
		cents     int
		quarters  int
		dimes     int
		nickels   int
		pennies   int
		remaining int
	)

	fmt.Print("Enter the amount of money you have($): ")
	fmt.Scanln(&money)

	fmt.Print("Enter the cost($): ")
	fmt.Scanln(&cost)

	change = money - cost
	cents = int(change * 100)
	quarters = cents / 25
	remaining = cents % 25
	dimes = remaining / 10
	remaining = remaining % 10
	nickels = remaining / 5
	pennies = remaining % 5

	fmt.Printf("Total change : %.1f\n%d quarters, %d dimes, %d nickels, %d pennies", change, quarters, dimes, nickels, pennies)
}
