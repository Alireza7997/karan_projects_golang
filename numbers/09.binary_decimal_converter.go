package numbers

import (
	"fmt"
	"strconv"
	"strings"
)

// Binary to Decimal and Back Converter :
// Develop a converter to convert a decimal number to binary or a binary number to its decimal equivalent.
func BinaryDecimalConverter() {
	var (
		convertType     string
		chosenType      string
		number          int
		convertedNumber int
	)

	fmt.Print("Enter the number type you want to convert. Binary(b), Decimal(d): ")
	fmt.Scanln(&convertType)
	convertType = strings.ToLower(convertType)
	if convertType != "b" && convertType != "d" {
		fmt.Println("Please enter a valid number type(b/d).")
		BinaryDecimalConverter()
		return
	}
	if convertType == "b" {
		chosenType = "Binary"
	} else {
		chosenType = "Decimal"
	}

	switch convertType {
	case "b":
		fmt.Printf("Enter a %s number: ", chosenType)
		fmt.Scanln(&number)

		convertedNumber = BinaryToDecimal(number)
		fmt.Printf("%d converted to decimal is : %d", number, convertedNumber)
		fmt.Println()

	case "d":
		fmt.Printf("Enter a %s number: ", chosenType)
		fmt.Scanln(&number)

		convertedNumber = DecimalToBinary(number)
		fmt.Printf("%d converted to binary is : %d", number, convertedNumber)
		fmt.Println()
	}
}

func DecimalToBinary(decimal int) (binary int) {
	if decimal == 0 {
		return 0
	}

	converting := ""
	for decimal > 0 {
		remainder := decimal % 2
		converting = fmt.Sprintf("%d%s", remainder, converting)
		decimal /= 2
	}
	binary, _ = strconv.Atoi(converting)

	return
}

func BinaryToDecimal(binary int) (decimal int) {
	binaryStr := fmt.Sprintf("%d", binary)
	power := 1

	for i := len(binaryStr) - 1; i >= 0; i-- {
		if binaryStr[i] == '1' {
			decimal += power
		}
		power *= 2
	}

	return
}
