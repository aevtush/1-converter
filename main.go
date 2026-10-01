package main

import "fmt"

func main() {
	const USDtoEUR = 0.85
	const USDtoRUB = 80.00
	const EURtoRUB = USDtoRUB / USDtoEUR
	fmt.Println(EURtoRUB)
}
