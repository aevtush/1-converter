package main

import (
	"errors"
	"fmt"
)

const USDtoEUR = 0.85
const USDtoRUB = 80.00
const EURtoRUB = USDtoRUB / USDtoEUR

func main() {
	fmt.Println("****__Конвертер валют__****")
	convert(userInput())
}

func convert(baseCurrency string, amount float64, targetCurrency string) error {
	return errors.New("The function is not implemented yet")
}

// func convert(baseCurrency string, amount float64, targetCurrency string) {
// 	if baseCurrency == "EUR" && targetCurrency == "RUB" {
// 		result := USDtoRUB / USDtoEUR * amount
// 		fmt.Print(result)
// 	} else if baseCurrency == "USD" && targetCurrency == "EUR" {
// 		result := USDtoEUR * amount
// 		fmt.Print(result)
// 	} else if baseCurrency == "USD" && targetCurrency == "RUB" {
// 		result := USDtoRUB * amount
// 		fmt.Print(result)
// 	} else {
// 		fmt.Println("Error: indicated convertation is unavalable")
// 	}
// }

func userInput() (string, float64, string) {
	var baseCurrency string
	var amount float64
	var targetCurrency string
	fmt.Print("Укажите исходную валюту :")
	fmt.Scan(&baseCurrency)
	fmt.Print("Укажите сумму в исходной валюте :")
	fmt.Scan(&amount)
	fmt.Print("Укажите валюту для конвертации :")
	fmt.Scan(&targetCurrency)
	return baseCurrency, amount, targetCurrency
}
