package numbers

import (
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
)

// Limit Calculator :
// Ask the user to enter f(x) and the limit value, then return the value of the limit statement.
func LimitCalculator() {
	var fx string
	var limitValue string

	fmt.Println("Enter the function f(x):")
	fmt.Scanln(&fx)

	fmt.Println("Enter the limit value:")
	fmt.Scanln(&limitValue)
	limitValue = strings.ToLower(limitValue)

	// Replace 'x' in the function with the limit value
	fx = strings.ReplaceAll(fx, "x", limitValue)

	// Create a new expression with the function
	expression, err := govaluate.NewEvaluableExpression(fx)
	if err != nil {
		fmt.Println("Error in parsing the function:", err)
		return
	}

	// Evaluate the expression
	result, err := expression.Evaluate(nil)
	if err != nil {
		fmt.Println("Error in evaluating the function:", err)
		return
	}

	fmt.Println("The value of the limit statement is:", result)
}
