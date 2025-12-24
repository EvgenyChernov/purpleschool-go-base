package main

import (
	"demo/app-http/geo"
	"flag"
	"fmt"
)

func main() {

	city := flag.String("city", "", "Город пользователя")
	// format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(geoData)

}
