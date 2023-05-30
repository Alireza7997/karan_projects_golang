package numbers

import (
	"fmt"
	"strconv"
	"strings"
)

func CreditCardValidator() {
	var (
		creditCardNumber string
		isCardValid      bool
		result           string = "Not Valid"
	)

	fmt.Print("Please enter your card number: ")
	fmt.Scanln(&creditCardNumber)

	if isCardValid = checksum(creditCardNumber); isCardValid {
		result = "Valid"
	}

	fmt.Printf("Card number %v is: %v\n", creditCardNumber, result)
}

func checksum(cardNumber string) bool {
	cardDigits := strings.ReplaceAll(cardNumber, " ", "")
	length := len(cardDigits)
	sum := 0

	for i := length - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(cardDigits[i]))
		if err != nil {
			return false
		}

		if (length-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}
