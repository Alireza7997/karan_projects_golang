package numbers

import (
	"fmt"
	"strings"
)

// Unit Converter (temp, currency, volume, mass and more) :
// The user enters the type of unit being entered, the type of unit they want to convert to and then the value.
// The program will then make the conversion.
func UnitConverter() {
	var (
		unit                  string
		availableTemperatures = []string{"\033[31mCelsius(c)", "Fahrenheit(f)", "Kelvin(k)\033[0m"}
		availableCurrencies   = []string{"\033[32mDollars", "Euros", "Pounds", "Yen\033[0m"}
		availableVolumes      = []string{"\033[34mLitre", "Gallons", "Mililitres\033[0m"}
		availableMasses       = []string{"\033[35mKilograms", "Pounds", "Grams\033[0m"}
		mainScale             string
		targetScale           string
		amount                float64
	)

	fmt.Print("\033[33mAvailable conversions : \033[0m")
	fmt.Print("\033[31mtemperature(t)\033[0m")
	fmt.Print("\033[33m, \033[32mcurrency(c)\033[0m")
	fmt.Print("\033[33m, \033[34mvolume(v)\033[0m")
	fmt.Print("\033[33m, \033[35mmass(m)\033[0m\n")
	fmt.Print("Which type of unit conversion do you want to proceed?\n")
	fmt.Scanln(&unit)
	unit = strings.ToLower(unit)

	fmt.Println("Please Enter your primary scale and the scale you want to convert to:")

	switch unit {
	case "temperature", "t":
		fmt.Println("Available temperatures: ", availableTemperatures)
		fmt.Print("\033[1;36mPrimary Scale:\033[0m ")
		fmt.Scanln(&mainScale)
		mainScale = strings.ToLower(mainScale)
		fmt.Print("\033[1;36mTarget Scale:\033[0m ")
		fmt.Scanln(&targetScale)
		targetScale = strings.ToLower(targetScale)
		fmt.Print("\033[32mAmount:\033[0m ")
		fmt.Scanln(&amount)

		converted := temperatureConvertion(mainScale, targetScale, amount)
		fmt.Println("\033[1;32mConverted:\033[0m ", converted)

	case "currency", "c":
		fmt.Println("Available currencies: ", availableCurrencies)
		fmt.Print("\033[1;36mPrimary Scale:\033[0m ")
		fmt.Scanln(&mainScale)
		mainScale = strings.ToLower(mainScale)
		fmt.Print("\033[1;36mTarget Scale:\033[0m ")
		fmt.Scanln(&targetScale)
		targetScale = strings.ToLower(targetScale)
		fmt.Print("\033[32mAmount:\033[0m ")
		fmt.Scanln(&amount)
	case "volume", "v":
		fmt.Println("Available volumes: ", availableVolumes)
		fmt.Print("\033[1;36mPrimary Scale:\033[0m ")
		fmt.Scanln(&mainScale)
		mainScale = strings.ToLower(mainScale)
		fmt.Print("\033[1;36mTarget Scale:\033[0m ")
		fmt.Scanln(&targetScale)
		targetScale = strings.ToLower(targetScale)
		fmt.Print("\033[32mAmount:\033[0m ")
		fmt.Scanln(&amount)
	case "mass", "m":
		fmt.Println("Available masses: ", availableMasses)
		fmt.Print("\033[1;36mPrimary Scale:\033[0m ")
		fmt.Scanln(&mainScale)
		mainScale = strings.ToLower(mainScale)
		fmt.Print("\033[1;36mTarget Scale:\033[0m ")
		fmt.Scanln(&targetScale)
		targetScale = strings.ToLower(targetScale)
		fmt.Print("\033[32mAmount:\033[0m ")
		fmt.Scanln(&amount)
	default:
		fmt.Println("Please enter a valid type of conversion")
		UnitConverter()
	}

}

func temperatureConvertion(main, target string, amount float64) (converted float64) {
	switch main {
	case "celsius", "c":
		switch target {
		case "celsius", "c":
			converted = amount
		case "fahrenheit", "f":
			converted = (amount * 9 / 5) + 32
		case "kelvin", "k":
			converted = amount + 273.15
		default:
			fmt.Printf("unsupported conversion to %s", target)
			fmt.Println()
		}
	case "fahrenheit", "f":
		switch target {
		case "celsius", "c":
			converted = (amount - 32) * 5 / 9
		case "fahrenheit", "f":
			converted = amount
		case "kelvin", "k":
			converted = (amount-32)*5/9 + 273.15
		default:
			fmt.Printf("unsupported conversion to %s", target)
			fmt.Println()
		}
	case "kelvin", "k":
		switch target {
		case "celsius", "c":
			converted = amount - 273.15
		case "fahrenheit", "f":
			converted = ((amount - 273.15) * 9 / 5) + 32
		case "kelvin", "k":
			converted = amount
		default:
			fmt.Printf("unsupported conversion to %s", target)
			fmt.Println()
		}
	default:
		fmt.Printf("unsupported conversion from %s", main)
		fmt.Println()
	}

	return
}
