package main

import "fmt"

func main() {
	const USD_TO_EUR = 1.13
	const USD_TO_RUB = 63.94
	EUR_TO_RUB := USD_TO_RUB / USD_TO_EUR

	fmt.Println(EUR_TO_RUB)
}