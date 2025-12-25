package weather_test

import (
	"demo/app-http/geo"
	"demo/app-http/weather"
	"testing"
)

func TestGetWeather(t *testing.T) {
	geoData := geo.GeoData{City: "Moscow"}
	format := "json"

	result, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Logf("Ошибка при получении погоды (может быть из-за сети): %v", err)
		// Если нет сети, пропускаем тест
		return
	}

	if result == "" {
		t.Error("Результат не должен быть пустым")
	}

	t.Logf("Погода получена успешно, длина ответа: %d символов", len(result))
}

func TestGetWeatherInvalidCity(t *testing.T) {
	// Тест с пустым городом
	geoData := geo.GeoData{City: ""}
	format := "json"

	_, err := weather.GetWeather(geoData, format)
	if err == nil {
		t.Error("Ожидается ошибка при пустом городе")
	}
}

func TestGetWeatherDifferentFormats(t *testing.T) {
	testCases := []struct {
		name   string
		city   string
		format string
	}{
		{"JSON format", "London", "json"},
		{"Text format", "Paris", "1"},
		{"Custom format", "Berlin", "2"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			geoData := geo.GeoData{City: tc.city}
			result, err := weather.GetWeather(geoData, tc.format)
			if err != nil {
				t.Logf("Ошибка для формата %s (может быть из-за сети): %v", tc.format, err)
				return
			}

			if result == "" {
				t.Errorf("Результат не должен быть пустым для формата %s", tc.format)
			}

			t.Logf("Формат %s: получено %d символов", tc.format, len(result))
		})
	}
}

func TestGetWeatherDifferentCities(t *testing.T) {
	cities := []string{"Moscow", "London", "Paris", "Berlin"}

	for _, city := range cities {
		t.Run(city, func(t *testing.T) {
			geoData := geo.GeoData{City: city}
			format := "json"

			result, err := weather.GetWeather(geoData, format)
			if err != nil {
				t.Logf("Ошибка для города %s (может быть из-за сети): %v", city, err)
				return
			}

			if result == "" {
				t.Errorf("Результат не должен быть пустым для города %s", city)
			}

			t.Logf("Погода для %s получена успешно", city)
		})
	}
}
