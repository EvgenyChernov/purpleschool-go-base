package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower float64 = 2
	var userKg float64
	var userHeight float64
	fmt.Print("___Калькулятор индекса массы тела___ \n")
	fmt.Print("Введите свой рост в м: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print(IMT)
}
