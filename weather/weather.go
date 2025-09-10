package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type OWMResponse struct {
	Weather []struct {
		Main string
	}
}

type Conditions struct {
	Summary string
}

type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: "https://api.openweathermap.org",
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c Client) FormatURL(latitude, longitude float64) string {
	return fmt.Sprintf("%s/data/2.5/weather?lat=%f&lon=%f&units=metric&appid=%s", c.BaseURL, latitude, longitude, c.APIKey)
}

func (c Client) GetWeather(latitude, longitude float64) (Conditions, error) {
	URL := c.FormatURL(latitude, longitude)

	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		return Conditions{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Conditions{}, fmt.Errorf("unexpected response status: %q", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Conditions{}, fmt.Errorf("unexpected issue reading response body: %v", err)
	}
	conditions, err := ParseResponse(data)
	if err != nil {
		return Conditions{}, err
	}
	return conditions, nil
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

func Get(apiKey string, latitude, longitude float64) (Conditions, error) {
	c := NewClient(apiKey)
	return c.GetWeather(latitude, longitude)
}
