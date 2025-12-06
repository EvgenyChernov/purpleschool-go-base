package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower float64 = 2
	userHeight := 1.8
	var userKg float64
	userKg = 100
	IMT := userKg / math.Pow(userHeight, IMTPower)
	const USD float64 = 132
	const EUR float64 = 0.86
	const RUB float64 = 76.80
	USDtoEUR := USD * EUR
	USDtoRUB := USD * RUB

	fmt.Print(USDtoEUR, USDtoRUB, IMT)
}
