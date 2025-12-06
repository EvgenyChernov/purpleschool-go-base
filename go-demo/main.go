package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight := 1.8
	var userKg float64
	userKg = 100
	IMT := userKg / math.Pow(userHeight, 2)
	fmt.Print(IMT)
}
