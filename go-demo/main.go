package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var count = 100
	var isReturn string = "y"
	for i := 0; i < count; i++ {
		fmt.Println("___Калькулятор индекса массы тела___")
		userHeight, userKg := getUserInput()
		IMT, err := calculateIMT(userHeight, userKg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if IMT < 16 {
			fmt.Println("У вас недостаток веса.")
		}
		outputResult(IMT)
		fmt.Print("повторить расчет? y n ")

		fmt.Scan(&isReturn)
		if isReturn == "n" {
			break
		}

	}

}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)
}

func calculateIMT(userHeight float64, userKg float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("no user height")
	}
	const IMTPower float64 = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userKg float64
	var userHeight float64

	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)

	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userKg)

	return userHeight, userKg
}
