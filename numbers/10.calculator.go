package numbers

import (
	"fmt"
	"math"
)

// Calculator :
// A simple calculator to do basic operators. Make it a scientific calculator for added complexity.
func Calculator() {
	var (
		operands = make([]float64, 2)
		operator string
		result   float64
	)

	fmt.Println("Calculator - Operators:(+, -, ÷, ×, %, √, ^, ^2, sin, cos, tan)\nEnter calculation (e.g. 3 + 3, 3 %):")
	fmt.Scanln(&operands[0], &operator, &operands[1])

	switch operator {
	case "+":
		result = operands[0] + operands[1]
	case "-":
		result = operands[0] - operands[1]
	case "×", "*":
		result = operands[0] * operands[1]
	case "÷", "/":
		result = operands[0] / operands[1]
	case "%":
		result = operands[0] / 100
	case "√":
		result = math.Sqrt(operands[0])
	case "^", "**":
		result = math.Pow(operands[0], operands[1])
	case "^2", "**2":
		result = operands[0] * operands[0]
	case "sin":
		result = math.Sin(operands[0])
	case "cos":
		result = math.Cos(operands[0])
	case "tan":
		result = math.Tan(operands[0])
	default:
		fmt.Println("Please enter a valid operation")
		Calculator()
	}

	fmt.Println("Result: ", result)
}
