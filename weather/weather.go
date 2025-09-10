package weather

import (
	"encoding/json"
	"fmt"
)

type WeatherItem struct {
	Main string
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
}

type Conditions struct {
	Summary string
}

func ParseResponse(data []byte) (Conditions, error) {
	resp := OWMResponse{}
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Conditions{}, fmt.Errorf("invalid API response %s: %w", data, err)
	}
	if len(resp.Weather) < 1 {
		return Conditions{}, fmt.Errorf("invalid API response %s: want at least one Weather element", data)
	}
	conditions := Conditions{
		Summary: resp.Weather[0].Main,
	}
	return conditions, nil
}

func FormatURL(baseURL string, latitude, longitude float64, apiKey string) string {
	return fmt.Sprintf("%s/data/2.5/weather?lat=%f&lon=%f&units=metric&appid=%s", baseURL, latitude, longitude, apiKey)
}
