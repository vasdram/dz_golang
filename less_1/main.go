package main

import "fmt"

func main() {
	const USD_TO_EUR = 1.13
	const USD_TO_RUB = 63.94
	const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR

	fmt.Println(EUR_TO_RUB)
}

func outputCurrency() {
	var currency int
	fmt.Println("input sum of currency")
	fmt.Scanf("%d", &currency)
}

func calcaleteCurrency(rate float64, from, to string, ) float64 {
}