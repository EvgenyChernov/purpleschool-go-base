package main

import "fmt"

func main() {
	transactions := [3]int{5, 10, -7}
	banks := [2]string{"Тинькофф", "Альфа"}

	fmt.Println(transactions)
	fmt.Println(banks)

	patiral := transactions[1:2]
	fmt.Println(patiral)
}
