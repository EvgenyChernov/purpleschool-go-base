package geo_test

import (
	"demo/app-http/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	city := "Moscow"
	expected := geo.GeoData{
		City: "London",
	}

	got, err := geo.GetMyLocation(city)
	if err != nil {
		t.Error("Ошибка получения данных о местоположении")
	}

	if got.City != expected.City {
		t.Errorf("Ожидается: %s, Получено: %s", expected.City, got.City)
	}
	t.Logf("Тест прошел успешно")
}
