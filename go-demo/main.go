package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower float64 = 2
	var userKg float64
	var userHeight float64
	fmt.Println("___Калькулятор индекса массы тела___")
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Print(result)
}
