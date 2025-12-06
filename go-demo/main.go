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
	fmt.Print(IMT)
}
