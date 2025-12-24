package main

import (
	"demo/app-http/geo"
	"demo/app-http/weather"
	"flag"
	"fmt"
)

func main() {

	city := flag.String("city", "", "Город пользователя")
	format := flag.String("format", "1", "Формат вывода погоды")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(geoData)
	weather, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(weather)
}
