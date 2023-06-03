package numbers

import "fmt"

// Number Names :
// Show how to spell out a number in English.
// You can use a preexisting implementation or roll your own, but you should support inputs up to at least one million
// (or the maximum value of your language's default bounded integer type, if that's less).
// Optional: Support for inputs other than positive integers (like zero, negative integers, and floating-point numbers).
func NumberNames() {
	var (
		number     int
		numberName string
	)

	fmt.Printf("Enter a number to show it's name: ")
	fmt.Scanln(&number)

	numberName = spellNumber(number)

	fmt.Printf("(%d)'s name is: %s\n", number, numberName)
}

func spellNumber(n int) string {
	if name, ok := numNames[n]; ok {
		return name
	} else if n < 100 {
		return fmt.Sprintf("%s-%s", numNames[n/10*10], numNames[n%10])
	} else if n < 1000 {
		var suffix string
		if n%100 != 0 {
			suffix = " " + spellNumber(n%100)
		}
		return fmt.Sprintf("%s hundred%s", numNames[n/100], suffix)
	} else if n < 1000000 {
		var suffix string
		if n%1000 != 0 {
			suffix = " " + spellNumber(n%1000)
		}
		return fmt.Sprintf("%s thousand%s", spellNumber(n/1000), suffix)
	} else {
		return "number out of range"
	}
}

func main() {
	for i := 1; i <= 1000000; i++ {
		fmt.Println(spellNumber(i))
	}
}

var numNames = map[int]string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}
