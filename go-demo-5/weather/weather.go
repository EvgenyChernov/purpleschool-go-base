package weather

import (
	"demo/app-http/geo"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetWeather(geo geo.GeoData, format string) (string, error) {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("format", string(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
