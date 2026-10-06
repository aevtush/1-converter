package main

import (
	"fmt"
)

const USDtoEUR = 0.85
const USDtoRUB = 80.00
const EURtoRUB = USDtoRUB / USDtoEUR

func main() {
	fmt.Println("****__Currency converter__****")
	for {
		baseCurrency := userInputBaseCurrency()
		amount := userInputAmount()
		quoteCurrency := userInputQuoteCurrency(baseCurrency)
		result := exchange(baseCurrency, amount, quoteCurrency)
		fmt.Println("Calculated amount is:", result)
		if !checkRepeatCalculation() {
			break
		}
	}
	println("Thanks for using the app!")
}

func exchange(baseCurrency string, amount float64, quoteCurrency string) float64 {
	switch baseCurrency {
	case "USD", "usd":
		switch quoteCurrency {
		case "EUR", "eur":
			return amount * USDtoEUR
		case "RUB", "rub":
			return amount * USDtoRUB
		default:
			return 0
		}
	case "EUR", "eur":
		switch quoteCurrency {
		case "USD", "usd":
			return amount / USDtoEUR
		case "RUB", "rub":
			return amount * EURtoRUB
		default:
			return 0
		}
	case "RUB", "rub":
		switch quoteCurrency {
		case "USD", "usd":
			return amount / USDtoRUB
		case "EUR", "eur":
			return amount / EURtoRUB
		default:
			return 0
		}
	default:
		return 0
	}

}

func userInputBaseCurrency() string {
	fmt.Println("Enter base currency (USD|EUR|RUB):")
	for {
		var baseCurrency string
		fmt.Scan(&baseCurrency)
		switch baseCurrency {
		case "USD", "usd", "EUR", "eur", "RUB", "rub":
			return baseCurrency
		default:
			fmt.Println("Invalid currency, enter USD|EUR|RUB:")
		}
	}
}

func userInputAmount() float64 {
	fmt.Println("Enter amount to convert:")
	for {
		var amount float64
		fmt.Scan(&amount)
		switch {
		case amount <= 0:
			fmt.Println("Invalid amount, enter strictly positive number, no spacing allowed:")
		default:
			return amount
		}
	}
}

func userInputQuoteCurrency(baseCurrency string) string {
	switch baseCurrency {
	case "USD", "usd":
		fmt.Println("Enter quote currency (EUR|RUB):")
		for {
			var quoteCurrency string
			fmt.Scan(&quoteCurrency)
			switch quoteCurrency {
			case "EUR", "eur", "RUB", "rub":
				return quoteCurrency
			default:
				fmt.Println("Invalid currency, enter EUR|RUB:")
			}
		}
	case "EUR", "eur":
		fmt.Println("Enter quote currency (USD|RUB):")
		for {
			var quoteCurrency string
			fmt.Scan(&quoteCurrency)
			switch quoteCurrency {
			case "USD", "usd", "RUB", "rub":
				return quoteCurrency
			default:
				fmt.Println("Invalid currency, enter USD|RUB:")
			}
		}
	case "RUB", "rub":
		fmt.Println("Enter quote currency (USD|EUR):")
		for {
			var quoteCurrency string
			fmt.Scan(&quoteCurrency)
			switch quoteCurrency {
			case "USD", "usd", "EUR", "eur":
				return quoteCurrency
			default:
				fmt.Println("Invalid currency, enter USD|EUR:")
			}
		}
	default:
		return ""
	}
}

func checkRepeatCalculation() bool {
	fmt.Println("Would you like to calculate again? (y/n): ")
	for {
		var userChoice string
		fmt.Scan(&userChoice)
		switch userChoice {
		case "y", "Y":
			return true
		case "n", "N":
			return false
		default:
			fmt.Println("Invalid input, enter y or n: ")
		}
	}
}
