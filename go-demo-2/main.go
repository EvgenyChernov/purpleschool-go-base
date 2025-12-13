package main

import (
	"fmt"
	"strconv"
)

func main() {
	transactions := []int{}
	var inputData string

	for {
		fmt.Println("Введите тарнзакцию или 'd' для подсчета суммы, спасибо.")
		fmt.Scan(&inputData)

		if isNumber(inputData) {
			numberTransition, _ := strconv.Atoi(inputData)
			transactions = append(transactions, numberTransition)
			fmt.Println(transactions)
		} else if inputData == "d" {
			fmt.Println("буква Д")
			sum := calculateAllTransition(transactions)
			fmt.Println("Сумма транзакций равна ", sum)
			break
		} else {
			break
		}
	}
}

func calculateAllTransition(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
